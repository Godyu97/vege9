package main

import (
	"github.com/Godyu97/vege9/vegeTools"
	"log"
)

func main() {
	//mux := gin.Default()
	//router.InitHttp(mux)
	//c, _ := vegeTools.TlsCertGenerateToMap()
	//pair, err := tls.X509KeyPair([]byte(c[vegeTools.CertName]), []byte(c[vegeTools.KeyName]))
	//if err != nil {
	//	panic(err)
	//}
	//listen, _ := tls.Listen("tcp", ":8080", &tls.Config{
	//	Certificates:   []tls.Certificate{pair},
	//	GetCertificate: nil,
	//})
	//mux.Run(":8080")
	//for {
	//	a := vegeTools.RandStringMask(20)
	//	log.Println(a)
	//	b := vegeTools.RandBytesMask(20)
	//	log.Println(b)
	//	time.Sleep(time.Second)
	//}
	hash := vegeTools.HashBySalt("123456asdasd", "jDnoKdXa")
	log.Println(hash, vegeTools.CheckBySalt("123456asdasd", hash, "jDnoKdXa"))
	log.Println("test version")
}
