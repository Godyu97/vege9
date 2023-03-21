package vegeTools

import (
	"go.uber.org/ratelimit"
)

// Rate 基于	"go.uber.org/ratelimit" 的漏桶算法限流器的封装
type Rate struct {
	ratelimit.Limiter
}

func NewRateLeakyBucket(n int64) *Rate {
	return &Rate{Limiter: ratelimit.New(int(n))}
}

func (r *Rate) RateLimitDo(fn func()) {
	r.Take()
	fn()
}
