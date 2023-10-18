package vegeRouter

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Godyu97/vege9/vege"
	"github.com/gin-gonic/gin"
)

const (
	SendOk      = "SendOk"
	SendBad     = "SendBad"
	CheckAuth   = "CheckAuth"
	UriToFnName = "UriToFnName"
)

func Call(ctx *gin.Context, service any, methodName string) (response any, err error) {
	if methodName == SendOk ||
		methodName == SendBad ||
		methodName == CheckAuth ||
		methodName == UriToFnName {
		return "",
			errors.New("YdNRJNuJ Invalid api")
	}

	method, ok := reflect.TypeOf(service).MethodByName(methodName)
	if !ok {
		return "",
			errors.New(fmt.Sprintf("pWbvMQYt %s not find", methodName))
	}

	// NumIn() : apiObj , *gin.Context, *req
	// NumOut(): *resp err
	if method.Type.NumIn() != 3 || method.Type.NumOut() != 2 {
		return "",
			errors.New("fwxbKFmo Invalid api")
	}
	parameter := method.Type.In(2)
	req := reflect.New(parameter.Elem()).Interface()
	err = ctx.ShouldBind(req)
	if err != nil {
		return "", errors.New(fmt.Sprintf("bcCggifz ctx.ShouldBind(&req) err:%s", err))
	}

	in := make([]reflect.Value, 0, 2)
	in = append(in, reflect.ValueOf(ctx))
	in = append(in, reflect.ValueOf(req))
	call := make([]reflect.Value, 0, 2)
	err = vege.PanicToErr(func() {
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
