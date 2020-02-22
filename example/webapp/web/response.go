package web

import (
	"encoding/json"
	"go-demo/example/webapp/defs"
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	_, _ = io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	_, _ = io.WriteString(w, resp)
}
