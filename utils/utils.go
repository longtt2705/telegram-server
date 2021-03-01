package utils

import (
	"encoding/json"
	"net/http"
)

// SendJSON send json respone to client
func SendJSON(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(value)
}
