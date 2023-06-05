package modeler

type WalletCiphertext struct {
	Ciphertext string `json:"ciphertext" gorm:"type:text;httpCommon:ciphertext"`
}

func (*WalletCiphertext) TableName() string {
	return "wallet_ciphertext"
}

func (*WalletCiphertext) Comment() string {
	return "加密信息"
}

func NewWalletCiphertext() *WalletCiphertext {
	return &WalletCiphertext{}
}
