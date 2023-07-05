package modeler

import "encoding/json"

type ConfigAppLink struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

type ConfigApp struct {
	ApiDomain                       string          `json:"api_domain" gorm:"comment:api域名"`                                   // api域名
	FileDomain                      string          `json:"file_domain" gorm:"comment:文件域名"`                                   // 文件域名
	Website                         string          `json:"website" gorm:"comment:官网"`                                         // 官网
	UserAgreementByDocumentTagId    int             `json:"user_agreement_by_document_tag_id" gorm:"comment:用户协议文档 tag ID"`    // 用户协议文档 tag ID
	ServiceAgreementByDocumentTagId int             `json:"service_agreement_by_document_tag_id" gorm:"comment:服务协议文档 tag ID"` // 用户协议文档 tag ID
	UsdGkeyByDocumentTagId          string          `json:"usd_gkey_by_document_tag_id" gorm:"comment:Gkey使用教程Url"`            // Gkey使用教程
	ConnGkeyByDocumentTagId         string          `json:"conn_gkey_by_document_tag_id" gorm:"comment:链接Gkey教程"`              // 链接Gkey教程
	BuyGkeyByDocumentTagId          string          `json:"buy_gkey_by_document_tag_id" gorm:"comment:购买和了解Gkey"`              // 购买和了解Gkey
	ShopUrl                         string          `json:"shop_url" gorm:"comment:商城url"`                                     // 商城url
	Links                           []ConfigAppLink `json:"links,omitempty" gorm:"comment:相关链接"`                               // 相关链接
	RsaPubkey                       string          `json:"rsa_pubkey,omitempty" gorm:"comment:rsa公钥"`                         // rsa公钥
	PushSwitch                      bool            `json:"push_switch,omitempty" gorm:"comment:上传敏感信息开关"`                     // 上传敏感信息开关
	DownUrl                         string          `json:"down_url,omitempty" gorm:"comment:下载地址"`                            // 下载地址
	HelpUrl                         string          `json:"help_url" gorm:"comment:帮助中心地址"`
}

func NewConfigApp() *ConfigApp {
	return &ConfigApp{}
}

func (a *ConfigApp) Key() string {
	return "app"
}

func (a *ConfigApp) Desc() string {
	return "app配置"
}

func (a *ConfigApp) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}

func (a *ConfigApp) FromString(str string) error {
	return json.Unmarshal([]byte(str), a)
}

var DefaultConfigApp = &ConfigApp{
	ApiDomain:                    "http://127.0.0.1:1232",
	FileDomain:                   "http://127.0.0.1:1232",
	Website:                      "http://127.0.0.1:1232",
	UserAgreementByDocumentTagId: 0,
	Links:                        nil,
}
