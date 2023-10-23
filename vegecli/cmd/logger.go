package cmd

import (
	"fmt"
)

var gLogger Logger = &defaultLog{}

type Logger interface {
	Debug(a ...any)
	Info(a ...any)
	Error(a ...any)
	Fatal(a ...any)
}

type defaultLog struct{}

func (*defaultLog) Debug(a ...any) {
	fmt.Println(a...)
}

func (*defaultLog) Info(a ...any) {
	fmt.Println(a...)
}

func (*defaultLog) Error(a ...any) {
	fmt.Println(a...)
}

func (*defaultLog) Fatal(a ...any) {
	fmt.Println(a...)
}

func Debug(a ...any) {
	gLogger.Debug(a...)
}

func Info(a ...any) {
	gLogger.Info(a...)
}

func Error(a ...any) {
	gLogger.Error(a...)
}

func Fatal(a ...any) {
	gLogger.Fatal(a...)
}

func SetLogger(logger Logger) {
	if logger != nil {
		gLogger = logger
	}
	panic("asYPuGxu SetLogger nil")
}
