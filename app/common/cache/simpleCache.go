package cache

import (
	"fmt"
	Redis "temp/app/common/redis"
	"time"
)

type DbGetFunc func() Redis.ResultType

type SimpleCache struct {
	Operation *Redis.Operation
	Expire    time.Duration
	DbGetter  DbGetFunc
	Policy    CachePolicy
}

func NewSimpleCache(operation *Redis.Operation, expire time.Duration, policy CachePolicy) *SimpleCache {
	policy.SetOperation(operation)
	return &SimpleCache{Operation: operation, Expire: expire, Policy: policy}
}

func (s *SimpleCache) SetCache(key string, value any) {
	s.Operation.Set(key, value, Redis.WithExpire(s.Expire)).Unwrap()
}

func (s *SimpleCache) HMSetCache(key string, value ...any) {
	s.Operation.HMSet(key, value...).Unwrap()
	s.Operation.Expire(key, s.Expire)
}

func (s *SimpleCache) ExpireCache(key string, ttl time.Duration) {
	s.Operation.Expire(key, ttl).Unwrap()
}

func (s *SimpleCache) DelCache(key string) {
	s.Operation.Del(key).Unwrap()
}

func (s *SimpleCache) GetCache(key string) (ret map[string]string) {
	if s.Policy != nil {
		s.Policy.Before(key)
	}
	res := s.Operation.HGetAll(key).UnwrapOrElse(s.DbGetter)
	fmt.Println("res:", res)

	a, ok := res.(map[string]any)
	if ok && len(a) != 0 {
		fmt.Printf("ret查db了 %v\n", a)
		s.HMSetCache(key, res)
	}
	m := make(map[string]any)
	m = map[string]any{"id": -1}

	if ok && len(a) == 0 {
		fmt.Println("ret查db了, 值还是空")
		fmt.Println("ret查db了, 值还是空", m)
		s.Policy.IfNil(key, m)
	}

	i, oki := res.(map[string]string)
	if oki && i["id"] == "-1" {
		fmt.Println("ret查缓存了, id值为：", i["id"])
		s.Policy.IfNil(key, m)
	}
	if oki && i["id"] != "-1" {
		fmt.Println("ret查缓存了, 更新ttl")
		s.HMSetCache(key, res)
	}
	ret = s.Operation.HGetAll(key).Unwrap().(map[string]string)

	return ret
}
