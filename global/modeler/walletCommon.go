package modeler

import (
	"errors"
	"fmt"
	"github.com/demdxx/gocast"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const (
	EthCoinBalanceDecimal int32 = 8 // 数字货币保留显示小数位
	BtcCoinShowDecimal    int32 = 10
	TronCoinShowDecimal   int32 = 4

	CurrencyShowDecimal int32 = 2 // 转化为法币价值保留小数位
	PriceDecimal              = 4 // 货币单价转化为法币价值保留小数位
)

type WalletType string

const (
	WalletTypeIdentity WalletType = "identity"
	WalletTypeUser     WalletType = "user"
	WalletTypeGKey     WalletType = "gkey"
)

var WalletTypeList = []WalletType{
	WalletTypeIdentity,
	WalletTypeUser,
	WalletTypeGKey,
}

type WalletChain struct {
	MysqlModel
	ChainId          uint            `json:"chain_id" gorm:"comment:链Id"`
	IdentityId       uint            `json:"identity_id" gorm:"comment:身份ID"`
	Identity         string          `json:"identity" gorm:"type:varchar(50);comment:身份ID"`
	ChainType        ChainType       `json:"chain_type" gorm:"type:varchar(20);comment:链类型"`
	WalletType       WalletType      `json:"wallet_type" gorm:"type:varchar(20);comment:钱包类型"`
	WalletAddress    string          `json:"wallet_address" gorm:"type:varchar(60);comment:钱包地址"`
	FirstBlockNumber *int64          `json:"first_block_number" gorm:"comment:首块"`
	LastBlockNumber  *int64          `json:"last_block_number" gorm:"comment:尾块"`
	Balance          decimal.Decimal `json:"balance" gorm:"type:decimal(50,18);comment:余额"`
}

func (*WalletChain) TableName() string {
	return "wallet_chain"
}

func (*WalletChain) Comment() string {
	return "用户钱包链信息"
}

func NewWalletChain() *WalletChain {
	return &WalletChain{}
}

func (this *WalletChain) GetOrCreate(db *gorm.DB) error {
	if err := db.
		Where(fmt.Sprintf("%s = ?", gocast.ToString("Wallet_Type")), this.WalletType).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("Chain_Id")), this.ChainId).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("Identity")), this.Identity).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("Identity_Id")), this.IdentityId).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("Wallet_Address")), this.WalletAddress).
		Debug().
		Find(&this).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return db.Create(&this).Error
		}
		return errors.New("SystemError")
	}
	if this.Id != 0 {
		return nil
	} else {
		return db.Create(&this).Error
	}
}

func (this *WalletChain) GetRecord(db *gorm.DB, column1, value1, column2, value2 interface{}) bool {
	if err := db.
		Where(fmt.Sprintf("%s = ?", gocast.ToString(column1)), value1).
		Where(fmt.Sprintf("%s = ?", gocast.ToString(column2)), value2).
		Debug().
		Find(&this).Error; err != nil {
		return false
	}
	if this.Id != 0 {
		return true
	}
	return false
}

func (this *WalletChain) UpdateBalance(db *gorm.DB) error {
	_ = this.GetRecord(db, "chain_id", this.ChainId, "wallet_address", this.WalletAddress)

	if err := db.Model(&WalletChain{}).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("Chain_Id")), this.ChainId).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("Wallet_Address")), this.WalletAddress).
		Updates(map[string]interface{}{
			"balance": this.Balance,
		}).Error; err != nil {
		return err
	}
	return nil
}

type WalletCoin struct {
	MysqlModel
	IdentityId    uint            `json:"identity_id" gorm:"comment:身份ID"`
	WalletChainId uint            `json:"wallet_chain_id" gorm:"comment:链Id"`
	CoinId        uint            `json:"coin_id" gorm:"comment:币种ID"`
	Coin          *Coin           `json:"coin" gorm:"comment:钱包币种"`
	Symbol        string          `json:"symbol" gorm:"type:varchar(20);comment:币种"`
	Address       string          `json:"address" gorm:"type:varchar(60);comment:地址"`
	Decimal       int32           `json:"decimal" gorm:"comment:精度"`
	Balance       decimal.Decimal `json:"balance" gorm:"type:decimal(50,18);comment:可用余额"`
}

func (*WalletCoin) TableName() string {
	return "wallet_coin"
}

func (*WalletCoin) Comment() string {
	return "用户钱包币种信息"
}

func NewWalletCoin() *WalletCoin {
	return &WalletCoin{}
}

func (c *WalletCoin) UpdateTokenBalance(db *gorm.DB, token string) error {
	coin := NewCoin()
	if !coin.GetRecord(db, "address", token) {
		return errors.New("coin info error")
	}
	c.Symbol = coin.Symbol
	c.CoinId = coin.Id
	if err := c.GetOrCreate(db); err != nil {
		return err
	}
	c.Decimal = coin.Decimal
	if err := db.Model(&WalletCoin{}).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("wallet_chain_id")), c.WalletChainId).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("coin_id")), coin.Id).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("address")), c.Address).Debug().
		Updates(map[string]interface{}{
			"balance":   c.Balance,
			"`decimal`": c.Decimal,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (c *WalletCoin) GetOrCreate(db *gorm.DB) error {
	if err := db.
		Where(fmt.Sprintf("%s = ?", gocast.ToString("identity_id")), c.IdentityId).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("symbol")), c.Symbol).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("coin_id")), c.CoinId).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("wallet_chain_id")), c.WalletChainId).
		Where(fmt.Sprintf("%s = ?", gocast.ToString("address")), c.Address).Debug().
		First(&c).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return db.Create(&c).Error
		}
		return errors.New("SystemError")
	}
	if c.Id != 0 {
		return nil
	} else {
		return db.Create(&c).Error
	}
}

type WalletCoinList []WalletCoin
