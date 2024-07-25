package vegelog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type Logger interface {
	Debug(msg ...any)
	Info(msg ...any)
	Warn(msg ...any)
	Error(msg ...any)
	Fatal(msg ...any)
}

type VegeLog struct {
	*zap.Logger
	zapcore.WriteSyncer
}

func (l *VegeLog) Debug(msg ...any) {
	if l != nil && l.Logger != nil {
		l.Logger.Debug(fmt.Sprintln(msg...))
	}
}

func (l *VegeLog) Info(msg ...any) {
	if l != nil && l.Logger != nil {
		l.Logger.Info(fmt.Sprintln(msg...))
	}
}

func (l *VegeLog) Warn(msg ...any) {
	if l != nil && l.Logger != nil {
		l.Logger.Warn(fmt.Sprintln(msg...))
	}
}

func (l *VegeLog) Error(msg ...any) {
	if l != nil && l.Logger != nil {
		l.Logger.Error(fmt.Sprintln(msg...))
	}
}

func (l *VegeLog) Fatal(msg ...any) {
	if l != nil && l.Logger != nil {
		l.Logger.Fatal(fmt.Sprintln(msg...))
	}
}

func (l *VegeLog) GetZapLogger() *zap.Logger {
	if l != nil && l.Logger != nil {
		return l.Logger
	}
	return nil
}

func (l *VegeLog) GetZapWriter() zapcore.WriteSyncer {
	if l != nil && l.WriteSyncer != nil {
		return l.WriteSyncer
	}
	return nil
}

func NewLogger(file string, level zapcore.Level) *VegeLog {
	l := &VegeLog{}
	encoder := GetEncoder()
	l.WriteSyncer = NewLogWriter(file)
	//LevelEnabler 设置打印的日志级别
	core := zapcore.NewCore(encoder, l.WriteSyncer, level)
	// zap.AddStacktrace  添加将调用栈信息记录到日志中
	l.Logger = zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
	return l
}

// 创建日志写入的zapcore.WriteSyncer,实现了io.Writer
func NewLogWriter(file string) zapcore.WriteSyncer {
	fileObj := &lumberjack.Logger{
		Filename: file, // 日志文件路径，默认 os.TempDir()
		MaxSize:  1000, // 每个日志文件保存1000M，默认 100M
		Compress: true, // 是否压缩，默认不压缩
	}
	// 利用io.MultiWriter支持文件和终端两个输出目标
	ws := io.MultiWriter(fileObj, os.Stdout)
	return zapcore.AddSync(ws)
}

func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器 ISO8601TimeEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}
