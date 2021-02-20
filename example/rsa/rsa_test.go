package rsa

import (
	"fmt"
	"log"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	text := "hello world"
	encryptBytes, err := Encrypt(text)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("encryptBytes: %v\n", encryptBytes)

	decryptBytes, err := Decrypt(encryptBytes)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("decryptText: ", string(decryptBytes))
}
