// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	fmt.Println(id.String())
}
