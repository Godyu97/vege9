package vegelog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var single *VegeLog

// InitLogger
// 初始化logger记录日志的文件
// level 通常使用 	zapcore.DebugLevel / 	zapcore.InfoLevel
func InitLogger(file string, level zapcore.Level) {
	single = &VegeLog{}
	encoder := GetEncoder()
	single.WriteSyncer = NewLogWriter(file)
	//LevelEnabler 设置打印的日志级别
	core := zapcore.NewCore(encoder, single.WriteSyncer, level)
	// zap.AddStacktrace  添加将调用栈信息记录到日志中
	single.Logger = zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
}

func LogDebug(msg ...any) {
	if single != nil && single.Logger != nil {
		single.Logger.Debug(fmt.Sprintln(msg...))
	}
}

func LogInfo(msg ...any) {
	if single != nil && single.Logger != nil {
		single.Logger.Info(fmt.Sprintln(msg...))
	}
}

func LogWarn(msg ...any) {
	if single != nil && single.Logger != nil {
		single.Logger.Warn(fmt.Sprintln(msg...))
	}
}

func LogError(msg ...any) {
	if single != nil && single.Logger != nil {
		single.Logger.Error(fmt.Sprintln(msg...))
	}
}

func LogFatal(msg ...any) {
	if single != nil && single.Logger != nil {
		single.Logger.Fatal(fmt.Sprintln(msg...))
	}
}

// 为其他组件提供单例*zap.Logger
func SingleZapLogger() *zap.Logger {
	return single.GetZapLogger()
}

// 为其他组件提供单例Writer接口
func SingleZapWriter() zapcore.WriteSyncer {
	return single.GetZapWriter()
}

func SingleVegeLogger() Logger {
	return single
}
