package modeler

type VersionGkeyModel struct {
	MysqlModel
	VersionId uint     `json:"version_id" gorm:";httpCommon:ciphertext"`
	Version   *Version `json:"version" gorm:"foreignKey:id;references:VersionId"`
	Icon      string   `json:"icon" gorm:";httpCommon:图标"`
	Name      string   `json:"name" gorm:";httpCommon:名称"`
	Desc      string   `json:"desc" gorm:";httpCommon:描述"`
	VersionNo string   `json:"version_no" gorm:";httpCommon:版本号"`
	Top       bool     `json:"top" gorm:";httpCommon:是否热门"`
	Sort      int      `json:"sort" gorm:"default:9999;httpCommon:排序"`
	Enable    bool     `json:"enable" gorm:";httpCommon:是否有效"`
}

func (*VersionGkeyModel) TableName() string {
	return "version_gkey_model"
}

func (*VersionGkeyModel) Comment() string {
	return "硬件版本信息"
}

func NewVersionGkeyModel() *VersionGkeyModel {
	return &VersionGkeyModel{}
}
