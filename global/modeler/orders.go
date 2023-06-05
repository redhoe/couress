package modeler

import (
	"github.com/shopspring/decimal"
)

type Orders struct {
	MysqlModel
	UserId uint            `json:"userId" gorm:"column:user_id;type:int;comment:用户Id"`
	Name   string          `json:"name" gorm:"column:name;type:varchar(100);comment:名称"`
	Price  decimal.Decimal `json:"price" gorm:"column:price;type:decimal(28,18);comment:单价"`
	Status bool            `json:"status" gorm:"column:status;comment:isOk;default:false"`
	Desc   string          `json:"desc" gorm:"column:desc;type:varchar(200);comment:描述;"`
	NoEdit string          `json:"-" gorm:"column:no_edit;type:varchar(100);comment:改字段不准修改"`
	Nums   int64           `json:"-" gorm:"column:nums;type:int;comment:nums改字段不准修改"`
}

func (*Orders) TableName() string {
	return "app_orders"
}

func (*Orders) Comment() string {
	return "订单表"
}

func (o *Orders) GetNoEdit() string {
	return o.NoEdit
}

func (o *Orders) SetNoEdit(str string) {
	o.NoEdit = str
	return
}

func (o *Orders) GetNums() int64 {
	return o.Nums
}

func (o *Orders) SetNums(num int64) {
	o.Nums = num
	return
}

func NewOrders() *Orders {
	return &Orders{}
}

func (*Orders) KeyMap() map[string]string {
	return reflectSignStruct(Orders{})
}
