/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Godyu97/vege9/vege"
	"github.com/spf13/cobra"
)

const (
	Cmd_newps    = "newps"    //随机生成字符串,默认8位,可输入数字指定长度
	Cmd_localip  = "localip"  //udp 得到 local ip
	Cmd_netinter = "netinter" //网卡上的全部 ip
	Cmd_ipip     = "ipip"     //http get ipip.net获取公网ip
	Cmd_ipsb     = "ipsb"     //http get ip.sb获取公网ip
	Cmd_hash     = "hash"     //通过hmac_sha2算法计算目标字符串的hash值
	Cmd_brackets = "brackets" //去除目标字符串不匹配括号
	Cmd_mac      = "mac"      //获取本机所有的mac地址
)

var MapCLI = map[string]CLI{
	Cmd_newps: {
		Cmd:     Cmd_newps,
		Func:    newps,
		Comment: "随机生成字符串,默认8位,可输入数字指定长度",
	},
	Cmd_localip: {
		Cmd:     Cmd_localip,
		Func:    localip,
		Comment: "udp 得到 local ip",
	},
	Cmd_netinter: {
		Cmd:     Cmd_netinter,
		Func:    netinter,
		Comment: "网卡上的全部 ip",
	},
	Cmd_ipip: {
		Cmd:     Cmd_ipip,
		Func:    ipip,
		Comment: "http get ipip.net获取公网ip",
	},
	Cmd_ipsb: {
		Cmd:     Cmd_ipsb,
		Func:    ipsb,
		Comment: "http get ip.sb获取公网ip",
	},
	Cmd_hash: {
		Cmd:     Cmd_hash,
		Func:    hmacsha2,
		Comment: "通过hmac_sha2算法计算目标字符串的hash值",
	},
	Cmd_brackets: {
		Cmd:     Cmd_brackets,
		Func:    brackets,
		Comment: "去除目标字符串不匹配括号",
	},
	Cmd_mac: {
		Cmd:     Cmd_mac,
		Func:    mac,
		Comment: "获取本机所有的mac地址",
	},
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
	//PreRun before Run
	PreRun: func(cmd *cobra.Command, args []string) {
		if vege.Ptr2Value(toggle) == true {
			Info("Help message for toggle cmd list")
			for _, item := range MapCLI {
				Info(item.Cmd, ":", item.Comment)
			}
			return
		}
		return
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if checkReturn() == true {
			return
		}
		argc := len(args)
		if argc > 0 {
			cli := MapCLI[args[0]]
			cli.Func(argc, args)
			return
		} else {
			Error("need param~")
			return
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
