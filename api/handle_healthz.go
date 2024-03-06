package api

import "net/http"

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type: text/plain; charset=utf-8", "*")
	var isOk []byte = []byte("OK")
	w.Write(isOk)
}
