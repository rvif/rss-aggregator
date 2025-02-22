package main

import "net/http"

//? Have to use this signature to define a http handler in a way which the go standard library expects
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}