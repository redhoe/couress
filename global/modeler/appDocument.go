package modeler

type DocumentTag struct {
	MysqlModel
	MysqlDeleteModel
	Name     string `json:"name"  gorm:""`
	Image    string `json:"image"  gorm:""`
	ParentId uint   `json:"parent_id" gorm:"default:0"`
	Lang     string `json:"lang"  gorm:""`
	Show     bool   `json:"show"  gorm:""`
	Sort     int64  `json:"sort" gorm:"default:9999"`
}

func (DocumentTag) TableName() string {
	return "document_tag"
}

func (DocumentTag) Comment() string {
	return "文档标签"
}

type DocumentBanner struct {
	MysqlModel
	MysqlDeleteModel
	Name       string `json:"name" gorm:""`
	Image      string `json:"image" gorm:""`
	Lang       string `json:"lang" gorm:""`
	Sort       int64  `json:"sort" gorm:"default:9999"`
	Show       bool   `json:"show" gorm:""`
	DocumentId int64  `json:"document_id" gorm:""`
}

func (DocumentBanner) TableName() string {
	return "document_banner"
}

func (DocumentBanner) Comment() string {
	return "轮播图"
}

type Document struct {
	MysqlModel
	MysqlDeleteModel
	Name    string       `json:"name" gorm:""`
	TagId   uint         `json:"tag_id" gorm:""`
	Tag     *DocumentTag `json:"tag" gorm:"foreignKey:tag_id;references:id;"`
	Content string       `json:"content" gorm:"type:text"`
	Lang    string       `json:"lang" gorm:"lang"`
	Sort    int64        `json:"sort" gorm:"default:9999"`
	Show    bool         `json:"show" gorm:""`
	Hot     bool         `json:"hot" gorm:""`
}

func (Document) TableName() string {
	return "document"
}

func (Document) Comment() string {
	return "文档"
}
