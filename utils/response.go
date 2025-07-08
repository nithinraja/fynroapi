package utils

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func SuccessResponse(w http.ResponseWriter, message string) {
    JSONResponse(w, http.StatusOK, map[string]string{"message": message})
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
    JSONResponse(w, status, map[string]string{"error": message})
}
