/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"fmt"
	"github.com/spf13/cobra"
	"log"
)

const (
	Cmd_newps    = "newps"    //随机生成字符串
	Cmd_localip  = "localip"  //udp 得到 local ip
	Cmd_netinter = "netinter" //网卡上的全部 ip
	Cmd_ipsb     = "ipsb"     //curl ip.sb
	Cmd_hash     = "hash"     //计算hmac sha2的hash值
	Cmd_brackets = "brackets"
)

var cmdL = []string{
	Cmd_newps,
	Cmd_localip,
	Cmd_netinter,
	Cmd_ipsb,
	Cmd_hash,
	Cmd_brackets,
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
		if *t {
			fmt.Println("helping cmd list:")
			for _, v := range cmdL {
				fmt.Println(v)
			}
			needReturn = true
			return
		}
		return
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if checkReturn() {
			return
		}
		argn := len(args)
		if argn > 0 {
			switch args[0] {
			case Cmd_newps:
				newps(args, argn)
				return
			case Cmd_localip:
				localip()
				return
			case Cmd_netinter:
				netinter()
				return
			case Cmd_ipsb:
				ipsb()
				return
			case Cmd_hash:
				hmacsha2(args, argn)
				return
			case Cmd_brackets:
				brackets(args, argn)
				return
			default:
				log.Println("unknown param~")
			}
		} else {
			log.Println("need param~")
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
