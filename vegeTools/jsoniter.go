package vegeTools

import (
	jsoniter "github.com/json-iterator/go"
)

func JsonUnmarshalFromString(str string, v any) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.
		UnmarshalFromString(str, v)
}

func JsonMarshalToString(v any) (string, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.
		MarshalToString(v)
}

func JsonMarshal(v any) ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.
		Marshal(v)
}

func JsonUnmarshal(data []byte, v any) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.
		Unmarshal(data, v)
}
