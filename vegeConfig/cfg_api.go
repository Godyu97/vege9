package vegeConfig

import (
	"github.com/spf13/viper"
)

// unmarshal cfg obj pointer
func InitYamlCfg(path string, file string, obj any) {
	v := viper.New()
	//path
	v.AddConfigPath(path)
	//filename 可忽略文件后缀
	v.SetConfigName(file)
	//解析配置文件的类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(obj); err != nil {
		panic(err)
	}
}

// unmarshal cfg obj pointer
func InitJsonCfg(path string, file string, obj any) {
	v := viper.New()
	//path
	v.AddConfigPath(path)
	//filename 可忽略文件后缀
	v.SetConfigName(file)
	//解析配置文件的类型
	v.SetConfigType("json")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(obj); err != nil {
		panic(err)
	}
}
