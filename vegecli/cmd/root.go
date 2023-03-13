/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"fmt"
	"github.com/spf13/cobra"

	"github.com/Godyu97/vege9/vegeTools"
	"log"
	"strconv"
)

const (
	Cmd_newps    = "newps"    //随机生成字符串
	Cmd_localip  = "localip"  //udp 得到 local ip
	Cmd_netinter = "netinter" //网卡上的全部 ip
	Cmd_ipsb     = "ipsb"     //curl ip.sb
)

var cmdL = []string{
	Cmd_newps,
	Cmd_localip,
	Cmd_netinter,
	Cmd_ipsb,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vegecli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if *t {
			fmt.Println("helping cmd list:")
			for _, v := range cmdL {
				fmt.Println(v)
			}
			return
		}
		if len(args) > 0 {
			switch args[0] {
			case Cmd_newps:
				n := 8
				if len(args) == 2 {
					var err error
					n, err = strconv.Atoi(args[1])
					if err != nil {
						log.Println("IZHSMxwu 请输入有效数字：", err)
						return
					}
				}
				fmt.Println(vegeTools.RandStringBytesMask(n))
			case Cmd_localip:
				ip, err := vegeTools.GetLocalIpv4ByUdp()
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Println(ip)
			case Cmd_netinter:
				list, err := vegeTools.GetLocalIpv4List()
				if err != nil {
					log.Println(err)
					return
				}
				for _, ip := range list {
					fmt.Println(ip)
				}
			case Cmd_ipsb:
				ip, err := vegeTools.GetPublicIp_ipsb()
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Println(ip)
			}

		} else {
			fmt.Println("need param~")
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

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
