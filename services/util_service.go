package services

import (
	"encoding/json"
	"errors"
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
