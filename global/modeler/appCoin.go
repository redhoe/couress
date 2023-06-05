package modeler

import (
	"fmt"
	"github.com/demdxx/gocast"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CoinType string

func (s CoinType) String() string {
	return string(s)
}

const (
	CoinTypeERC20 CoinType = "ERC20"
	CoinTypeBEP20 CoinType = "BEP20"
	CoinTypeHTC20 CoinType = "HTC20"
	CoinTypeTRX   CoinType = "TRX"
	CoinTypeTRC20 CoinType = "TRC20"
	CoinTypeTRC10 CoinType = "TRC10"
)

type Coin struct {
	MysqlModel
	Name        string           `json:"name" gorm:"type:varchar(200);comment:全称"`
	ChainId     uint             `json:"chain_id" gorm:";comment:全称"`
	Chain       *Chain           `json:"chain" gorm:"foreignKey:chain_id;references:id;"`
	Symbol      string           `json:"symbol" bun:"comment:币名称"`
	Decimal     int32            `json:"decimal" gorm:"comment:精度"`
	Icon        string           `json:"icon" gorm:"comment:图标"`
	Website     string           `json:"website" gorm:"comment:官网"`
	Type        CoinType         `json:"type" gorm:"comment:币类型"`
	Address     string           `json:"address" gorm:"comment:合约地址"`
	Sort        int              `json:"sort" gorm:"default:9999;comment:排序"`
	Hot         bool             `json:"hot" gorm:"comment:是否热门"`
	MarketApiId *string          `json:"-" gorm:"comment:市场id"`
	UsdRate     *decimal.Decimal `json:"-" gorm:"comment:汇率"`
}

func (*Coin) TableName() string {
	return "coin"
}

func (*Coin) Comment() string {
	return "币信息"
}

func NewCoin() *Coin {
	return &Coin{}
}

func (this *Coin) GetRecord(db *gorm.DB, column, value interface{}) bool {
	if err := db.Where(fmt.Sprintf("%s = ?", gocast.ToString(column)), value).
		Find(&this).Error; err != nil {
		return false
	}
	if this.Id != 0 {
		return true
	}
	return false
}

func (*Coin) GetList(db *gorm.DB, chainId uint) (CoinList, error) {
	results := make(CoinList, 0)
	err := db.Model(&Coin{}).
		Where("chain_id = ?", chainId).
		Order("sort asc").Find(&results).Error
	return results, err
}

func (*Coin) GetHotList(db *gorm.DB, chainId uint) (CoinList, error) {
	results := make(CoinList, 0)
	err := db.Model(&Coin{}).
		Where("chain_id = ?", chainId).
		Where("hot = true").
		Order("sort asc").Find(&results).Error
	return results, err
}

func (c *Coin) CoinInfo() *CoinInfo {
	return &CoinInfo{
		Name:    c.Name,
		Symbol:  c.Symbol,
		Decimal: c.Decimal,
		Icon:    c.Icon,
		Website: c.Website,
		Address: c.Address,
		Type:    c.Type,
		UsdRate: c.UsdRate,
	}
}

type CoinList []Coin

type CoinInfo struct {
	Name    string           `json:"name"`
	Symbol  string           `json:"symbol"`
	Decimal int32            `json:"decimal"`
	Icon    string           `json:"icon"`
	Website string           `json:"website"`
	Address string           `json:"address"`
	Type    CoinType         `json:"type"`
	UsdRate *decimal.Decimal `json:"usdRate"`
}

func (list CoinList) ToInfo() []CoinInfo {
	return lo.Map(list, func(item Coin, index int) CoinInfo {
		return CoinInfo{
			Name:    item.Name,
			Symbol:  item.Symbol,
			Decimal: item.Decimal,
			Icon:    item.Icon,
			Website: item.Website,
			Address: item.Address,
			Type:    item.Type,
		}
	})
}
