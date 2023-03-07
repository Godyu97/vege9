package config

import (
	"github.com/spf13/viper"
)

type Cfg struct {
}

var cfgObj *Cfg

// 返回值类型，无法被外部更改
func GetCfg() Cfg {
	if cfgObj != nil {
		return *cfgObj
	} else {
		panic("mBYKPPhT need InitCfg")
	}
}

func InitYamlCfg(path string, file string) {
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
	cfgObj = new(Cfg)
	if err := v.Unmarshal(&cfgObj); err != nil {
		panic(err)
	}
}

func InitJsonCfg(path string, file string) {
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
	cfgObj = new(Cfg)
	if err := v.Unmarshal(&cfgObj); err != nil {
		panic(err)
	}
}
