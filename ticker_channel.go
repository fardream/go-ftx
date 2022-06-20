package ftx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
)

type TickerChannelUpdate struct {
	Bid  *decimal.Decimal `json:"bid,omitempty"`
	Ask  *decimal.Decimal `json:"ask,omitempty"`
	Last *decimal.Decimal `json:"last,omitempty"`
	Time decimal.Decimal  `json:"time"`
}

const pingMsg = "{op: \"ping\"}"

func isNormalClose(err error) bool {
	return websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway)
}

func runPing(ctx context.Context, conn *websocket.Conn) error {
	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()

	err_count := 0

ping_loop:
	for {
		select {
		case <-ticker.C:
			if err := conn.WriteControl(websocket.PingMessage, []byte(pingMsg), time.Now().Add(time.Second*2)); err != nil {
				if isNormalClose(err) {
					break ping_loop
				}

				err_count++
				log.Warnf("failed to send ping msg: %w", err)
				if err_count >= 15 {
					return fmt.Errorf("failed to send ping message: %w", err)
				}
			}
		case <-ctx.Done():
			break ping_loop
		}
	}

	return nil
}

func SubscribeTicker(ctx context.Context, endpoint, market, channel string, outputChan chan<- *ChannelResponse[TickerChannelUpdate]) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	inner_ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 60 * time.Second,
	}

	conn, rsp, err := dialer.DialContext(inner_ctx, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to websocket %s: %w\nresponse is %#v", endpoint, err, rsp)
	}
	defer conn.Close()

	wg.Add(1)
	go func() {
		wg.Done()
		runPing(inner_ctx, conn)
	}()

	msg_chan := make(chan []byte)
	err_chan := make(chan error)
	// start looping read
	wg.Add(1)
	go func() {
		defer wg.Done()
		// close message channel
		defer close(msg_chan)
		// close error channel.
		defer close(err_chan)

		select {
		case err_chan <- loopRead(inner_ctx, conn, msg_chan):
		case <-inner_ctx.Done():
		}
	}()

	sub_request := ChannelRequest{
		Channel: channel,
		Market:  market,
		Op:      "subscribe",
	}

	unsub_request := ChannelRequest{
		Channel: channel,
		Market:  market,
		Op:      "unsubscribe",
	}

	if err := conn.WriteJSON(sub_request); err != nil {
		return fmt.Errorf("failed to send subscribe request: %w", err)
	}

write_loop:
	for {
		select {
		case <-inner_ctx.Done():
			break write_loop
		case msg, ok := <-msg_chan:
			if !ok {
				break write_loop
			}

			resp := new(ChannelResponse[TickerChannelUpdate])

			err := json.Unmarshal(msg, resp)
			if err != nil {
				log.Warnf("faile to parse data: %v", err)
				continue write_loop
			}

			if resp.Type == "error" {
				cancel()
				unsubscribeAndClose(conn, &unsub_request)
				return fmt.Errorf("subscription erro: %s", resp.Msg)
			}

			select {
			case <-inner_ctx.Done():
				break write_loop
			case outputChan <- resp:
			}

		case err := <-err_chan:
			log.Warnf("received err: %v", err)
			if err != nil {
				return err
			}
		}
	}

	if err := unsubscribeAndClose(conn, &unsub_request); err != nil {
		return err
	}

	// drain the error channel
	return <-err_chan
}

func unsubscribeAndClose(conn *websocket.Conn, unsubscribe any) error {
	// request cancelled.
	// Close the connection.
	if unsubscribe != nil {
		if err := conn.WriteJSON(unsubscribe); err != nil {
			return err
		}
	}
	if err := conn.WriteControl(websocket.CloseMessage, nil, time.Now().Add(time.Second*2)); err != nil {
		return err
	}

	return nil
}

func loopRead(ctx context.Context, conn *websocket.Conn, output chan<- []byte) error {
	for {

		_, msg, err := conn.ReadMessage()
		if err != nil {
			if isNormalClose(err) {
				return nil
			}
			log.Warnf("error reading websocket: %#v", err)
			log.Warnf("quit read loop")
			return err
		}

		log.Debugf("message received from websocket: %s", string(msg))

		select {
		case output <- msg:
		case <-ctx.Done():
			return nil
		}
	}
}
