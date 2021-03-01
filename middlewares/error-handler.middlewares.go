package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// ErrorHandler handles all errors
func ErrorHandler(err error, w http.ResponseWriter, r *http.Request, status int) {
	log.Fatal(err)
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
