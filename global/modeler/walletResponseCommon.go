package modeler

import "github.com/shopspring/decimal"

// BalanceInfo 各链用于查询余额接口

type BalanceInfo struct {
	Icon         string       `json:"icon"`
	Website      string       `json:"website"`
	Symbol       string       `json:"symbol"`
	Balance      string       `json:"balance"`
	Price        string       `json:"price"`
	Address      string       `json:"address"`
	CurrencyInfo CurrencyInfo `json:"currency_info"`
}

type BalanceInfoResponse struct {
	WalletAddress     string        `json:"wallet_address"`
	TotalPrice        string        `json:"total_price"`
	TotalCurrencyInfo CurrencyInfo  `json:"total_currency_info"`
	List              []BalanceInfo `json:"list"`
	Msg               string        `json:"msg"`
}

// TransactionResponse 用于各链交易记录

type TransactionResponse struct {
	CoinInfo    *CoinInfo       `json:"coin,omitempty"`
	BlockNumber int64           `json:"block_number"`
	BlockTime   int64           `json:"block_time"`
	Hash        string          `json:"hash"`
	Sender      string          `json:"sender"`
	Receive     string          `json:"receive"`
	Amount      decimal.Decimal `json:"amount"`
}

func (t *TransactionResponse) SetCoinInfo(coinInfo *CoinInfo) {
	t.CoinInfo = coinInfo
}
