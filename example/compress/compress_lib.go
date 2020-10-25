package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/golang/snappy"
)

func main() {
	useGzip()
}

func useGzip() {
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	_, err := writer.Write([]byte("hello gzip 12345678"))
	if err != nil {
		panic(err)
	}
	_ = writer.Flush()
	fmt.Println(len(buffer.Bytes()))

	reader, _ := gzip.NewReader(&buffer)
	carrier, _ := ioutil.ReadAll(reader)
	fmt.Println(string(carrier))
}

func useSnappy() {
	var buffer bytes.Buffer
	writer := snappy.NewBufferedWriter(&buffer)
	_, err := writer.Write([]byte("hello snappy 12345678"))
	if err != nil {
		panic(err)
	}
	_ = writer.Flush()
	fmt.Println(len(buffer.Bytes()))

	reader := snappy.NewReader(&buffer)
	carrier, _ := ioutil.ReadAll(reader)
	fmt.Println(string(carrier))
}
