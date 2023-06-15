package tronerModel

import (
	"fmt"
	"github.com/redhoe/couress/global/modeler"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TxType string

const (
	RecordTypeInterior TxType = "interior"
	RecordTypeScan     TxType = "scan"
)

type QueueTransactionStatus int64

const (
	StatusInteriorPending QueueTransactionStatus = 4
	StatusInteriorFail    QueueTransactionStatus = 3
	StatusInteriorSuccess QueueTransactionStatus = 2 // pending = 1
	StatusSuccess         QueueTransactionStatus = 1 // 成功 1
	StatusFail            QueueTransactionStatus = 0 // 失败 0
	StatusPending         QueueTransactionStatus = 6
	StatusConfirmed       QueueTransactionStatus = 7
)

type TransactionTron struct {
	modeler.MysqlModel
	RecordType       TxType           `json:"record_type" gorm:"type:varchar(50);comment:记录类型:interior 内部发起 scan 扫块记录"`
	TransferCoinType modeler.CoinType `json:"transfer_coin_type" gorm:"type:varchar(50);comment:转账币种类型"`
	ChainId          uint             `json:"chain_id" gorm:"comment:链ID 本站"`
	CoinId           uint             `json:"coin_id" gorm:"comment:币种ID 本站"`
	BlockNumber      int64            `json:"block_number" gorm:"comment:区块号"`
	BlockTime        int64            `json:"block_time"  gorm:"comment:区块扫描时间"`
	Symbol           string           `json:"symbol" gorm:"type:varchar(20);comment:币种名"`
	SymbolDecimal    int32            `json:"symbol_decimal" gorm:"comment:币种精度"`
	ContractAddress  string           `json:"contract_address" gorm:"type:varchar(50);comment:合约地址"` // Trc10 : 1000016
	Sender           string           `json:"sender" gorm:"type:varchar(50);comment:发送地址"`
	Receive          string           `json:"receive" gorm:"type:varchar(50);comment:接收地址"`
	Amount           decimal.Decimal  `json:"amount" gorm:"type:decimal(50,18);comment:交易金额"`
	GasLimit         decimal.Decimal  `json:"gas_limit" gorm:"type:decimal(50,18);comment:最大矿工费"`
	Hash             string           `json:"hash" gorm:"type:varchar(80);unique;not null;comment:交易ID"`
	// 内部交易专用
	TxBeforeSign *string `json:"tx_before_sign" gorm:"type:text;comment:签名前交易信息"`
	TxAfterSign  *string `json:"tx_after_sign" gorm:"type:text;comment:签名后交易信息"`
	// 区块扫描专用
	GasAmount         decimal.Decimal `json:"gas_amount" gorm:"comment:扫块获得实际试用矿工费"`
	GasAmountUsd      decimal.Decimal `json:"gas_amount_usd" gorm:"comment:矿工费折合USD"`
	Gas               int64           `json:"gas" gorm:"comment:实际矿工费gaWei"`
	NetFee            int64           `json:"net_fee" gorm:"comment:带宽费用gaWei"`
	EnergyFee         int64           `json:"energy_fee" gorm:"comment:能量费用gaWei"`
	NetUsage          int64           `json:"net_usage" gorm:"comment:带宽消耗数"`
	EnergyUsageTotal  int64           `json:"energy_usage_total" gorm:"comment:能量总消耗"`
	OriginEnergyUsage int64           `json:"origin_energy_usage" gorm:"comment:合约消耗能量"`
	// 公共消息
	Status        QueueTransactionStatus `json:"status" gorm:"comment:交易状态:通用 1 待签名与广播 2 广播成功 3 广播失败 4 交易成功 5 交易失败"`
	Message       string                 `json:"message" gorm:"type:text;comment:消息内容 失败原因"`
	IsRead        bool                   `json:"is_read" gorm:"default:false;comment:消息是否已读"`
	SenderIsRead  bool                   `json:"sender_is_read" gorm:"default:false;comment:发送者是否已读"`
	ReceiveIsRead bool                   `json:"receive_is_read" gorm:"default:false;comment:接受者是否已读"`
}

func (*TransactionTron) TableName() string {
	return "transaction_tron"
}

func (*TransactionTron) Comment() string {
	return "tron区块交易记录"
}

func NewTransactionTron() *TransactionTron {
	return &TransactionTron{}
}

func (t *TransactionTron) GetRecord(db *gorm.DB, column, value string) bool {
	if err := db.Where(fmt.Sprintf("%s = ?", column), value).
		Find(&t).Error; err != nil {
		return false
	}
	if t.Id != 0 {
		return true
	}
	return false
}

func (t *TransactionTron) UpdateOrCreate(db *gorm.DB) error {
	txFind := NewTransactionTron()
	if err := db.Where("hash", t.Hash).
		Find(&txFind).Error; err != nil {
		return err
	}
	if txFind.Id != 0 {
		// 修改交易
		//t.Id = txFind.Id
		//t.TransferType = txFind.TransferType
		txFind.Status = t.Status
		txFind.Message = t.Message
		txFind.ContractAddress = t.ContractAddress
		txFind.GasAmount = t.GasAmount
		txFind.GasAmountUsd = t.GasAmountUsd
		txFind.Amount = t.Amount
		txFind.BlockNumber = t.BlockNumber
		return db.Save(&txFind).Debug().Error
	}
	// 创建交易
	return db.Create(&t).Error
}

func (*TransactionTron) GetAllPage(db *gorm.DB, paging *modeler.Paging) ([]TransactionTron, error) {
	db = db.Where("Status < ?", StatusInteriorFail) // 过滤未成功的订单
	results := make([]TransactionTron, 0)
	var err error
	err = db.Model(&TransactionTron{}).Debug().Count(&paging.Total).Error
	if err != nil {
		return results, err
	}
	paging.GetPages()
	if paging.Total < 1 {
		return results, nil
	}
	//Omit("content").
	err = db.Model(&TransactionTron{}).Debug().
		Order("id desc").Limit(paging.PageSize).Offset(paging.StartNums).Find(&results).Error
	return results, err
}

func (*TransactionTron) CountReadNum(db *gorm.DB) (nums int64, err error) {
	db = db.Where("Status < ?", StatusInteriorSuccess) // 未成功的订单
	err = db.Model(&TransactionTron{}).Debug().Count(&nums).Error
	return
}

type TransactionTronList []TransactionTron

type TxInfo struct {
	Direction             string                 `json:"direction"`
	RecordType            TxType                 `json:"record_type" gorm:"type:varchar(50);comment:记录类型:interior 内部发起 scan 扫块记录"`
	TransferCoinType      modeler.CoinType       `json:"transfer_coin_type" gorm:"type:varchar(50);comment:转账币种类型"`
	ChainId               uint                   `json:"chain_id" gorm:"comment:链ID 本站"`
	CoinId                uint                   `json:"coin_id" gorm:"comment:币种ID 本站"`
	BlockNumber           int64                  `json:"block_number" gorm:"comment:区块号"`
	BlockTime             int64                  `json:"block_time"  gorm:"comment:区块扫描时间"`
	Symbol                string                 `json:"symbol" gorm:"type:varchar(20);comment:币种名"`
	SymbolDecimal         int32                  `json:"symbol_decimal" gorm:"comment:币种精度"`
	ContractAddress       string                 `json:"contract_address" gorm:"type:varchar(50);comment:合约地址"` // Trc10 : 1000016
	Sender                string                 `json:"sender" gorm:"type:varchar(50);comment:发送地址"`
	Receive               string                 `json:"receive" gorm:"type:varchar(50);comment:接收地址"`
	Amount                decimal.Decimal        `json:"amount" gorm:"type:decimal(50,18);comment:交易金额"`
	GasLimit              decimal.Decimal        `json:"gas_limit" gorm:"type:decimal(50,18);comment:最大矿工费"`
	Hash                  string                 `json:"hash" gorm:"type:varchar(80);unique;not null;comment:交易ID"`
	Status                QueueTransactionStatus `json:"status" gorm:"comment:交易状态:通用 1 待签名与广播 2 广播成功 3 广播失败 4 交易成功 5 交易失败"`
	Message               string                 `json:"message" gorm:"type:text;comment:消息内容 失败原因"`
	GasAmount             interface{}            `json:"gas_amount" gorm:"comment:扫块获得实际试用矿工费"`
	GasAmountUsd          interface{}            `json:"gas_amount_usd" gorm:"comment:矿工费折合USD"`
	GasAmountCurrencyInfo interface{}            `json:"gas_amount_currency_info"`
	IsRead                bool                   `json:"is_read" gorm:"default:false;comment:消息是否已读"`
	IsPending             bool                   `json:"is_pending" gorm:"default:false;comment:消息是否已读"`
	Gas                   int64                  `json:"gas" gorm:"comment:实际矿工费gaWei"`
	NetUsage              int64                  `json:"net_usage" gorm:"comment:带宽消耗数"`
	EnergyUsageTotal      int64                  `json:"energy_usage_total" gorm:"comment:能量总消耗"`
	OriginEnergyUsage     int64                  `json:"origin_energy_usage" gorm:"comment:合约消耗能量"`
	CoinInfoObj           *modeler.CoinInfo      `json:"coin" gorm:"comment:币种信息"`
	WebUrl                string                 `json:"web_url"`
}

func (t TransactionTronList) ToInfo(db *gorm.DB, webUrl, walletAddress string, currencyEngine *modeler.CurrencyExchangeRate) []TxInfo {
	return lo.Map(t, func(item TransactionTron, index int) TxInfo {
		if item.Receive == walletAddress && item.Status < StatusInteriorSuccess && !item.ReceiveIsRead {
			item.ReceiveIsRead = true
			db.Save(&item)
		}
		if item.Sender == walletAddress && item.Status < StatusInteriorSuccess && !item.SenderIsRead {
			item.SenderIsRead = true
			db.Save(&item)
		}
		coin := modeler.NewCoin()
		coin.Type = item.TransferCoinType
		coin.Decimal = item.SymbolDecimal
		coin.Symbol = item.Symbol
		coin.Icon = ""
		coin.Website = ""
		coin.Address = item.ContractAddress
		if item.TransferCoinType != modeler.CoinTypeTRX {
			_ = coin.GetRecord(db, "address", item.ContractAddress)
		}
		return TxInfo{
			IsPending:             lo.If(item.Status == StatusInteriorSuccess, true).Else(false),
			Direction:             lo.If(item.Receive == walletAddress, "in").Else("out"),
			RecordType:            item.RecordType,
			TransferCoinType:      item.TransferCoinType,
			ChainId:               item.ChainId,
			CoinId:                item.CoinId,
			BlockNumber:           item.BlockNumber,
			BlockTime:             item.BlockTime,
			Symbol:                item.Symbol,
			SymbolDecimal:         item.SymbolDecimal,
			ContractAddress:       item.ContractAddress,
			Sender:                item.Sender,
			Receive:               item.Receive,
			Amount:                item.Amount.RoundFloor(modeler.TronCoinShowDecimal),
			GasLimit:              item.GasLimit,
			Hash:                  item.Hash,
			Status:                item.Status,
			Message:               item.Message,
			GasAmount:             item.GasAmount,
			GasAmountUsd:          item.GasAmountUsd,
			GasAmountCurrencyInfo: currencyEngine.Exchange(item.GasAmountUsd),
			IsRead:                item.IsRead,            //     bool                   `json:"is_read" gorm:"default:false;comment:消息是否已读"`
			Gas:                   item.Gas,               //int64           `json:"Gas" gorm:"comment:实际矿工费gaWei"`
			NetUsage:              item.NetUsage,          //          int64           `json:"NetUsage" gorm:"comment:带宽消耗数"`
			EnergyUsageTotal:      item.EnergyUsageTotal,  //int64           `json:"EnergyUsageTotal" gorm:"comment:能量总消耗"`
			OriginEnergyUsage:     item.OriginEnergyUsage, //int64           `json:"OriginEnergyUsage" gorm:"comment:合约消耗能量"`
			CoinInfoObj:           coin.CoinInfo(),
			WebUrl:                fmt.Sprintf("%stransaction/%s", webUrl, item.Hash),
		}
	})
}
