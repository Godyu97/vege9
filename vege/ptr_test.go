package vege

import (
	"reflect"
	"testing"
)

type PtrObj struct {
	Name  string
	Hobby *string
}

func TestPtr2Value(t *testing.T) {
	o := &PtrObj{
		Name:  "vege9",
		Hobby: nil,
	}
	v := Ptr2Value(o)
	t.Log(reflect.DeepEqual(*o, v), v)

	vh := Ptr2Value(o.Hobby)
	t.Log(vh)
}

func TestValue2Ptr(t *testing.T) {
	o := PtrObj{
		Name:  "vege9",
		Hobby: nil,
	}
	ptr := Value2Ptr(o)
	t.Logf("%t,%p,%p", reflect.DeepEqual(&o, ptr), &o, ptr)
}
