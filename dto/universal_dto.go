package dto

import "mime/multipart"

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

type GetListDTO struct {
	Page   int64  `json:"page"`
	Limit  int64  `json:"limit"`
	Filter string `json:"filter"`
}
