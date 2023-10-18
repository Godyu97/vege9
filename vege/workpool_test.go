package vege

import (
	"testing"
	"time"
)

func TestWorkerPool_Submit(t *testing.T) {
	tasks := NewWorkerPool(128)
	for i := 0; i < 1000; i++ {
		ii := i
		tasks.Submit(func() {
			time.Sleep(time.Second)
			t.Log(ii, " done")
		})
	}
	t.Log("tasks StopWait log")
	tasks.StopWait()
}

func TestNewWorkerPool_SubmitBlockBySize(t *testing.T) {
	tasks := NewWorkerPool(128)
	for i := 0; i < 1000; i++ {
		ii := i
		tasks.SubmitBlockBySize(func() {
			time.Sleep(time.Second)
			t.Log(ii, " done")
		})
	}
	t.Log("tasks StopWait log")
	tasks.StopWait()
}
