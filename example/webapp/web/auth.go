package web

import (
	"go_learning_notes/example/webapp/defs"
	"go_learning_notes/example/webapp/session"
	"net/http"
)

var HeaderFieldSession = "X-Session-Id"
var HeaderFieldUname = "X-User-Name"

func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HeaderFieldSession)
	if len(sid) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HeaderFieldUname, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HeaderFieldUname)
	if len(uname) == 0 {
		SendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
