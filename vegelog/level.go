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
