package dto

import (
	"mime/multipart"
	"time"
)

type PayloadResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp`
	Data      Response  `json:"data"`
}

type Response struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type MultipartData struct {
	FormValues map[string]string                `json:"form_values"`
	Files      map[string]*multipart.FileHeader `json:"files"`
}

type FileDTO struct {
	FileName    string `json:"file_name"`
	FileSize    int64  `json:"file_size"`
	FileType    string `json:"file_type"`
	FileContent []byte
}
