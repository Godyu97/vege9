package vegeConfig

import (
	"testing"
)

func TestVegeConfig(t *testing.T) {
	cfgObj = new(Cfg)
	InitYamlCfg("./", "test_cfg.yaml", cfgObj)
	if cfgObj.DSN != "this is a test cfg yaml" {
		t.Error("Init yaml failed!")
	}
	t.Log(GetCfg())
	InitJsonCfg("./", "test_cfg.json", cfgObj)
	if cfgObj.DSN != "this is a test cfg json" {
		t.Error("Init json failed!")
	}
	t.Log(GetCfg())
}
