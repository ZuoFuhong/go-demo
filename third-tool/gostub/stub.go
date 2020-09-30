package main

import (
	"fmt"
	"time"

	"github.com/prashantv/gostub"
)

// 1.为全局变量打桩
var counter = 100

func stubGlobalVariable() {
	stubs := gostub.Stub(&counter, 200)
	defer stubs.Reset()
	fmt.Println("Counter:", counter)
}

// 2.为函数打桩
var Exec = func() {
	fmt.Println("Exec")
}

func stubFunc() {
	stubs := gostub.Stub(&Exec, func() {
		fmt.Println("Stub Exec")
	})
	Exec()
	defer stubs.Reset()
}

// 3.为过程打桩
var timeNow = time.Now

func GetDate() int {
	return timeNow().Year()
}

func stubPath() {
	stubs := gostub.StubFunc(&timeNow, time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC))
	defer stubs.Reset()
	fmt.Println(GetDate())
}

func main() {
	//stubGlobalVariable()
	//stubFunc()
	stubPath()
}
