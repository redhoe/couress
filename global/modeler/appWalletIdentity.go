package modeler

import (
	"gorm.io/gorm"
)

type WalletIdentity struct {
	MysqlModel
	Identity string `json:"identity" gorm:"unique;common:身份ID"`
}

func (WalletIdentity) TableName() string {
	return "wallet_identity"
}

func (WalletIdentity) Comment() string {
	return "身份钱包Id"
}

func GetWalletIdentity(db *gorm.DB, identity string) (*WalletIdentity, error) {
	walletIdentity := &WalletIdentity{}
	err := db.Model(&WalletIdentity{}).Where("`identity` = ?", identity).First(&walletIdentity).Error
	if err != nil && gorm.ErrRecordNotFound != err {
		return nil, err
	}
	if gorm.ErrRecordNotFound == err {
		walletIdentity.Identity = identity
		if err = db.Create(&walletIdentity).Error; err != nil {
			return nil, err
		}
	}
	return walletIdentity, nil
}
