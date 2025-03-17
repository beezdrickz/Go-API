package endpoint

import (
	"encoding/json"
	"errors"
	"go-api-frame/dto"
	"go-api-frame/logger"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("Test")

func writeHttpResponseResult(token bool, w http.ResponseWriter, r *http.Request, handler func(*http.Request, *dto.UserContext) (dto.Response, error)) {
	var payloadResponse dto.PayloadResponse
	var response dto.Response
	var err error
	var contextModel *dto.UserContext
	if token {
		contextModel, err = checkTokenJwt(r.Header.Get("Authorization"))
		if err != nil {
			payloadResponse.Status = "UNAUTHORIZED"
			payloadResponse.Data.Message = err.Error()
			payloadResponse.Timestamp = time.Now()

			logger.LogError(r.Method, r.URL.Path, err.Error())
			w.WriteHeader(http.StatusUnauthorized) // Set appropriate error status
		}
		response, err = handler(r, contextModel) // Execute the handler function
		// Prepare payload response
		payloadResponse = dto.PayloadResponse{
			Timestamp: time.Now(),
			Status:    "SUCCESS",
			Data:      response,
		}

	} else {
		response, err = handler(r, contextModel) // Execute the handler function
		// Prepare payload response
		payloadResponse = dto.PayloadResponse{
			Timestamp: time.Now(),
			Status:    "SUCCESS",
			Data:      response,
		}
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
	// Set content type and encode JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payloadResponse)
}

func checkTokenJwt(tokenString string) (contextModel *dto.UserContext, err error) {
	var token *jwt.Token
	if tokenString == "" {
		return nil, errors.New("Unauthroized Token!!!")
	}
	//todo validate token if jwt
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unauthorized: Invalid token signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Unauthorized: Invalid token")
	}
	contextModel, err = parseToken(token)

	return
}

func parseToken(token *jwt.Token) (*dto.UserContext, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Convert claims to ContextModel
		contextModel := &dto.UserContext{}

		// Extract user ID
		if userID, ok := claims["userid"].(string); ok {
			contextModel.UserID = userID
		}

		// Extract role if available
		if role, ok := claims["role"].(string); ok {
			contextModel.Role = role
		}

		return contextModel, nil
	}

	return nil, errors.New("invalid token claims")
}
