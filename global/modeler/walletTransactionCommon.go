package modeler

import "github.com/shopspring/decimal"

// Transaction 用于各链交易记录保存的基础信息
type Transaction struct {
	ChainId       uint            `json:"chain_id" gorm:"index;comment:链ID:chain"`
	CoinId        uint            `json:"coin_id" gorm:"index;comment:币种ID:coin"`
	Symbol        string          `json:"Symbol" gorm:"type:varchar(20);comment:币种名"`
	SymbolDecimal int32           `json:"SymbolDecimal" gorm:"comment:币种精度"`
	BlockNumber   int64           `json:"block_number" gorm:"index;comment:区块号"`
	BlockTime     int64           `json:"block_time" gorm:"comment:区块扫描时间"`
	Hash          string          `json:"hash" gorm:"type:varchar(100);index"`
	Sender        string          `json:"sender" gorm:"type:varchar(50);index"`
	FoundSender   bool            `json:"found_sender" gorm:"default:false"`
	Receive       string          `json:"receive" gorm:"type:varchar(50);index"`
	FoundReceive  bool            `json:"found_receive" gorm:"default:false"`
	Amount        decimal.Decimal `json:"amount" gorm:"type:decimal(50,18);comment:交易金额 单位decimal"`
	AmountInt     decimal.Decimal `json:"amount_int" gorm:"type:decimal(50);comment:交易金额:wei"`
}

func (tx *Transaction) TransactionResponse() TransactionResponse {
	return TransactionResponse{
		BlockNumber: tx.BlockNumber,
		BlockTime:   tx.BlockTime,
		Hash:        tx.Hash,
		Sender:      tx.Sender,
		Receive:     tx.Receive,
		Amount:      tx.Amount,
	}
}
