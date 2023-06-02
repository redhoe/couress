package cache

import (
	"fmt"
	"sync"
)

type LocalCache struct {
	list sync.Map
}

func NewLocalCache() *LocalCache {
	return &LocalCache{sync.Map{}}
}

func (u *LocalCache) key(category, key string) string {
	return fmt.Sprintf("%s-%s", category, key)
}

func (u *LocalCache) Add(category, key string, data any) {
	u.list.Store(u.key(category, key), data)
}

func (u *LocalCache) Find(category, key string) (any, bool) {
	data, ok := u.list.Load(u.key(category, key))
	if !ok {
		return nil, false
	}
	return data, ok
}

func (u *LocalCache) Finds(category string, keys ...string) (any, bool) {
	for _, arg := range keys {
		data, ok := u.list.Load(u.key(category, arg))
		if ok {
			return data, false
		}
	}
	return nil, false
}

const (
	GbCACHECategoryWalletAddress       = "WalletAddress"       // 监听地址本地缓存key
	GbCACHECategoryWalletAddressLastId = "WalletAddressLastId" // 监听地址最后的Id 用于更新写入
)

// 地址管理

func (u *LocalCache) AddAddress(key string, data any) {
	u.list.Store(u.key(GbCACHECategoryWalletAddress, key), data)
}

func (u *LocalCache) FindAddress(key string) (any, bool) {
	data, ok := u.list.Load(u.key(GbCACHECategoryWalletAddress, key))
	if !ok {
		return nil, false
	}
	return data, ok
}

func (u *LocalCache) FindAddressMore(keys ...string) (any, bool) {
	for _, arg := range keys {
		data, ok := u.list.Load(u.key(GbCACHECategoryWalletAddress, arg))
		if ok {
			return data, true
		}
	}
	return nil, false

}

func (u *LocalCache) SetAddressLastId(lastId uint) {
	u.list.Store(u.key(GbCACHECategoryWalletAddressLastId, ""), lastId)
}

func (u *LocalCache) GetAddressLastId() (uint, bool) {
	data, ok := u.list.Load(u.key(GbCACHECategoryWalletAddressLastId, ""))
	if !ok {
		return 0, false
	}
	return data.(uint), ok
}
