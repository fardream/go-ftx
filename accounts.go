package ftx

import "github.com/shopspring/decimal"

type Positions struct {
	Success bool `json:"success"`
	Result  []struct {
		Cost                         decimal.Decimal `json:"cost"`
		EntryPrice                   decimal.Decimal `json:"entryPrice"`
		EstimatedLiquidationPrice    decimal.Decimal `json:"estimatedLiquidationPrice"`
		Future                       string          `json:"future"`
		InitialMarginRequirement     decimal.Decimal `json:"initialMarginRequirement"`
		LongOrderSize                decimal.Decimal `json:"longOrderSize"`
		MaintenanceMarginRequirement decimal.Decimal `json:"maintenanceMarginRequirement"`
		NetSize                      decimal.Decimal `json:"netSize"`
		OpenSize                     decimal.Decimal `json:"openSize"`
		RealizedPnl                  decimal.Decimal `json:"realizedPnl"`
		ShortOrderSize               decimal.Decimal `json:"shortOrderSize"`
		Side                         string          `json:"side"`
		Size                         decimal.Decimal `json:"size"`
		UnrealizedPnl                decimal.Decimal `json:"unrealizedPnl"`
	} `json:"result"`
}
