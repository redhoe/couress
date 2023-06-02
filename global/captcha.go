package global

import (
	"github.com/demdxx/gocast"
	"time"
)

var CaptchaStore = NewDefaultRedisStore()

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
}

func (rs *RedisStore) Set(id string, value string) error {
	err := GbREDIS.Set(GbREDIS.Context(), rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := GbREDIS.Get(GbREDIS.Context(), key).Result()
	if err != nil {
		return ""
	}
	if clear {
		err := GbREDIS.Del(GbREDIS.Context(), key).Err()
		if err != nil {
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}

// 增加错误次数

func (rs *RedisStore) AddTime(id string) error {
	times, _ := GbREDIS.Get(GbREDIS.Context(), rs.PreKey+id).Result()
	timesInt := gocast.ToInt(times)
	return GbREDIS.Set(GbREDIS.Context(), rs.PreKey+id, timesInt+1, rs.Expiration).Err()
}

func (rs *RedisStore) GetErrorTime(id string) (string, error) {
	return GbREDIS.Get(GbREDIS.Context(), rs.PreKey+id).Result()
}
