package vegeConfig

type Cfg struct {
	DSN string
}

var cfgObj *Cfg

// GetCfg
// 经过测试：返回值类型，无法被外部更改
func GetCfg() Cfg {
	if cfgObj != nil {
		return *cfgObj
	} else {
		panic("mBYKPPhT need InitCfg")
	}
}

func InitCfg(path string, file string) {
	cfgObj = new(Cfg)
	InitYamlCfg(path, file, cfgObj)
}
