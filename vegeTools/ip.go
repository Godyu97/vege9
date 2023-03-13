package vegeTools

import (
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
)

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

// http获取 ip.sb 的响应值
func GetIpsb() (ip string, err error) {
	resp, err := http.Get("http://ip.sb")
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(resp.Body)
	return string(b), nil
}
