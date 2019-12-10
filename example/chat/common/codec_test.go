package common

import (
	"fmt"
	"testing"
)

func Test_Make(t *testing.T) {
	bytes := make([]byte, 20)
	fmt.Print(len(bytes))
}

func Test_bounds(t *testing.T) {
	var a = 1
	switch a {
	case 1:
		fmt.Println("hello 1")
		fallthrough
	case 2:
		fmt.Println("hello 2")
	default:
		fmt.Println("hello defajlt")
	}

}
