package etherModel

import (
	"github.com/redhoe/couress/global/modeler"
	"github.com/shopspring/decimal"
)

type TransactionEth struct {
	modeler.MysqlModel
	modeler.Transaction
	Status            int             `json:"status"`
	GasPrice          decimal.Decimal `json:"gas_price" gorm:"type:decimal(20)"`
	GasLimit          int64           `json:"gas_limit" gorm:""`
	SourceTransaction string          `json:"source_transaction" gorm:"type:text"`
	SourceReceipt     string          `json:"source_receipt" gorm:"type:text"`
	Nonce             int64           `json:"nonce" gorm:""`
	Message           *string         `json:"message" gorm:""`
	GasAmount         decimal.Decimal `json:"gas_amount,omitempty" gorm:"type:decimal(20)"`
	GasAmountUsd      decimal.Decimal `json:"-" gorm:"type:decimal(10,4)"`
}

func (*TransactionEth) TableName() string {
	return "transaction_eth"
}

func NewTransactionEth() *TransactionEth {
	return &TransactionEth{}
}

func (*TransactionEth) Comment() string {
	return "Eth交易记录"
}

func (t *TransactionEth) TransactionEthResponse() TransactionEthResponse {
	resp := TransactionEthResponse{
		TransactionResponse:   t.Transaction.TransactionResponse(),
		Status:                t.Status,
		Direction:             "",
		Message:               t.Message,
		GasLimit:              t.GasLimit,
		GasPrice:              t.GasPrice,
		GasAmount:             t.GasPrice.Mul(decimal.NewFromInt(t.GasLimit)).Div(decimal.New(1, 18)).RoundFloor(modeler.EthCoinBalanceDecimal),
		GasAmountCurrencyInfo: nil,
		GasPriceView:          "",
		WebUrl:                "",
		Nonce:                 t.Nonce,
		IsPending:             false,
	}

	return resp
}
