package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse call when operation successful
func JSONResponse(w http.ResponseWriter, statusCode int, payload any) {
	dat, err := json.Marshal(payload) // encode as json
	if err != nil {
		defer log.Fatal("Failed to marshal json response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(dat)
	if err != nil {
		log.Fatal(err)
	}

}

// ErrResponse call when error during execution
func ErrResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w) // pass the writer whom to send encoded data
	err := encoder.Encode(map[string]string{
		"error": message,
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

type PgErrDesc struct {
	Message string
	Code    string
}

// PgErrResponse call when postgres error occurs
func PgErrResponse(w http.ResponseWriter, statusCode int, payload PgErrDesc) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(payload)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
