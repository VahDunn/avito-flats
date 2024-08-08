package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
}

func sendErrorResponse(w http.ResponseWriter, message string, requestID string, statusCode int) {
	errorResponse := ErrorResponse{
		Message:   message,
		RequestID: requestID,
		Code:      statusCode,
	}

	response, _ := json.Marshal(errorResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
