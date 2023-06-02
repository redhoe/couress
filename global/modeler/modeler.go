package modeler

import (
	"fmt"
	"gorm.io/gorm"
	"math"
	"reflect"
	"strings"
	"time"
)

const (
	EthCoinBalanceDecimal int32 = 8
	PriceDecimal                = 2
)

// 分页器

type Paging struct {
	Page      int   `json:"page"`      //当前页
	PageSize  int   `json:"pageSize" ` //每页条数
	Total     int64 `json:"total"`     //总条数
	PageCount int   `json:"pageCount"` //总页数
	StartNums int   `json:"startNums"` //起始条数
}

func NewPaging() *Paging {
	return &Paging{}
}

func (p *Paging) GetPages() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	p.StartNums = p.PageSize * (p.Page - 1)
	pageCount := math.Ceil(float64(p.Total) / float64(p.PageSize))
	p.PageCount = int(pageCount)
}

type MysqlModel struct {
	MysqlIdModel
	MysqlTimeModel
}

func (a *MysqlModel) CreateTime() time.Time {
	return a.CreatedAt
}

func (a *MysqlModel) CheckRowIsExist(db *gorm.DB, column, value interface{}) bool {
	if err := db.Where(fmt.Sprintf("%s = ?", column), value).
		Find(&a).Error; err != nil {
		return false
	}
	if a.Id != 0 {
		return true
	}
	return false
}

// 获取模型中字段与描述信息
func reflectSignStruct(i any) map[string]string {
	modelType := reflect.TypeOf(i)
	ky := map[string]string{}
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		if field.Name == "MysqlModel" {
			ky1 := reflectSignStruct(MysqlIdModel{})
			ky = mapAdd(ky, ky1)
			ky2 := reflectSignStruct(MysqlTimeModel{})
			ky = mapAdd(ky, ky2)
			continue
		}
		if field.Name == "MysqlIdModel" {
			ky1 := reflectSignStruct(MysqlIdModel{})
			ky = mapAdd(ky, ky1)
			continue
		}
		if field.Name == "MysqlTimeModel" {
			ky1 := reflectSignStruct(MysqlTimeModel{})
			ky = mapAdd(ky, ky1)
			continue
		}
		fieldJson := field.Tag.Get("json")
		if fieldJson == "-" {
			continue
		}
		fieldJsons := strings.Split(fieldJson, ",") // 处理这种情况:fieldName,omitempty
		ky[fieldJsons[0]] = ""
		fieldGorm := field.Tag.Get("gorm")
		gormTags := strings.Split(fieldGorm, ";")
		for _, tag := range gormTags {
			tag_ := strings.Split(tag, ":")
			if len(tag_) == 2 && strings.ToLower(tag_[0]) == "comment" {
				ky[fieldJson] = tag_[1]
			}
		}
	}

	return ky
}

func mapAdd(a map[string]string, b map[string]string) map[string]string {
	c := map[string]string{}
	for k, v := range b {
		c[k] = v
	}
	for k, v := range a {
		c[k] = v
	}
	return c
}

type MysqlIdModel struct {
	Id uint `json:"id,omitempty" gorm:"primarykey;comment:id"`
}

func (a MysqlModel) IdNo() uint {
	return a.Id
}

type MysqlTimeModel struct {
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:更新时间"`
}

type MysqlDeleteModel struct {
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
}
