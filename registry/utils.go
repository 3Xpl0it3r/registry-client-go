package registry

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

func newInsecureHttpClient()*http.Client{
	return &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
}



func newTlsHttpClient(cas ...string)*http.Client{
	ca_pool := x509.NewCertPool()
	for _,ca := range cas{
		if ca_raw,err := ioutil.ReadFile(ca);err != nil{
			panic(err)
		} else {
			ca_pool.AppendCertsFromPEM(ca_raw)
		}
	}
	return &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false, RootCAs: ca_pool}}}
}
