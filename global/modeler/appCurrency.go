package modeler

import (
	"fmt"
	"github.com/demdxx/gocast"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

var DefaultCurrency = "USD"

type CurrencyExchangeRate struct {
	MysqlModel
	Uint     *string `json:"uint" gorm:"column:unit;type:VARCHAR(4);comment:单位符号"`
	Symbol   string  `json:"symbol" gorm:"comment:币种名;type:VARCHAR(20)"`
	Rate     float64 `json:"rate" gorm:"comment:汇率"`
	BaseCode string  `json:"base_code" gorm:"comment:基础币种;type:VARCHAR(20)"`
	Enable   bool    `json:"enable"  gorm:"comment:是否有效"`
	Sort     int     `json:"sort" gorm:" default:9999;comment:排序"`
}

func (*CurrencyExchangeRate) TableName() string {
	return "currency_exchange_rate"
}

func (*CurrencyExchangeRate) Comment() string {
	return "法币汇率表"
}
func NewCurrencyExchangeRate() *CurrencyExchangeRate {
	return &CurrencyExchangeRate{}
}

func (c *CurrencyExchangeRate) GetRecord(db *gorm.DB, column, value interface{}) bool {
	if err := db.Where(fmt.Sprintf("%s = ?", gocast.ToString(column)), value).Debug().
		Find(&c).Error; err != nil {
		return false
	}
	if c.Id != 0 {
		return true
	}
	return false
}

// ExchangePrice Usd单价转化为指定币种单价
func (c *CurrencyExchangeRate) ExchangePrice(usdPrice decimal.Decimal) decimal.Decimal {
	return usdPrice.Mul(decimal.NewFromFloat(c.Rate)).RoundFloor(CurrencyShowDecimal)
}

// Exchange 把对应USD价值 转化为 指定币种价值
func (c *CurrencyExchangeRate) Exchange(usdAmount decimal.Decimal) CurrencyInfo {
	info := CurrencyInfo{
		Uint:   c.Uint,
		Symbol: c.Symbol,
		Value:  usdAmount.Mul(decimal.NewFromFloat(c.Rate)).RoundFloor(CurrencyShowDecimal),
	}
	return info
}

func (c *CurrencyExchangeRate) CurrencyInfo(usdAmount decimal.Decimal) CurrencyInfo {
	info := CurrencyInfo{
		Uint:   c.Uint,
		Symbol: c.Symbol,
		Value:  usdAmount.Mul(decimal.NewFromFloat(c.Rate)).RoundFloor(2),
	}
	return info
}

func (c *CurrencyExchangeRate) ToInfoFromUsd2(usdAmount decimal.Decimal) CurrencyInfo {
	info := CurrencyInfo{
		Uint:   c.Uint,
		Symbol: c.Symbol,
		Value:  usdAmount.Mul(decimal.NewFromFloat(c.Rate)),
	}
	return info
}

type CurrencyInfo struct {
	Uint   *string         `json:"uint"`
	Symbol string          `json:"symbol"`
	Value  decimal.Decimal `json:"value"`
}

func CurrencyBySymbol(db *gorm.DB, currencySymbol string) (*CurrencyExchangeRate, error) {
	currencyExchangeRate := NewCurrencyExchangeRate()
	if err := db.Model(&CurrencyExchangeRate{}).Where("Symbol", currencySymbol).Find(&currencyExchangeRate).Error; err != nil {
		return nil, err
	}
	return currencyExchangeRate, nil
}

func ToCurrencyInfoByUsd(db *gorm.DB, currencySymbol string, usdAmount decimal.Decimal) (*CurrencyInfo, error) {
	currencyExchangeRate, err := CurrencyBySymbol(db, currencySymbol)
	if err != nil {
		return nil, err
	}
	return lo.ToPtr(currencyExchangeRate.CurrencyInfo(usdAmount)), nil
}
