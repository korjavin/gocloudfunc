package mypkg

import (
	"g.3pm.ai/go/lib/log"
	"net/http"
	"os"
)

//LambdaX for gcloud func
func LambdaX(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there " + os.Getenv("TEST")))
	log.Printf("requested")
}
