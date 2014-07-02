package xero

import (
	"crypto/rsa"

	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"os"
)

func loadPrivateKeyFromFile(keyFilename string) (*rsa.PrivateKey, error) {
	f, err := os.Open(keyFilename)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(f)
	f.Close()
	if err != nil {
		return nil, fmt.Errorf("reading private key: %s", err.Error())
	}
	p, _ := pem.Decode(buf)
	if p == nil {
		return nil, fmt.Errorf("Could not load private key from %s", keyFilename)
	}
	pk, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return nil, fmt.Errorf("pasring private key: %s", err.Error())
	}
	return pk, nil
}

func loadCertificateFromFile(filename string) (*x509.Certificate, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(f)
	f.Close()
	if err != nil {
		return nil, fmt.Errorf("reading certificate: %s", err.Error())
	}
	p, _ := pem.Decode(buf)
	if p == nil {
		return nil, fmt.Errorf("Could not load certificate from %s", filename)
	}
	crt, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		return nil, fmt.Errorf("pasring certificate: %s", err.Error())
	}
	return crt, nil
}
