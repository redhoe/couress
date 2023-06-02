package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/redhoe/couress/global/core/confer"
	"time"
)

var cacheEngineLiving *redis.Client

func cacheEngineInit() error {
	addr := fmt.Sprintf("%s:%d", confer.AppConfServer.Redis.Host, confer.AppConfServer.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: confer.AppConfServer.Redis.Auth,
		DB:       confer.AppConfServer.Redis.Index,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return err
	}
	cacheEngineLiving = client
	return nil
}

func GetCacheEngine() *redis.Client {
	if cacheEngineLiving == nil {
		if err := cacheEngineInit(); err != nil {
			panic(err)
		}
	}
	return cacheEngineLiving
}

// CheckRepeatCache 去重检查
func CheckRepeatCache(serviceName string, userId uint, timeSt time.Duration) bool {
	if cacheEngineLiving == nil {
		if err := cacheEngineInit(); err != nil {
			panic(err)
		}
	}
	serviceNameCache := serviceKeyCheck(serviceName, userId)
	nxBool, err := cacheEngineLiving.SetNX(cacheEngineLiving.Context(), serviceNameCache, userId, timeSt).Result()
	// 是否需要返回缓存剩余时间
	if err != nil {
		return false
	}
	_, err = cacheEngineLiving.TTL(cacheEngineLiving.Context(), serviceNameCache).Result()
	if err != nil {
		return false
	}
	if !nxBool {
		return false
	}
	return true
}

func serviceKeyCheck(a string, u uint) string {
	return fmt.Sprintf("CRCACHE:%s:%d", a, u)
}
