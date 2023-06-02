package modeler

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

func confList() []ConfigInterface {
	return []ConfigInterface{&ConfigApp{}}
}

func ConfKeys() []string {
	list := make([]string, 0)
	for _, confObj := range confList() {
		list = append(list, confObj.Key())
	}
	return list
}

func ConfMaps() []map[string]any {
	list := make([]map[string]any, 0)
	for _, confObj := range confList() {
		list = append(list, map[string]any{
			confObj.Key(): confObj,
		})
	}
	return list
}

func GetConfigInterface(str string) ConfigInterface {
	for _, confObj := range confList() {
		if strings.ToLower(confObj.Key()) == strings.ToLower(str) {
			return confObj
		}
	}
	return nil
}

type ConfTest struct {
	IsTestConf string          `json:"isTestConf"`
	IsBool     bool            `json:"isBool"`
	IsNumber   float64         `json:"isNumber"`
	IsDecimal  decimal.Decimal `json:"isDecimal"`
}

func (c *ConfTest) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c *ConfTest) Desc() string {
	return fmt.Sprintf("测试配置")
}

func (c *ConfTest) Key() string {
	return fmt.Sprintf("ConfTest")
}

func (c *ConfTest) FromString(str string) error {
	return json.Unmarshal([]byte(str), &c)
}

type AppBaseConf struct {
	AppName          string `json:"appName"`
	AppApiBaseUrl    string `json:"appApiBaseUrl"`
	AppGetAddressUrl string `json:"appGetAddressUrl"`
	AppWithdrawUrl   string `json:"appWithdrawUrl"`
}

func (c *AppBaseConf) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c *AppBaseConf) Desc() string {
	return fmt.Sprintf("基础配置")
}

func (c *AppBaseConf) Key() string {
	return fmt.Sprintf("AppBaseConf")
}

func (c *AppBaseConf) FromString(str string) error {
	return json.Unmarshal([]byte(str), &c)
}

type Config struct {
	MysqlModel
	Key   string `json:"key" gorm:"column:key;primaryKey"`
	Value string `json:"value" gorm:"column:value;type:json;comment:值"`
	Desc  string `json:"desc" gorm:"column:desc;comment:备注"`
}

func (Config) TableName() string {
	return "config"
}

func (Config) Comment() string {
	return "app总配置"
}
