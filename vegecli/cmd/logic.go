package cmd

import (
	"github.com/Godyu97/vege9/vege"
	"strconv"
)

type LogicFunc func(argc int, argv []string)

type CLI struct {
	Cmd     string
	Func    LogicFunc
	Comment string
}

func newps(argc int, argv []string) {
	n := 8
	if argc == 2 {
		var err error
		n, err = strconv.Atoi(argv[1])
		if err != nil {
			Error("IZHSMxwu 请输入有效数字：", err)
			return
		}
	}
	Info(vege.RandStringMask(n))
}

func localip(argc int, argv []string) {
	ip, err := vege.GetLocalIpv4ByUdp()
	if err != nil {
		Error(err)
		return
	}
	Info(ip)
}

func netinter(argc int, argv []string) {
	list, err := vege.GetLocalIpv4List()
	if err != nil {
		Error(err)
		return
	}
	for _, ip := range list {
		Info(ip)
	}
}

func ipsb(argc int, argv []string) {
	ip, err := vege.GetPubIpVipsb(vege.IpV4Typ)
	if err != nil {
		Error(err)
		return
	}
	Info(ip)
}

func ipip(argc int, argv []string) {
	ip, err := vege.GetPubIpVipip(vege.IpV4Typ)
	if err != nil {
		Error(err)
		return
	}
	Info(ip)
}

func hmacsha2(argc int, argv []string) {
	if argc != 2 {
		Error("bad param~")
		return
	}
	Info(vege.HmacHashWithSalt(argv[1], ""))
}

func brackets(argc int, argv []string) {
	if argc != 2 {
		Error("bad param~")
		return
	}
	s := vege.RemoveInvalidParentheses(argv[1], [2]rune{'(', ')'})
	s = vege.RemoveInvalidParentheses(s, [2]rune{'（', '）'})
	Info(s)
}

func mac(argc int, argv []string) {
	macs, err := vege.GetMacAddr()
	if err != nil {
		Error(err)
		return
	}
	for k, addr := range macs {
		Info(k, ":", addr)
	}
}
