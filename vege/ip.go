package vege

import (
	"encoding/binary"
	"errors"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

// GetLocalIpv4List
// 获取所有非lo的网卡ip
func GetLocalIpv4List() ([]string, error) {
	//获取所有网卡
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				res = append(res, ipnet.IP.String())
			}
		}
	}
	if len(res) == 0 {
		err = errors.New("can not find ip address~")
		return nil, err
	}
	return res, nil

}

// GetLocalIpv4ByUdp
// 获取 localIp by  net.Dial("udp", "8.8.8.8:53")
func GetLocalIpv4ByUdp() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// GetMacMap
// 获取本机的MAC地址
func GetMacMap() map[string]string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	res := make(map[string]string)
	for _, inter := range interfaces {
		res[inter.Name] = inter.HardwareAddr.String() //获取本机MAC地址
	}
	return res
}

// GetPublicIp_ipsb
// curl ip.sb 的响应值 被GWF服务器不可用
func GetPublicIp_ipsb() (ip string, err error) {
	command := exec.Command("curl", "ip.sb")
	//Output 为将命令执行并返回输出
	output, err := command.Output()
	if err != nil {
		return "", err
	}
	//多余换行符
	res := string(output[:len(output)-1])
	return res, nil
}

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}
