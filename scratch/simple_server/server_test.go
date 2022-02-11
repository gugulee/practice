package simpleserver

import (
	"net/http"
	"testing"
)

func Test_server(t *testing.T) {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
