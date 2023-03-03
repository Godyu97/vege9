package vegeTools

import (
	"testing"
)

func TestMapToObj(t *testing.T) {
	m := map[string]interface{}{
		"Id":    "0001",
		"Name":  "hongyu",
		"Phone": "133",
	}
	type user struct {
		Id    string
		Name  string
		Phone string
	}
	obj := &user{}
	err := MapToObj(m, obj)
	if err != nil {
		t.Error(err)
	}
	t.Log(obj)
}

func TestObjToMap(t *testing.T) {
	m := map[string]interface{}{}
	type user struct {
		Id    string
		Name  string
		Phone string
	}
	obj := &user{
		Id:    "0001",
		Name:  "hongyu",
		Phone: "133",
	}
	err := ObjToMap(*obj, &m)
	if err != nil {
		t.Error(err)
	}
	t.Log(m)
}
