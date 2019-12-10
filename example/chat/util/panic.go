package util

import "fmt"

func RecoverPanic() {
	err := recover()
	if err != nil {
		fmt.Println("recover panic全局异常", err)
	}
}
