package global

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/redhoe/couress/global/core/cache"
	"github.com/redhoe/couress/global/core/confer"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GbCONFIG confer.Server
	GbDB     *gorm.DB
	GbREDIS  *redis.Client // redis缓存
	GbVP     *viper.Viper
	GbLOG    *zap.Logger
	GbSLOG   *zap.SugaredLogger // 日志糖
	GbCACHE  *cache.LocalCache  // 本地缓存
	lock     sync.RWMutex
)

const (
	ConfigDefaultFile = "config.json"
)

func init() {
	GbCACHE = cache.NewLocalCache()
}

// Viper

func Viper(path ...string) *viper.Viper {
	var configPath string
	if len(path) == 0 {
		flag.StringVar(&configPath, "c", "", "choose config file.")
		flag.Parse()
		if configPath == "" { // 判断命令行参数是否为
			configPath = ConfigDefaultFile
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", configPath)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		configPath = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", configPath)
	}
	v := viper.New()
	v.SetConfigFile(configPath)
	viper.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		// 读取配置文件失败 则初始化内容到confer在指定路径创建改配置文件
		confer.DefaultConfInit(v)
		if err := v.WriteConfig(); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
			return nil
		}
	}

	v.WatchConfig() // 开启观察者模式
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&GbCONFIG); err != nil {
			fmt.Println(err)
		}
		if err = v.Unmarshal(&confer.AppConfServer); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&GbCONFIG); err != nil {
		fmt.Println(err)
	}
	if err = v.Unmarshal(&confer.AppConfServer); err != nil {
		fmt.Println(err)
	}
	return v
}
