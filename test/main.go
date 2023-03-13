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
	//mux.RunListener(listen)

	log.Println(vegeTools.GetIpsb())

}
