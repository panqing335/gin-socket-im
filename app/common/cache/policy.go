package cache

import (
	"regexp"
	"temp/app/common/redis"
	"temp/app/constants/errorCode"
	util "temp/app/utils"
	"time"
)

type CachePolicy interface {
	Before(key string)
	IfNil(key string, values ...any)
	SetOperation(opt *redis.Operation)
}

type CrossPolicy struct {
	KeyRegx string
	Expire  time.Duration
	opt     *redis.Operation
}

func NewCrossPolicy(keyRegx string, expire time.Duration) *CrossPolicy {
	return &CrossPolicy{KeyRegx: keyRegx, Expire: expire}
}

func (p *CrossPolicy) Before(key string) {
	if !regexp.MustCompile(p.KeyRegx).MatchString(key) {
		util.Fail(errorCode.BAD_REQUEST, "error cache key")
	}
}

func (p *CrossPolicy) IfNil(key string, values ...any) {
	p.opt.HMSet(key, values...).Unwrap()
	p.opt.Expire(key, p.Expire)
}

func (p *CrossPolicy) SetOperation(opt *redis.Operation) {
	p.opt = opt
}
