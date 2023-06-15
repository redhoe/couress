package modeler

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

// tips: 为前端准备的配置说明需要配置的内容 需要给前端用需要把结构体放入清单
func confList() []ConfigInterface {
	return []ConfigInterface{NewConfigApp(), NewConfTest()}
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
			confObj.Key(): reflectModelToMap(confObj),
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
	IsTestConf string          `json:"isTestConf" gorm:"comment:string类型测试配置"`
	IsBool     bool            `json:"isBool" gorm:"comment:bool类型"`
	IsNumber   float64         `json:"isNumber" gorm:"comment:number类型"`
	IsDecimal  decimal.Decimal `json:"isDecimal" gorm:"comment:decimal类型"`
}

func NewConfTest() *ConfTest {
	return &ConfTest{}
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

type Config struct {
	MysqlModel
	Key   string `json:"key" gorm:"column:key;primaryKey"`
	Value string `json:"value" gorm:"column:value;type:json;comment:值"`
	Desc  string `json:"desc" gorm:"column:desc;comment:备注"`
}

func (*Config) TableName() string {
	return "config"
}

func (*Config) Comment() string {
	return "app总配置"
}

func NewConfig() *Config {
	return &Config{}
}
