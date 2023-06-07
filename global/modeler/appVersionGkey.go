package modeler

type VersionGkeyModel struct {
	MysqlModel
	VersionId uint     `json:"version_id" gorm:";comment:ciphertext"`
	Version   *Version `json:"version" gorm:"foreignKey:id;references:VersionId"`
	Icon      string   `json:"icon" gorm:";comment:图标"`
	Name      string   `json:"name" gorm:";comment:名称"`
	Desc      string   `json:"desc" gorm:";comment:描述"`
	VersionNo string   `json:"version_no" gorm:";comment:版本号"`
	Top       bool     `json:"top" gorm:";comment:是否热门"`
	Sort      int      `json:"sort" gorm:"default:9999;comment:排序"`
	Enable    bool     `json:"enable" gorm:";comment:是否有效"`
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
