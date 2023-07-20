package modeler

type VersionType string

const (
	Android VersionType = "android"
	Ios     VersionType = "ios"
	Gkey    VersionType = "gkey"
)

type Version struct {
	MysqlModel
	Type    VersionType `json:"type" gorm:"type:varchar(10)"`
	Uri     string      `json:"uri" gorm:"uri"`
	Version string      `json:"version" gorm:"version"`
	Enable  bool        `json:"enable" gorm:"enable"`
	Force   bool        `json:"force" gorm:"force"`
	Icon    string      `json:"icon" gorm:"type:varchar(200);comment:图标"`
	Name    string      `json:"name" gorm:"type:varchar(200);comment:名称"`
	Desc    string      `json:"desc" gorm:"type:varchar(200);comment:描述"`
}

func (*Version) TableName() string {
	return "version"
}

func (*Version) Comment() string {
	return "版本信息"
}

func NewVersion() *Version {
	return &Version{}
}

type VersionDocument struct {
	MysqlModel
	MysqlDeleteModel
	VersionId uint     `json:"version_id" gorm:"version_id"`
	Version   *Version `json:"version" gorm:"foreignKey:id;references:VersionId"`
	Content   string   `json:"content" gorm:"type:text"`
	Lang      string   `json:"lang" gorm:"lang"`
	Show      bool     `json:"show" gorm:"show"`
}

func (*VersionDocument) TableName() string {
	return "version_document"
}

func (*VersionDocument) Comment() string {
	return "版本信息文档"
}

func NewVersionDocument() *VersionDocument {
	return &VersionDocument{}
}
