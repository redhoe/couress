package modeler

import "gorm.io/gorm"

const (
	SLA = "SLA" // service agreement 服务协议
	NDA = "NDA" // user concert form 用户协议
)

type ProtocolsType struct {
	Type   string `json:"type"`
	CnName string `json:"cnName"`
}

var ProtocolsTypes = []ProtocolsType{
	{SLA, "服务协议"},
	{NDA, "用户协议"},
}

type Protocols struct {
	MysqlModel
	MysqlDeleteModel
	Type    string `json:"type" gorm:"type:varchar(10);comment:协议类型;"`
	Lang    string `json:"lang" gorm:"type:varchar(10)"`
	Name    string `json:"name" gorm:"type:varchar(255);comment:协议标题;"`
	Content string `json:"content" gorm:"type:text;comment:协议内容;"`
	Show    bool   `json:"show" gorm:"comment:协议类型;"`
	Sort    int    `json:"sort" gorm:"type:int(4);default:9999"` // 排序，由小到大
}

func (*Protocols) TableName() string {
	return "protocols"
}

func (*Protocols) Comment() string {
	return "协议"
}

func NewProtocols() *Protocols {
	return &Protocols{}
}

func (t *Protocols) SLA(db *gorm.DB, lang string) uint {
	db.Model(&Protocols{}).Where("lang", lang).Where("type", SLA).Find(&t)
	return t.Id
}

func (t *Protocols) NDA(db *gorm.DB, lang string) uint {
	db.Model(&Protocols{}).Where("lang", lang).Where("type", NDA).Find(&t)
	return t.Id
}
