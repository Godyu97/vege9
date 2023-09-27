package vegeTools

import (
	"encoding/gob"
	"bytes"
)

//json实现 dst必须是指针
func DeepCopyByJson(src, dst any) error {
	if tmp, err := JsonMarshalV2(src); err != nil {
		return err
	} else {
		err = JsonUnmarshalV2(tmp, dst)
		return err
	}
}

//gob实现 dst必须是指针
func DeepCopyByGob(src, dst any) error {
	buffer := new(bytes.Buffer)
	if err := gob.NewEncoder(buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(buffer).Decode(dst)
}
