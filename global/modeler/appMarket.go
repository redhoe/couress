package modeler

import (
	coingecko "github.com/superoo7/go-gecko/v3"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type MarketCoin struct {
	MysqlModel
	ApiId  string `json:"api_id" gorm:"unique"`
	Symbol string `json:"symbol" gorm:""`
	Name   string `json:"name" gorm:""`
}

func (MarketCoin) TableName() string {
	return "market_coin"
}

func (MarketCoin) Comment() string {
	return "市场所有币种"
}

func NewMarketCoin() *MarketCoin {
	return &MarketCoin{}
}

func InitMarketCoinTable(db *gorm.DB) error {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	client := coingecko.NewClient(httpClient)
	coinList, err := client.CoinsList()
	if err != nil {
		return err
	}
	for _, item := range *coinList {
		marketCoin := NewMarketCoin()
		err := db.Model(&marketCoin).Where("api_id", item.ID).Find(&marketCoin).Error
		if err != nil {
			return err
		}
		if marketCoin.Id == 0 {
			newMarketCoin := MarketCoin{
				ApiId:  item.ID,
				Symbol: item.Symbol,
				Name:   item.Name,
			}
			err = db.Create(&newMarketCoin).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Market struct {
	MysqlModel
	ApiId                        string    `json:"api_id" gorm:"unique"`
	Symbol                       string    `json:"symbol" gorm:""`
	Name                         string    `json:"name" gorm:""`
	Image                        string    `json:"image" gorm:""`
	CurrentPrice                 float64   `json:"current_price" gorm:""`
	MarketCap                    int64     `json:"market_cap" gorm:""`
	MarketCapRank                int       `json:"market_cap_rank" gorm:""`
	TotalVolume                  int64     `json:"total_volume" gorm:""`
	High24h                      float64   `json:"high_24h" gorm:"column:high_24h"`
	Low24h                       float64   `json:"low_24h" gorm:"column:low_24h"`
	PriceChange24h               float64   `json:"price_change_24h" gorm:"column:price_change_24h"`
	PriceChangePercentage24H     float64   `json:"price_change_percentage_24h" gorm:"column:price_change_percentage_24h"`
	MarketCapChange24H           int64     `json:"market_cap_change_24h" gorm:"column:market_cap_change_24h"`
	MarketCapChangePercentage24H float64   `json:"market_cap_change_percentage_24h" gorm:"column:market_cap_change_percentage_24h"`
	CirculatingSupply            int       `json:"circulating_supply" gorm:""`
	TotalSupply                  int       `json:"total_supply" gorm:""`
	Ath                          int       `json:"ath" gorm:""`
	AthChangePercentage          float64   `json:"ath_change_percentage" gorm:""`
	AthDate                      time.Time `json:"ath_date" gorm:""`
	LastUpdated                  time.Time `json:"last_updated" gorm:""`
	Sort                         int       `json:"sort" gorm:"default:9999"`
	Hot                          bool      `json:"hot" gorm:"default:false"`
}

func NewMarket() *Market {
	return &Market{}
}

func (Market) TableName() string {
	return "market"
}

func (Market) Comment() string {
	return "当前市场币种行情"
}
