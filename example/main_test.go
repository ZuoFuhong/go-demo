package main

import (
	"testing"
)

func Test_case(t *testing.T) {
	t.Log("hello test log")
	t.Logf("name %d", 22)
	t.Error("error log")
	// 只标记错误不终止测试的方法
	t.Fail()
	// 终止当前测试用例时，可以使用 FailNow
	t.Fatal("fatal log")
}

func Benchmark_case(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1To20()
	}
}

func Print1To20() {
	//fmt.Println("hello")
}
