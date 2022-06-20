package ftx

type ChannelRequest struct {
	Op      string `json:"op"`
	Channel string `json:"channel"`
	Market  string `json:"market"`
}

type ChannelResponseHeader struct {
	Channel string `json:"channel"`
	Market  string `json:"market"`
	Type    string `json:"type"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
}

type ChannelResponse[T any] struct {
	ChannelResponseHeader
	Data *T `json:"data,omitempty"`
}
