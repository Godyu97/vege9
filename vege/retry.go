package vege

import (
	"time"
)

type RetryConfig struct {
	Count    int
	SleepDur time.Duration
	Log      func(msg ...any)
}

func RetryFn(cfg RetryConfig, fn func() error) (err error) {
	for i := 0; i < cfg.Count; i++ {
		err = fn()
		if err != nil {
			if cfg.Log != nil {
				cfg.Log("YRxqLuKw RetryFn failed", i+1, err)
			}
			time.Sleep(cfg.SleepDur)
			continue
		} else {
			return nil
		}
	}
	return err
}
