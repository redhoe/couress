package confer

import (
	"github.com/spf13/viper"
)

var AppConfServer Server

type Server struct {
	JWT        JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System     System  `mapstructure:"system" json:"system" yaml:"system"`
	Redis      Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql      Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	MysqlSlave Mysql   `mapstructure:"mysql-slave" json:"mysql-slave" yaml:"mysql-slave"`
	Local      Local   `mapstructure:"local" json:"local" yaml:"local"`
	Oss        Oss     `mapstructure:"oss" json:"oss" yaml:"oss"`
	Zap        Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Captcha    Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}

func DefaultConfInit(v *viper.Viper) {
	v.SetDefault("jwt", JWT{
		SigningKey:  "signingKey",
		ExpiresTime: "100h",
		BufferTime:  "3s",
		Issuer:      "hoer",
	})
	v.SetDefault("system", System{
		Env:           "dev",
		Addr:          1203,
		DbType:        "mysql",
		OssType:       "local",
		UseMultipoint: false,
		UseRedis:      false,
		LimitCountIP:  100,
		LimitTimeIP:   3,
		RouterPrefix:  "",
	})
	v.SetDefault("redis", Redis{
		Host:  "127.0.0.1",
		Port:  6379,
		Auth:  "",
		Index: 5,
	})
	v.SetDefault("mysql", Mysql{
		Host:   "127.0.0.1",
		Port:   "3306",
		User:   "root",
		Secret: "3nzd100W",
		Name:   "wallet_app_new",
	})
	v.SetDefault("mysqlSlave", Mysql{
		Host:   "127.0.0.1",
		Port:   "3306",
		User:   "root",
		Secret: "3nzd100W",
		Name:   "wallet_app_new",
	})
	v.SetDefault("local", Local{
		Path:      "http://127.0.0.1:1113",
		StorePath: "static/uploads",
		MaxSize:   10,
		Types:     "jpg,png,jpeg,csv",
	})
	v.SetDefault("oss", Oss{
		Endpoint:        "oss",
		AccessKeyId:     "access-key-id",
		AccessKeySecret: "access-key-secret",
		BucketName:      "bucket name",
		BucketUrl:       "url",
		BasePath:        "bucket01",
	})
	v.SetDefault("zap", Zap{
		Level:         "info",
		Prefix:        "app--",
		Format:        "",
		Director:      "logs",
		EncodeLevel:   "CapitalColorLevelEncoder",
		StacktraceKey: "error",
		MaxAge:        2,
		ShowLine:      true,
		LogInConsole:  true,
	})
	v.SetDefault("captcha", Captcha{
		KeyLong:            4,
		ImgWidth:           100,
		ImgHeight:          60,
		OpenCaptcha:        3,
		OpenCaptchaTimeOut: 30,
	})
}
