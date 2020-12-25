package main

import (
	"go_learning_notes/example/webapp/defs"
	"go_learning_notes/example/webapp/taskrunner"
	"go_learning_notes/example/webapp/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middlewareHandler struct {
	r *httprouter.Router
	l *web.ConnLimiter
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		web.SendErrorResponse(w, defs.ErrorTooManyRequests)
		return
	}

	web.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func NewMiddlewareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middlewareHandler{}
	m.r = r
	m.l = web.NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/login", web.Login)
	return router
}

func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	mh := NewMiddlewareHandler(r, 2)
	_ = http.ListenAndServe(":8080", mh)
}
