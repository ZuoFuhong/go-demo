// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package testify

import "fmt"

func Eval(op rune) {
	switch op {
	case '+':
		fmt.Printf("//////// + ////////\n")
	case '-':
		fmt.Printf("//////// - ////////\n")
	case '*':
		fmt.Printf("//////// * ////////\n")
	case '/':
		fmt.Printf("//////// / ////////\n")
	}
}
