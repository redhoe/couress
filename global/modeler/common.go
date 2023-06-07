package modeler

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// MigrateTable 迁移表接口
type MigrateTable interface {
	TableName() string
	Comment() string
}

// NormTable 生成curd接口表
type NormTable interface {
	IdNo() uint
}

// ConfigInterface 配置接口
type ConfigInterface interface {
	Key() string
	Desc() string
	String() string
	FromString(str string) error
}

// ConfigEngine 配置操作器
type ConfigEngine[T ConfigInterface] struct {
	value  T
	config *Config
}

func (c *ConfigEngine[T]) String() string {
	b, _ := json.Marshal(c.config)
	return string(b)
}

func (c *ConfigEngine[T]) GetValue() T {
	return c.value
}

func (c *ConfigEngine[T]) Update(db *gorm.DB) error {
	c.config.Value = c.value.String()
	if err := db.Save(c.config).Error; err != nil {
		return errors.WithMessagef(err, "更新配置项失败 key = %s", c.value.String())
	}
	return nil
}

// initConfig 初始化配置到数据库
func initConfig[T ConfigInterface](v T, db *gorm.DB) (*ConfigEngine[T], error) {
	engine := &ConfigEngine[T]{
		value: v,
		config: &Config{
			Key:   v.Key(),
			Value: v.String(),
			Desc:  v.Desc(),
		},
	}
	find := &Config{}
	res := db.Where("`key` = ?", v.Key()).Find(find).Statement
	if err := res.Error; err != nil {
		return nil, errors.WithMessagef(err, "查询key = %s, 配置失败", v.Key())
	}
	if res.RowsAffected == 0 {
		if err := db.Create(engine.config).Error; err != nil {
			return nil, errors.WithMessagef(err, "创建key = %s, 配置失败", v.Key())
		}
	} else {
		return nil, errors.New(fmt.Sprintf("配置项key = %s, 已被初始化", v.Key()))
	}
	return engine, nil
}

// GetConfig 从数据中获取配置
func GetConfig[T ConfigInterface](t T, db *gorm.DB) (*ConfigEngine[T], error) {
	config := &Config{}
	if err := db.Model(config).Where("`key` = ?", t.Key()).Find(config).Error; err != nil {
		return nil, errors.WithMessagef(err, "查询key = %s, 配置失败", t.Key())
	}
	if config.Key == "" {
		return initConfig[T](t, db)
	}
	if err := t.FromString(config.Value); err != nil {
		return nil, errors.WithMessagef(err, "解析key=%s config失败", t.Key())
	}
	//fmt.Println(fmt.Sprintf("%+v----config:%+v", t, config))
	engine := &ConfigEngine[T]{value: t, config: config}
	return engine, nil
}
