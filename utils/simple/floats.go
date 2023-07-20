package simple

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

func FloatToDecimal(f float64, n int32) decimal.Decimal {
	return decimal.NewFromFloat(f).RoundFloor(n)
}

// float64 精度处理
func FloatToString(f float64) string {
	return decimal.NewFromFloat(f).RoundFloor(6).String()
}

// DecimalString num 精度 最少3位
func DecimalString(str string, num int32) string {
	if num < 4 {
		num = 4
	}
	nums := strings.Split(str, ".")
	if len(nums) != 2 {
		return str
	}
	var s, ss int32 = 0, 0
	aft := ""
	for _, i := range nums[1] {
		if i == '0' {
			s++
			continue
		}
		if ss < num-3 {
			aft += string(i)
		}
		ss++
	}
	if s > 2 {
		return fmt.Sprintf("%s.{%d}%s", nums[0], s, aft)
	}
	decimals := nums[1]
	if len(nums[1]) >= int(num) {
		decimals = decimals[:num]
	}
	return fmt.Sprintf("%s.%s", nums[0], decimals)
}
