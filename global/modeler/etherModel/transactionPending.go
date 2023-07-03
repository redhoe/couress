package etherModel

import (
	"github.com/redhoe/couress/global/modeler"
	"github.com/shopspring/decimal"
)

type TransactionEthPending struct {
	modeler.MysqlModel
	modeler.MysqlDeleteModel
	ChainId           uint            `json:"chain_id" gorm:""`
	Chain             *modeler.Chain  `json:"chain" gorm:""`
	CoinId            *uint           `json:"coin_id" gorm:""`
	CoinDecimal       int32           `json:"coin_decimal" gorm:""`
	Hash              string          `json:"hash" gorm:"index;type:varchar(100)"`
	Sender            string          `json:"sender" gorm:"index;type:varchar(42)"`
	Receive           string          `json:"receive" gorm:"index;"`
	Amount            decimal.Decimal `json:"amount" gorm:"type:decimal(50)"`
	GasPrice          decimal.Decimal `json:"gas_price" gorm:"type:decimal(20)"`
	GasLimit          int64           `json:"gas_limit" gorm:""`
	Nonce             int64           `json:"nonce" gorm:""`
	SourceTransaction string          `json:"-" gorm:"type:text"`
}

func (*TransactionEthPending) TableName() string {
	return "transaction_eth_pending"
}

func NewTransactionEthPending() *TransactionEthPending {
	return &TransactionEthPending{}
}

func (*TransactionEthPending) Comment() string {
	return "Eth待上块交易记录"
}

func (t *TransactionEthPending) TransactionEthResponse() TransactionEthResponse {
	resp := TransactionEthResponse{
		TransactionResponse: modeler.TransactionResponse{
			BlockNumber: 0,
			BlockTime:   t.CreatedAt.Unix(),
			Hash:        t.Hash,
			Sender:      t.Sender,
			Receive:     t.Receive,
			Amount:      t.Amount,
		},
		Status:    0,
		Direction: "",
		Message:   nil,
		GasLimit:  t.GasLimit,
		GasPrice:  t.GasPrice,
		GasAmount: t.GasPrice.Mul(decimal.NewFromInt(t.GasLimit)).Div(decimal.New(1, 18)),
		WebUrl:    "",
		Nonce:     t.Nonce,
		IsPending: true,
	}

	if *t.CoinId > 0 {
		resp.Amount = resp.Amount.Div(decimal.New(1, int32(t.CoinDecimal))).RoundFloor(modeler.EthCoinBalanceDecimal)
	} else {
		resp.Amount = resp.Amount.Div(decimal.New(1, 18)).RoundFloor(modeler.EthCoinBalanceDecimal)
	}

	return resp
}

func (t *TransactionEthPending) TransactionEthResponseWithDirection(walletAddress string) TransactionEthResponse {
	resp := t.TransactionEthResponse()

	if walletAddress == resp.Sender {
		resp.Direction = "out"
	}

	if walletAddress == resp.Receive {
		resp.Direction = "in"
	}

	if walletAddress == resp.Sender && walletAddress == resp.Receive {
		resp.Direction = "out"
	}

	return resp
}

func (t *TransactionEthPending) TransactionEthResponseWithDirectionAndGasPriceView(walletAddress string) *TransactionEthResponse {
	resp := t.TransactionEthResponse()

	if walletAddress == resp.Sender {
		resp.Direction = "out"
	}

	if walletAddress == resp.Receive {
		resp.Direction = "in"
	}

	if walletAddress == resp.Sender && walletAddress == resp.Receive {
		resp.Direction = "out"
	}

	resp.GasPriceView = resp.GasPrice.Div(decimal.New(1, 9)).RoundFloor(2).String() + " Gwei"

	return &resp
}
