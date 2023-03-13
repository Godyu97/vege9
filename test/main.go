package main

import (
	"crypto/tls"
	"github.com/Godyu97/vege9/router"
	"github.com/Godyu97/vege9/vegeTools"
	"github.com/gin-gonic/gin"
)

func main() {
	mux := gin.Default()
	router.InitHttp(mux)
	c, _ := vegeTools.TlsCertGenerateToMap()
	pair, err := tls.X509KeyPair([]byte(c[vegeTools.CertName]), []byte(c[vegeTools.KeyName]))
	if err != nil {
		panic(err)
	}
	listen, _ := tls.Listen("tcp", ":8080", &tls.Config{
		Certificates:   []tls.Certificate{pair},
		GetCertificate: nil,
	})
	mux.RunListener(listen)

}
