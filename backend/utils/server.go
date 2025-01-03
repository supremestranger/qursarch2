package utils

import (
	"fmt"
	"net/http"
)

const (
	API_VERSION = "/v1"
)

func EnableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func RegisterOnGet(url string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("GET %s", API_VERSION+url)
	http.HandleFunc(pattern, handler)
}

func RegisterOnPost(url string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("POST %s", API_VERSION+url)
	http.HandleFunc(pattern, handler)
}
