package main

import (
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

// testing server connection
func TestStartServer(t *testing.T) {

	r := StartServer() //router call from main.go
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil) // get request returns ping

	r.ServeHTTP(w, req)

	//assert that http request is okay
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
func TestViewProverbs(t *testing.T) {

}
