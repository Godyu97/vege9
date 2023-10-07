package main

import (
	"github.com/Godyu97/vege9/vegeConfig"
	"fmt"
)

func main() {
	vegeConfig.InitCfg("./vegeConfig", "test_cfg.yaml")
	cfg := vegeConfig.GetCfg()
	cfg.DSN = "modify"
	fmt.Printf("out:%p\n", &cfg)
	cfg2 := vegeConfig.GetCfg()
	fmt.Printf("out:%p\n", &cfg2)
}
