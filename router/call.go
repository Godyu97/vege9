package router

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Godyu97/vege9/vegeTools"
)

const (
	SendOk  = "SendOk"
	SendBad = "SendBad"
)

func Call(ctx context.Context, service interface{}, methodName string, request string) (response interface{}, err error) {
	if methodName == SendOk || methodName == SendBad {
		return "",
			errors.New("YdNRJNuJ 非可用的api")
	}

	method, ok := reflect.TypeOf(service).MethodByName(methodName)
	if !ok {
		return "",
			errors.New(fmt.Sprintf("YdNRJNuJ %s not find", methodName))
	}

	// NumIn() : apiObj , *gin.Context, *req
	// NumOut(): *resp err
	if method.Type.NumIn() != 3 || method.Type.NumOut() != 2 {
		return "",
			errors.New(fmt.Sprintf("fwxbKFmo 非可用的api"))
	}
	parameter := method.Type.In(2)
	req := reflect.New(parameter.Elem()).Interface()
	if len(request) != 0 {
		err = vegeTools.UnmarshalFromString(request, req)
		if err != nil {
			log.Println("bcCggifz ", err)
			return "",
				errors.New("bcCggifz req 解析错误")
		}
	}
	in := make([]reflect.Value, 0, 2)
	in = append(in, reflect.ValueOf(ctx))
	in = append(in, reflect.ValueOf(req))
	call := make([]reflect.Value, 0, 2)
	err = vegeTools.PanicToErr(func() {
		call = reflect.ValueOf(service).MethodByName(methodName).Call(in)
	})
	if err != nil {
		return "", err
	}
	if call[1].Interface() != nil {
		err = call[1].Interface().(error)
		return "", err
	}
	return call[0].Interface(), nil
}
