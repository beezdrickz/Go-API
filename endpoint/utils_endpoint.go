package endpoint

import (
	"encoding/json"
	"errors"
	"go-api-frame/dto"
	"go-api-frame/logger"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var jwtSecret = []byte("JWT Rocks!")

func writeHttpResponseResult(token bool, w http.ResponseWriter, r *http.Request, handler func(interface{}, *dto.UserContext) (dto.Response, error)) {
	var (
		payloadResponse dto.PayloadResponse
		response        dto.Response
		contextModel    *dto.UserContext
		requestBody     = make(map[string]interface{})
		err             error
	)

	// Handle JWT validation if required
	if token {
		contextModel, err = checkTokenJwt(r.Header.Get("Authorization"))
		if err != nil {
			writeErrorResponse(w, r, "UNAUTHORIZED", err.Error(), http.StatusUnauthorized)
			return
		}
	}

	// Parse request body if present
	if r.Body != nil && r.ContentLength > 0 {
		err = ParseRequestBody(r, &requestBody)
		if err != nil {
			writeErrorResponse(w, r, "ERROR", "Invalid request body", http.StatusBadRequest)
			return
		}
	}
	// Extract path & query parameters and inject into requestBody
	extractParam(r, &requestBody)

	response, err = handler(requestBody, contextModel)
	if err != nil {
		statusCode := http.StatusBadRequest
		message := err.Error()

		if strings.Contains(err.Error(), "sql") {
			message = "Internal Server Error, please Contact Admin!"
			statusCode = http.StatusInternalServerError
		}

		writeErrorResponse(w, r, "ERROR", message, statusCode)
		return
	}

	// Success response
	payloadResponse = dto.PayloadResponse{
		Timestamp: time.Now(),
		Status:    "SUCCESS",
		Data:      response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payloadResponse)
	logger.LogRequest(r.Method, r.URL.Path)
}

func checkTokenJwt(tokenString string) (contextModel *dto.UserContext, err error) {
	var token *jwt.Token
	if tokenString == "" {
		return nil, errors.New("Unauthroized Token!!!")
	}
	splittedToken := strings.Split(tokenString, " ")
	if len(splittedToken) != 2 {
		return nil, errors.New("Unauthroized Token!!!")
	}
	//todo validate token if jwt
	token, err = jwt.Parse(splittedToken[1], func(token *jwt.Token) (interface{}, error) {
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

// ParseRequestBody reads and decodes the JSON request body into the provided interface.
func ParseRequestBody(r *http.Request, target *map[string]interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}
func extractParam(r *http.Request, requestBody *map[string]interface{}) {
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			(*requestBody)[key] = values[0]
		}
	}
	vars := mux.Vars(r)
	for key, value := range vars {
		if key == "ID" {
			numValue, _ := strconv.Atoi(value)
			(*requestBody)[key] = int64(numValue)
		} else {
			(*requestBody)[key] = value
		}
	}
}

func writeErrorResponse(w http.ResponseWriter, r *http.Request, status, message string, statusCode int) {
	payloadResponse := dto.PayloadResponse{
		Timestamp: time.Now(),
		Status:    status,
		Data: dto.Response{
			Message: message,
			Result:  nil,
		},
	}

	logger.LogError(r.Method, r.URL.Path, message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payloadResponse)
}
