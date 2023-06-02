package couress

import (
	"github.com/redhoe/couress/global"
	"github.com/redhoe/couress/global/core/cache"
	"github.com/redhoe/couress/global/core/dataer"
	"github.com/redhoe/couress/global/core/loger"
)

func SetGlobals() {
	// 生成与读取配置文件
	global.GbVP = global.Viper()
	loger.LogerInit() // 日志引擎初始化
	// 初始化全局变化
	global.GbDB = dataer.Living()
	global.GbREDIS = cache.GetCacheEngine()
	global.GbLOG = loger.NewLogger(loger.App)
	global.GbSLOG = global.GbLOG.Sugar()
}
