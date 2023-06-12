package etherModel

import (
	"github.com/redhoe/couress/global/modeler"
	"github.com/shopspring/decimal"
	"math/big"
)

type TransactionEthResponse struct {
	modeler.TransactionResponse
	ID                    int64                 `json:"-"`
	Status                int                   `json:"status"`
	Direction             string                `json:"direction"`
	Message               *string               `json:"message"`
	GasLimit              int64                 `json:"gas_limit"`
	GasPrice              decimal.Decimal       `json:"gas_price"`
	GasAmount             decimal.Decimal       `json:"gas_amount"`
	GasAmountUsd          decimal.Decimal       `json:"-"`
	GasAmountCurrencyInfo *modeler.CurrencyInfo `json:"gas_amount_currency_info,omitempty"`
	GasPriceView          string                `json:"gas_price_view,omitempty"`
	WebUrl                string                `json:"web_url"`
	Nonce                 int64                 `json:"nonce"`
	IsPending             bool                  `json:"is_pending"`
	IsRead                *bool                 `json:"is_read,omitempty"`
}

// SetDirection 设置交易方向
func (t *TransactionEthResponse) SetDirection(walletAddress string) {
	if walletAddress == t.Sender {
		t.Direction = "out"
	}
	if walletAddress == t.Receive {
		t.Direction = "in"
	}
	if walletAddress == t.Sender && walletAddress == t.Receive {
		t.Direction = "out"
	}
}

// SetGasAmountUsd 设置 gas amount usd price
func (t *TransactionEthResponse) SetGasAmountUsd(usdPrice decimal.Decimal) {
	t.GasAmountUsd = usdPrice
}

// SetGasAmountCurrencyInfo 设置 gas amount 法币详情
func (t *TransactionEthResponse) SetGasAmountCurrencyInfo(currencyInfo *modeler.CurrencyInfo) {
	t.GasAmountCurrencyInfo = currencyInfo
}

type BuildTransferByEthereum struct {
	ChainId  int             `json:"chain_id"`
	Nonce    uint64          `json:"nonce"`
	GasLimit uint64          `json:"gas_limit"`
	GasPrice decimal.Decimal `json:"gas_price"`
	To       string          `json:"to"`
	Value    *big.Int        `json:"value"`
	Data     string          `json:"data"`
}

type BuildTransferByEthereumResponse struct {
	BuildTransfer  BuildTransferByEthereum `json:"build_transfer"`
	ReceiveAddress string                  `json:"receive_address"`
	ChainNetAlias  string                  `json:"chain_net_alias"`
}

type ReplaceTransferByEthereumResponse struct {
	*BuildTransferByEthereumResponse
	Original *TransactionEthResponse `json:"original"`
	Replace  *TransactionEthResponse `json:"replace"`
	Coin     *modeler.CoinInfo       `json:"coin"`
}
