package cmd

import (
	"fmt"
	"github.com/Godyu97/vege9/vegelog"
)

var gLogger vegelog.Logger = &defaultLog{}

type defaultLog struct{}

func (*defaultLog) Debug(a ...any) {
	fmt.Println(a...)
}

func (*defaultLog) Info(a ...any) {
	fmt.Println(a...)
}

func (l *defaultLog) Warn(a ...any) {
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

func SetLogger(logger vegelog.Logger) {
	if logger != nil {
		gLogger = logger
	}
	panic("asYPuGxu SetLogger nil")
}
