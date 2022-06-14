package ftx

import (
	"time"

	"github.com/shopspring/decimal"
)

type NewOrder struct {
	Market                  string          `json:"market"`
	Side                    string          `json:"side"`
	Price                   decimal.Decimal `json:"price"`
	Type                    string          `json:"type"`
	Size                    decimal.Decimal `json:"size"`
	ReduceOnly              bool            `json:"reduceOnly"`
	Ioc                     bool            `json:"ioc"`
	PostOnly                bool            `json:"postOnly"`
	ExternalReferralProgram string          `json:"externalReferralProgram"`
	// ClientID                string  `json:"clientId"`
}

type NewOrderResponse struct {
	Success bool  `json:"success"`
	Result  Order `json:"result"`
}

type Order struct {
	CreatedAt     time.Time       `json:"createdAt"`
	FilledSize    decimal.Decimal `json:"filledSize"`
	Future        string          `json:"future"`
	ID            int64           `json:"id"`
	Market        string          `json:"market"`
	Price         decimal.Decimal `json:"price"`
	AvgFillPrice  decimal.Decimal `json:"avgFillPrice"`
	RemainingSize decimal.Decimal `json:"remainingSize"`
	Side          string          `json:"side"`
	Size          decimal.Decimal `json:"size"`
	Status        string          `json:"status"`
	Type          string          `json:"type"`
	ReduceOnly    bool            `json:"reduceOnly"`
	Ioc           bool            `json:"ioc"`
	PostOnly      bool            `json:"postOnly"`
	ClientID      string          `json:"clientId"`
}

type OpenOrders struct {
	Success bool    `json:"success"`
	Result  []Order `json:"result"`
}

type OrderHistory struct {
	Success     bool    `json:"success"`
	Result      []Order `json:"result"`
	HasMoreData bool    `json:"hasMoreData"`
}

type NewTriggerOrder struct {
	Market           string          `json:"market"`
	Side             string          `json:"side"`
	Size             decimal.Decimal `json:"size"`
	Type             string          `json:"type"`
	ReduceOnly       bool            `json:"reduceOnly"`
	RetryUntilFilled bool            `json:"retryUntilFilled"`
	TriggerPrice     decimal.Decimal `json:"triggerPrice,omitempty"`
	OrderPrice       decimal.Decimal `json:"orderPrice,omitempty"`
	TrailValue       decimal.Decimal `json:"trailValue,omitempty"`
}

type NewTriggerOrderResponse struct {
	Success bool         `json:"success"`
	Result  TriggerOrder `json:"result"`
}

type TriggerOrder struct {
	CreatedAt        time.Time       `json:"createdAt"`
	Error            string          `json:"error"`
	Future           string          `json:"future"`
	ID               int64           `json:"id"`
	Market           string          `json:"market"`
	OrderID          int64           `json:"orderId"`
	OrderPrice       decimal.Decimal `json:"orderPrice"`
	ReduceOnly       bool            `json:"reduceOnly"`
	Side             string          `json:"side"`
	Size             decimal.Decimal `json:"size"`
	Status           string          `json:"status"`
	TrailStart       decimal.Decimal `json:"trailStart"`
	TrailValue       decimal.Decimal `json:"trailValue"`
	TriggerPrice     decimal.Decimal `json:"triggerPrice"`
	TriggeredAt      string          `json:"triggeredAt"`
	Type             string          `json:"type"`
	OrderType        string          `json:"orderType"`
	FilledSize       decimal.Decimal `json:"filledSize"`
	AvgFillPrice     decimal.Decimal `json:"avgFillPrice"`
	OrderStatus      string          `json:"orderStatus"`
	RetryUntilFilled bool            `json:"retryUntilFilled"`
}

type OpenTriggerOrders struct {
	Success bool           `json:"success"`
	Result  []TriggerOrder `json:"result"`
}

type TriggerOrderHistory struct {
	Success     bool           `json:"success"`
	Result      []TriggerOrder `json:"result"`
	HasMoreData bool           `json:"hasMoreData"`
}

type Triggers struct {
	Success bool `json:"success"`
	Result  []struct {
		Error      string          `json:"error"`
		FilledSize decimal.Decimal `json:"filledSize"`
		OrderSize  decimal.Decimal `json:"orderSize"`
		OrderID    int64           `json:"orderId"`
		Time       time.Time       `json:"time"`
	} `json:"result"`
}
