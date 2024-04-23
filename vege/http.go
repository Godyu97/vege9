package vege

import (
	"net/http"
	"crypto/tls"
)

//InsecureSkipVerify  DisableKeepAlives true
var HttpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableKeepAlives: true,
	},
}
