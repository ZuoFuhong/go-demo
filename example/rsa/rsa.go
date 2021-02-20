package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var certDir = "%s/cert/"

func init() {
	path, _ := os.Getwd()
	certDir = fmt.Sprintf(certDir, path)
}

func Encrypt(text string) ([]byte, error) {
	certBytes, err := ioutil.ReadFile(certDir + "public.key")
	if err != nil {
		log.Panic(err)
	}
	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(text))
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	certBytes, err := ioutil.ReadFile(certDir + "private.key")
	if err != nil {
		log.Panic(err)
	}
	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
