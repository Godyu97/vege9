package vegeTools

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/bytedance/sonic"
)

// jsoniter.ConfigCompatibleWithStandardLibrary 实现

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

//sonic.ConfigFastest 实现

func JsonUnmarshalFromStringV2(str string, v any) error {
	return sonic.ConfigFastest.
		UnmarshalFromString(str, v)
}

func JsonMarshalToStringV2(v any) (string, error) {
	return sonic.ConfigFastest.
		MarshalToString(v)
}

func JsonMarshalV2(v any) ([]byte, error) {
	return sonic.ConfigFastest.
		Marshal(v)
}

func JsonUnmarshalV2(data []byte, v any) error {
	return sonic.ConfigFastest.
		Unmarshal(data, v)
}
