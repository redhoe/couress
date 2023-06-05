package simple

import "github.com/shopspring/decimal"

func FloatToDecimal(f float64, n int32) decimal.Decimal {
	return decimal.NewFromFloat(f).RoundFloor(n)
}
