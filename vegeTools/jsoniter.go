package vegeTools

import (
	jsoniter "github.com/json-iterator/go"
)

func JsonUnmarshalFromString(str string, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.
		UnmarshalFromString(str, v)
}

func JsonMarshalToString(v interface{}) (string, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.
		MarshalToString(v)
}

func JsonMarshal(v interface{}) ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.
		Marshal(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.
		Unmarshal(data, v)
}
