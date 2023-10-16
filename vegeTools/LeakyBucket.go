package vegeTools

import (
	"go.uber.org/ratelimit"
)

// Rate 基于	"go.uber.org/ratelimit" 的漏桶算法限流器的封装
// 固定速率执行业务函数
type Rate struct {
	ratelimit.Limiter
}

//创建QPS为n的限流器
func NewRateLeakyBucket(n int64) *Rate {
	return &Rate{Limiter: ratelimit.New(int(n))}
}

//定速执行业务函数
func (r *Rate) RateLimitDo(fn func()) {
	r.Take()
	fn()
}
