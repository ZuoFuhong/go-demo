package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	calcSha256("I like donutsca07ca")
}

func calcSha256(data string) {
	// 前三个字节都是0的Hash
	hashInBytes := sha256.Sum256([]byte(data))
	hashVal := hex.EncodeToString(hashInBytes[:])
	fmt.Println(hashVal)
}
