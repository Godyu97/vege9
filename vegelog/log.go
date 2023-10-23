package vegelog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var writer zapcore.WriteSyncer

// InitLogger
// 初始化logger记录日志的文件
// level 通常使用 	zapcore.DebugLevel / 	zapcore.InfoLevel
func InitLogger(file string, level zapcore.Level) {
	encoder := getEncoder()
	writeSyncer := getLogWriter(file)
	//LevelEnabler 设置打印的日志级别
	core := zapcore.NewCore(encoder, writeSyncer, level)
	// zap.AddStacktrace  添加将调用栈信息记录到日志中
	logger = zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
}

func LogDebug(msg ...any) {
	if logger != nil {
		logger.Debug(fmt.Sprintln(msg...))
	}
}

func LogInfo(msg ...any) {
	if logger != nil {
		logger.Info(fmt.Sprintln(msg...))
	}
}

func LogWarn(msg ...any) {
	if logger != nil {
		logger.Warn(fmt.Sprintln(msg...))
	}
}

func LogError(msg ...any) {
	if logger != nil {
		logger.Error(fmt.Sprintln(msg...))
	}
}

func LogFatal(msg ...any) {
	if logger != nil {
		logger.Fatal(fmt.Sprintln(msg...))
	}
}

func getLogWriter(file string) zapcore.WriteSyncer {
	if writer != nil {
		return writer
	}
	fileObj := &lumberjack.Logger{
		Filename: file, // 日志文件路径，默认 os.TempDir()
		MaxSize:  1000, // 每个日志文件保存1000M，默认 100M
		Compress: true, // 是否压缩，默认不压缩
	}
	// 利用io.MultiWriter支持文件和终端两个输出目标
	ws := io.MultiWriter(fileObj, os.Stdout)
	writer = zapcore.AddSync(ws)
	return writer
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器 ISO8601TimeEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 为其他组件提供logger接口
func GetZapLogger() *zap.Logger {
	if logger != nil {
		return logger
	}
	panic("InitLogger first")
}

// 为其他组件提供Writer接口
func GetLogWriter() io.Writer {
	if writer != nil {
		return writer
	}
	return os.Stdout
}
