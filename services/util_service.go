package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api-frame/dto"
	"io"
	"os"
	"path/filepath"
)

func MapToStruct(data interface{}, result interface{}) error {
	// Convert map to JSON
	var newData = make(map[string]interface{})
	newData = data.(map[string]interface{})

	jsonData, err := json.Marshal(newData)
	if err != nil {
		return errors.New("error marshaling map: " + err.Error())
	}

	// Unmarshal JSON into the provided struct
	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return errors.New("error unmarshaling JSON to struct: " + err.Error())
	}
	return nil
}

func MapToStructGetList(data interface{}, result *dto.GetListDTO) error {
	// Convert map to JSON
	var newData = make(map[string]interface{})
	newData = data.(map[string]interface{})

	jsonData, err := json.Marshal(newData)
	if err != nil {
		return errors.New("error marshaling map: " + err.Error())
	}

	// Unmarshal JSON into the provided struct
	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return errors.New("error unmarshaling JSON to struct: " + err.Error())
	}
	return nil
}

func MapToStructMultipart(data interface{}, result interface{}, resultFile *dto.FileDTO) error {
	dataMultipart := data.(*dto.MultipartData)
	newData := dataMultipart.FormValues["form_values"]
	err := json.Unmarshal([]byte(newData), result)
	if err != nil {
		return errors.New("error unmarshaling JSON to struct: " + err.Error())
	}
	for fileKey, fileHeader := range dataMultipart.Files {
		file, err := fileHeader.Open()
		if err != nil {
			return errors.New("failed to open file " + fileKey + ": " + err.Error())
		}
		defer file.Close()
		resultFile.FileName = fileHeader.Filename
		resultFile.FileType = fileHeader.Header.Get("Content-Type")
		resultFile.FileSize = fileHeader.Size
		fileContent, err := io.ReadAll(file)
		if err != nil {
			return errors.New("failed to read file content: " + err.Error())
		}
		resultFile.FileContent = fileContent

	}
	return nil
}

// to be adjust again
func SaveFileFromDTO(resultFile *dto.FileDTO, saveDirectory string, contextModel *dto.UserContext) error {
	// Ensure the save directory exists
	err := os.MkdirAll(saveDirectory, os.ModePerm)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create directory %s: %s", saveDirectory, err.Error()))
	}

	// Generate the full file path
	filePath := filepath.Join(saveDirectory, resultFile.FileName)

	// Create or open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create file %s: %s", filePath, err.Error()))
	}
	defer file.Close() // Ensure file is closed after writing

	// Write the file content to disk (from []byte in FileDTO)
	_, err = file.Write(resultFile.FileContent)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to write content to file %s: %s", filePath, err.Error()))
	}

	// Successfully saved the file
	return nil
}
