package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

// http://localhost:8080/info/dazuo/131
func Info(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("id: " + vars["id"])
	fmt.Println("username: " + vars["username"])

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("user info！"))
}

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"pwd" validate:"required"`
}

// POST 请求application/json
func Login(w http.ResponseWriter, r *http.Request) {
	loginReq := &LoginReq{}
	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		return
	}
	fmt.Println("loginReq: ", loginReq)

	// 参数validate
	validate := validator.New()
	err := validate.Struct(loginReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{\"code\": 1, \"msg\": \"Welcome!\"}"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/info/{username}/{id:[0-9]+}", Info).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./"))))

	_ = http.ListenAndServe("127.0.0.1:8080", router)
}
