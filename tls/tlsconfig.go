package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

const (
	RootCAFile = ""
	CertFile   = ""
	KeyFile    = ""
)

func SetCredentials() (*tls.Config, error) {
	rootCA, err := ioutil.ReadFile(RootCAFile)
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(rootCA)
	cert, err := tls.LoadX509KeyPair(CertFile, KeyFile)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		RootCAs:            pool,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		NextProtos:         []string{"x-amzn-mqtt-ca"},
	}, nil
}
