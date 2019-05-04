package example

import (
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T)  {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}