package vege

import (
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

var SkipHttpsClient = &http.Client{Transport: &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}}

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
		err = errors.New("RcpWHuHB not find ip address~")
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

// GetMacAddr
// 获取本机的MAC地址
func GetMacAddr() (map[string]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("ePbWfcoN net.Interfaces() err:%w", err)
	}
	res := make(map[string]string, len(interfaces))
	for _, inter := range interfaces {
		m := inter.HardwareAddr.String() //获取本机MAC地址
		if m != "" {
			res[inter.Name] = m
		}
	}
	return res, nil
}

type IpTyp uint8

const (
	IpV4Typ IpTyp = iota
	IpV6Typ
)

const (
	IPSB4URL = `http://api-ipv4.ip.sb/ip`
	IPSB6URL = `http://api-ipv6.ip.sb/ip`
)

// GetPubIpVipsb 通过ip.sb获取公网ip
func GetPubIpVipsb(typ IpTyp) (ip string, err error) {
	//http实现
	url := ""
	switch typ {
	case IpV4Typ:
		url = IPSB4URL
		break
	case IpV6Typ:
		url = IPSB6URL
		break
	default:
		return "", fmt.Errorf("NXmZuhNz no support type:%d", typ)
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ip = strings.TrimSpace(string(b))
	return ip, nil
}

const (
	IPIP4URL = `http://myip.ipip.net/s`
)

// GetPubIpVipip 通过ip.sb获取公网ip 仅支持ipv4
func GetPubIpVipip(typ IpTyp) (ip string, err error) {
	//http实现
	url := ""
	switch typ {
	case IpV4Typ:
		url = IPIP4URL
		break
	case IpV6Typ:
		return "", fmt.Errorf("wZWdFhsH no support type:%d", typ)
	default:
		return "", fmt.Errorf("WMaPBIjm no support type:%d", typ)
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ip = strings.TrimSpace(string(b))
	return ip, nil
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

// Ip2Long 将 IPv4 字符串形式转为 uint32
func Ip2Long(str string) uint32 {
	ip := net.ParseIP(str)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}
