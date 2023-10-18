package vege

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
)

type Quit struct {
	ch chan os.Signal
}

func NewQuit() *Quit {
	return &Quit{ch: make(chan os.Signal)}
}

func (q *Quit) Close() error {
	if q == nil {
		return errors.New("VAdjAqPi nil quit")
	}
	q.ch <- syscall.SIGKILL
	return nil
}

func (q *Quit) WaitCloseWithFn(fn func()) {
	signal.Notify(q.ch, syscall.SIGINT, syscall.SIGTERM)
	<-q.ch
	//处理关闭逻辑
	fn()
	//关闭成功，退出进程
	os.Exit(0)
}
