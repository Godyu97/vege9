package vegeTools

import (
	"testing"
	"time"
)

func TestCostTimeDur(t *testing.T) {
	t.Log(CostTimeDur(func() {
		time.Sleep(time.Second * 5)
	}))
}
