package web

import (
	"encoding/json"
	"go_learning_notes/example/webapp/dbops"
	"go_learning_notes/example/webapp/defs"
	"go_learning_notes/example/webapp/session"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, e := json.Marshal(su); e != nil {
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		SendNormalResponse(w, string(resp), 201)
	}
}
