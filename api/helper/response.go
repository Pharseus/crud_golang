package helper

import (
	"encoding/json"
	"net/http"

	"github.com/Pharseus/crud_golang.git/api/payloads"
)

func RespondSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payloads.SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func RespondError(w http.ResponseWriter, status int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	json.NewEncoder(w).Encode(payloads.ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   errMsg,
	})
}
