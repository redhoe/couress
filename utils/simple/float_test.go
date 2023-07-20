package simple

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestFloor(t *testing.T) {
	t.Log("RoundDown:", decimal.NewFromFloat(1233324.23456789).RoundDown(5).String())
	t.Log("RoundFloor:", decimal.NewFromFloat(124422.23456789).RoundFloor(5).String())
	t.Log("RoundCash:", decimal.NewFromFloat(1.23456789).RoundCash(5).String())
	t.Log("RoundUp:", decimal.NewFromFloat(1.23456789).RoundUp(5).String())
}

func TestStr(t *testing.T) {
	tsStr := "12.00023456"
	t.Log(DecimalString(tsStr, 6))
}
