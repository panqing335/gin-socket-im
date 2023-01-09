package cache

import (
	"github.com/spf13/viper"
	"sync"
	"temp/app/common/redis"
	"time"
)

var Pool *sync.Pool

func init() {
	Pool = &sync.Pool{
		New: func() interface{} {
			ttl := time.Duration(viper.GetInt("modelCache.ttl")) * time.Second
			emptyTTL := time.Duration(viper.GetInt("modelCache.emptyTTL")) * time.Second
			return NewSimpleCache(redis.NewOperation(), ttl, NewCrossPolicy("", emptyTTL))
		},
	}
}

func Cache() *SimpleCache {
	return Pool.Get().(*SimpleCache)
}

func ReleaseCache(cache *SimpleCache) {
	Pool.Put(cache)
}
