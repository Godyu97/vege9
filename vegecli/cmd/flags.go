package cmd

var t *bool

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vege.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	t = rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle cmd list")
}

var needReturn bool

func checkReturn() bool {
	return needReturn
}
