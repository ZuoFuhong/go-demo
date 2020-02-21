package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server failed err:", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	baseDir, _ := filepath.Abs("./")
	tmpl, e := template.ParseFiles(baseDir + "/example/template/t.html")
	if e != nil {
		panic(e)
	}

	user := struct {
		Name string
		Age  int
	}{"dazuo", 22}
	_ = tmpl.Execute(w, user)
}
