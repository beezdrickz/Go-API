package utils

import (
	"go-api-frame/dto"
	"mime/multipart"
	"net/http"
)

// for endpoint that use multipart for request body
func WriteMultipartHttpResponseResult(token bool, w http.ResponseWriter, r *http.Request, handler func(interface{}, *dto.UserContext) (dto.Response, error)) {
	var (
		contextModel *dto.UserContext
		requestData  interface{}
		err          error
	)

	// Handle JWT validation if required
	if token {
		contextModel, err = checkTokenJwt(r.Header.Get("Authorization"))
		if err != nil {
			writeErrorResponse(w, r, "UNAUTHORIZED", err.Error(), http.StatusUnauthorized)
			return
		}
	}

	// Parse multipart form
	requestData, err = parseMultipartForm(r)
	if err != nil {
		writeErrorResponse(w, r, "ERROR", "Invalid multipart form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the handler function
	handleAndRespond(w, r, handler, requestData, contextModel)
}

func parseMultipartForm(r *http.Request) (*dto.MultipartData, error) {
	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
	if err != nil {
		return nil, err
	}

	formData := &dto.MultipartData{
		FormValues: make(map[string]string),
		Files:      make(map[string]*multipart.FileHeader),
	}

	// Extract text fields
	for key, values := range r.MultipartForm.Value {
		if len(values) > 0 {
			formData.FormValues[key] = values[0]
		}
	}

	// Extract file uploads
	for key, fileHeaders := range r.MultipartForm.File {
		if len(fileHeaders) > 0 {
			formData.Files[key] = fileHeaders[0] // Taking the first file only
		}
	}

	return formData, nil
}
