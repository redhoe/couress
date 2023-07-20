package modeler

type WalletCiphertext struct {
	WalletIdentityId uint   `json:"walletIdentityId" gorm:"comment:id"`
	Ciphertext       string `json:"ciphertext" gorm:"type:text;comment:ciphertext"`
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
