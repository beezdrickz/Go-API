package license

import (
	"go-api-frame/dto"
	"go-api-frame/dto/in"
	"go-api-frame/services"
)

func (l licenseService) InsertLicenseWithFile(r interface{}, contextModel *dto.UserContext) (result dto.Response, err error) {
	var inputStruct in.LicenseFileDTORequest
	var resultInsert int64

	err = mapToStructMultipartDTO(r, &inputStruct)
	if err != nil {
		return
	}

	err = validateInsertDTO(inputStruct.LicenseDTO)
	if err != nil {
		return
	}

	err = services.SaveFileFromDTO(&inputStruct.FileLicense, "./uploads/license/", contextModel)
	if err != nil {
		return
	}

	resultInsert, err = l.doInsertLicense(inputStruct.LicenseDTO)
	if err != nil {
		return
	}

	result.Message = "Success Insert Data License with file"
	result.Result = resultInsert
	return
}
