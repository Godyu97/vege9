package cmd

import (
	"fmt"
	"github.com/Godyu97/vege9/vege"
	"log"
	"strconv"
)

func newps(args []string, argn int) {
	n := 8
	if argn == 2 {
		var err error
		n, err = strconv.Atoi(args[1])
		if err != nil {
			log.Println("IZHSMxwu 请输入有效数字：", err)
			return
		}
	}
	fmt.Println(vege.RandStringMask(n))
}

func localip() {
	ip, err := vege.GetLocalIpv4ByUdp()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ip)
}

func netinter() {
	list, err := vege.GetLocalIpv4List()
	if err != nil {
		log.Println(err)
		return
	}
	for _, ip := range list {
		fmt.Println(ip)
	}
}

func ipsb() {
	ip, err := vege.GetPublicIp_ipsb()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ip)
}

func hmacsha2(args []string, argn int) {
	if argn != 2 {
		log.Println("bad param~")
		return
	}
	fmt.Println(vege.HashBySalt(args[1], ""))
}

func brackets(args []string, argn int) {
	if argn != 2 {
		log.Println("bad param~")
		return
	}
	s := vege.RemoveInvalidParentheses(args[1], [2]rune{'(', ')'})
	s = vege.RemoveInvalidParentheses(s, [2]rune{'（', '）'})
	fmt.Println(s)
}
