package etherModel

import "github.com/redhoe/couress/global/modeler"

type WalletMessageType string

const (
	WalletMessageType_Transaction WalletMessageType = "transaction"
)

type WalletMessage struct {
	modeler.MysqlModel
	modeler.MysqlDeleteModel
	ChainId       uint              `json:"chain_id" gorm:"index"`
	IdentityId    uint              `json:"identity_id" gorm:"index"`
	WalletChainId uint              `json:"wallet_chain_id" gorm:"index"`
	Type          WalletMessageType `json:"type" gorm:""`
	Table         string            `json:"table" gorm:""`
	TableId       uint              `json:"table_id" gorm:""`
	IsRead        bool              `json:"is_read" gorm:""`
	Message       *string           `json:"message" gorm:""`
}

func (*WalletMessage) TableName() string {
	return "wallet_message"
}

func NewWalletMessage() *WalletMessage {
	return &WalletMessage{}
}

func (*WalletMessage) Comment() string {
	return "交易通知"
}
