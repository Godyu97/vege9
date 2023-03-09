package vegeTools

import (
	"testing"
	"time"
)

func TestQuit(t *testing.T) {
	q := NewQuit()
	go func() {
		time.Sleep(time.Second * 5)
		q.Close()
	}()
	q.WaitCloseWithFn(func() {
		t.Log("close ing~~~")
	})
}
