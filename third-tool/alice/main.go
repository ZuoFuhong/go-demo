package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"time"
)

// alice 中间件链

type Middleware struct {
}

func (m Middleware) LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

// RecoverHandler recover panic
func (m Middleware) RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("user login ...")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Welcome!"))
}

func main() {
	m := Middleware{}
	mc := alice.New(m.LoggingHandler, m.RecoverHandler)

	router := mux.NewRouter()
	router.Handle("/login", mc.ThenFunc(Login)).Methods("GET")

	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		log.Println("Server startup failed.")
	}
}
