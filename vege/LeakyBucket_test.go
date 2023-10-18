package vege

import (
	"testing"
	"time"
)

func TestRateLimitDo(t *testing.T) {
	r := NewRateLeakyBucket(1)
	go func() {
		for {
			r.RateLimitDo(func() {
				t.Log("hello in")
			})
		}
	}()
	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		r.RateLimitDo(func() {
			t.Log("hello out")
		})
	}
}
