package authorizer

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"path/filepath"
)

var rsaPriv *rsa.PrivateKey
var rsaPub interface{}

func getPubKey() (interface{}, error) {
	publicKeyPath, err := filepath.Abs(filepath.Join("keys", "rabbit-key.pub"))

	if err != nil {
		return nil, err
	}

	derBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(derBytes)

	rsaPub, err = x509.ParsePKIXPublicKey(block.Bytes)

	return rsaPub, err
}

func getPrivKey() (*rsa.PrivateKey, error) {

	if rsaPriv != nil {
		return rsaPriv, nil
	}

	privateKeyPath, err := filepath.Abs(filepath.Join("keys", "rabbit-key.key"))
	if err != nil {
		return nil, err
	}

	der, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	block2, _ := pem.Decode(der)
	rsaPriv, err = x509.ParsePKCS1PrivateKey(block2.Bytes)

	if err != nil {
		return nil, err
	}

	return rsaPriv, nil
}
