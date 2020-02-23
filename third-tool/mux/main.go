package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// http://localhost:8080/login/dazuo/131
func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("id: " + vars["id"])
	fmt.Println("username: " + vars["username"])

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Welcome！"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login/{username}/{id:[0-9]+}", Login).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("/Users/dazuo/workplace/go-demo/third-tool/mux"))))

	_ = http.ListenAndServe(":8080", router)
}
