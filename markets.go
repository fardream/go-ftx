package ftx

import (
	"time"

	"github.com/shopspring/decimal"
)

type HistoricalPrices struct {
	Success bool `json:"success"`
	Result  []struct {
		Close     decimal.Decimal `json:"close"`
		High      decimal.Decimal `json:"high"`
		Low       decimal.Decimal `json:"low"`
		Open      decimal.Decimal `json:"open"`
		StartTime time.Time       `json:"startTime"`
		Volume    decimal.Decimal `json:"volume"`
	} `json:"result"`
}

type Trades struct {
	Success bool `json:"success"`
	Result  []struct {
		ID          int64           `json:"id"`
		Liquidation bool            `json:"liquidation"`
		Price       decimal.Decimal `json:"price"`
		Side        string          `json:"side"`
		Size        decimal.Decimal `json:"size"`
		Time        time.Time       `json:"time"`
	} `json:"result"`
}
