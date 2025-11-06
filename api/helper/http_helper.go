package helper

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, httpStatusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(data)
}


func WriteError(w http.ResponseWriter, httpStatusCode int, message string) {
	WriteJSON(w, httpStatusCode, map[string]string{"error": message})
}