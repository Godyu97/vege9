package vegeTools

import (
	"fmt"
)

func PanicToErr(fn func()) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	fn()
	return err
}
