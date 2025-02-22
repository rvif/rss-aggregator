package main

import "net/http"

//? Have to use this signature to define a http handler in a way which the go standard library expects
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}