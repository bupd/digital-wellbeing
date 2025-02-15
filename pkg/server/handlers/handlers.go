package handlers

import (
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func Health(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("healthy"))
}
