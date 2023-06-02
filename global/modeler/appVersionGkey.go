package modeler

type VersionGkeyModel struct {
	MysqlModel
	VersionId uint     `json:"version_id" gorm:";common:ciphertext"`
	Version   *Version `json:"version" gorm:"foreignKey:id;references:VersionId"`
	Icon      string   `json:"icon" gorm:";common:图标"`
	Name      string   `json:"name" gorm:";common:名称"`
	Desc      string   `json:"desc" gorm:";common:描述"`
	VersionNo string   `json:"version_no" gorm:";common:版本号"`
	Top       bool     `json:"top" gorm:";common:是否热门"`
	Sort      int      `json:"sort" gorm:"default:9999;common:排序"`
	Enable    bool     `json:"enable" gorm:";common:是否有效"`
}

func (VersionGkeyModel) TableName() string {
	return "version_gkey_model"
}

func (VersionGkeyModel) Comment() string {
	return "硬件版本信息"
}
