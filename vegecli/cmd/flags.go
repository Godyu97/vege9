package cmd

import (
	"fmt"
	"github.com/Godyu97/vege9/vege"
)

const version = "v1.1.1"

// flags
var (
	ftoggle  *bool
	fversion *bool
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vege.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	ftoggle = rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle cmd list")
	fversion = rootCmd.Flags().BoolP("version", "v", false, fmt.Sprintf("vegecli version : %s", version))
}

func checkReturn() bool {
	switch {
	case vege.Ptr2Value(ftoggle) == true:
		return true
	case vege.Ptr2Value(fversion) == true:
		return true
	default:
		return false
	}
}
