package apiserver

import "net/http"

type responseWrite struct {
	http.ResponseWriter
	code int
}

func (w *responseWrite) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
