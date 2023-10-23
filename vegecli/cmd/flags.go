package cmd

import (
	"github.com/Godyu97/vege9/vege"
)

var toggle *bool

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vege.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	toggle = rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle cmd list")
}

func checkReturn() bool {
	switch {
	case vege.Ptr2Value(toggle) == true:
		return true
	default:
		return false
	}
}
