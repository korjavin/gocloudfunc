package mypkg

import (
	"net/http"
	"os"
)

//LambdaX for gcloud func
func LambdaX(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there " + os.Getenv("TEST")))
}
