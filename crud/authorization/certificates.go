package authorization

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"sync"
)

var (
	signKey *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once sync.Once
)

//LoadCertificate singleton
func LoadCertificate(private, public string) error {
	var err error
	once.Do(func() {
		err = loadCertificates(private, public)
	})
	return err
}


//loadCertificates receives the path of the respective certificates
func loadCertificates(private, public string) error {
	privateBytes, err := ioutil.ReadFile(private)
	if err != nil {
		return err
	}

	publicBytes, err := ioutil.ReadFile(public)
	if err != nil {
		return err
	}

	return parseRSA(privateBytes, publicBytes)
}

//parseRSA validate the content of the ceriticates
func parseRSA(private, public []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		return err
	}

	return nil
}