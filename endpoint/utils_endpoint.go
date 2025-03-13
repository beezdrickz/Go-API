package endpoint

import (
	"encoding/json"
	"errors"
	"go-api-frame/dto"
	"go-api-frame/logger"
	"net/http"
	"strings"
	"time"
)

func writeHttpResponseResult(token bool, w http.ResponseWriter, r *http.Request, handler func(*http.Request) (dto.Response, error)) {
	var payloadResponse dto.PayloadResponse
	var err error
	if token {
		err = checkTokenJwt(r.Header.Get("Authorization"))
		if err != nil {
			payloadResponse.Status = "UNAUTHORIZED"
			payloadResponse.Data.Message = err.Error()
			payloadResponse.Timestamp = time.Now()

			logger.LogError(r.Method, r.URL.Path, err.Error())
			w.WriteHeader(http.StatusUnauthorized) // Set appropriate error status
		}
	} else {

		response, err := handler(r) // Execute the handler function

		// Prepare payload response
		payloadResponse = dto.PayloadResponse{
			Timestamp: time.Now(),
			Status:    "SUCCESS",
			Data:      response,
		}

		if err != nil {
			payloadResponse.Status = "ERROR"
			payloadResponse.Data.Message = err.Error()
			payloadResponse.Data.Result = nil

			logger.LogError(r.Method, r.URL.Path, err.Error())
			if !strings.Contains(err.Error(), "sql") {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				payloadResponse.Data.Message = "Internal Server Error, please Contact Admin!"
				w.WriteHeader(http.StatusInternalServerError) // Set appropriate error status
			}

		} else {
			logger.LogRequest(r.Method, r.URL.Path)
			w.WriteHeader(http.StatusOK)
		}
	}
	// Set content type and encode JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payloadResponse)
}

func checkTokenJwt(token string) error {
	if token == "" {
		return errors.New("Unauthroized Token!!!")
	}
	//todo validate token if jwt
	return nil
}
