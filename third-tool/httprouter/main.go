package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/login/:username/:password", Login)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("username = " + ps.ByName("username"))
	fmt.Println("password = " + ps.ByName("password"))
	_, _ = fmt.Fprint(w, "Welcome!\n")
}
