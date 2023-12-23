package vegelog

import (
	"go.uber.org/zap/zapcore"
)

// debug mode：Info模式false，Debug模式true
func JudgeLogLevel(debug bool) zapcore.Level {
	if debug == false {
		return zapcore.InfoLevel
	} else {
		return zapcore.DebugLevel
	}
}

const ISO8601 = "2006-01-02T15:04:05.000Z0700"
