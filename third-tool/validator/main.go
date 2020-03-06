package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserReq struct {
	Id   int    `json:"id" validate:"isdefault"`
	Name string `json:"name" validate:"-"`
}

func main() {
	user := UserReq{0, "dazuo"}
	validate := validator.New()
	err := validate.Struct(user)
	fmt.Println(err)
}
