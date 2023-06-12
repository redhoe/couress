package btcerModel

import (
	"github.com/redhoe/couress/global/modeler"
	"github.com/shopspring/decimal"
)

type TransactionBtc struct {
	modeler.MysqlModel
	ChainId           uint            `json:"chain_id" gorm:""`
	Chain             *modeler.Chain  `json:"chain" gorm:""`
	CoinId            *uint           `json:"coin_id" gorm:""`
	WalletChainId     uint            `json:"wallet_chain_id" gorm:""`
	BlockNumber       int64           `json:"block_number" gorm:""`
	BlockTime         int64           `json:"block_time" gorm:""`
	Hash              string          `json:"hash" gorm:""`
	Fees              decimal.Decimal `json:"fees" gorm:"type:decimal(20)"`
	Sender            string          `json:"sender" gorm:"type:varchar(42)"`
	Receive           string          `json:"receive" gorm:""`
	Amount            decimal.Decimal `json:"amount" gorm:"type:decimal(50)"`
	SourceTransaction string          `json:"source_transaction" gorm:"type:text"`
}

func (*TransactionBtc) TableName() string {
	return "transaction_btc"
}

func NewTransactionBtc() *TransactionBtc {
	return &TransactionBtc{}
}

func (*TransactionBtc) Comment() string {
	return "Btc交易记录"
}

type TransactionBtcSort []TransactionBtc

func (s TransactionBtcSort) Len() int { return len(s) }

func (s TransactionBtcSort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s TransactionBtcSort) Less(i, j int) bool { return s[i].BlockNumber > s[j].BlockNumber }

type TransactionBtcPending struct {
	modeler.MysqlModel
	modeler.MysqlDeleteModel
	ChainId           uint            `json:"chain_id" gorm:""`
	Chain             *modeler.Chain  `json:"chain" gorm:""`
	CoinId            *uint           `json:"coin_id" gorm:""`
	WalletChainId     uint            `json:"wallet_chain_id" gorm:""`
	Hash              string          `json:"hash" gorm:""`
	Fees              decimal.Decimal `json:"fees" gorm:"type:decimal(20)"`
	Sender            string          `json:"sender" gorm:"type:varchar(42)"`
	Receive           string          `json:"receive" gorm:""`
	Amount            decimal.Decimal `json:"amount" gorm:"type:decimal(50)"`
	SourceTransaction string          `json:"source_transaction" gorm:"type:text"`
	Usd               bool            `json:"usd" bun:"usd"`
}

func (*TransactionBtcPending) TableName() string {
	return "transaction_btc_pending"
}

func NewTransactionBtcPending() *TransactionBtcPending {
	return &TransactionBtcPending{}
}

func (*TransactionBtcPending) Comment() string {
	return "Btc待上链交易记录"
}
