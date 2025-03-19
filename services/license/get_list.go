package license

import (
	"go-api-frame/dto"
	"go-api-frame/dto/out"
	"go-api-frame/repository"
	"go-api-frame/services"
)

func (l *licenseService) GetList(
	r interface{},
	contextModel *dto.UserContext,
) (
	result dto.Response,
	err error,
) {

	var inputStruct dto.GetListDTO
	err = services.MapToStructGetList(r, &inputStruct)
	if err != nil {
		return
	}
	result.Result, err = l.doGetList(inputStruct, contextModel)
	if err != nil {
		return
	}
	result.Message = "Success Get List License"

	return
}

func (l *licenseService) doGetList(
	inputStruct dto.GetListDTO,
	contextModel *dto.UserContext,
) (
	result []out.ListLicenseDTOOut,
	err error,
) {
	var resultDB []repository.LicenseModel
	resultDB, err = l.licenseDao.ListLicense(inputStruct.Page, inputStruct.Limit, inputStruct.Filter)
	if err != nil {
		return
	}

	result = convertToList(resultDB)
	return

}

func convertToList(resultDB []repository.LicenseModel) (result []out.ListLicenseDTOOut) {
	for _, data := range resultDB {
		var temp out.ListLicenseDTOOut
		temp.ID = data.ID.Int64
		temp.MachineUUID = data.MachineUUID.String
		temp.PublicKey = data.LicenseKey.String
		temp.StoreID = data.StoreID.Int64
		result = append(result, temp)
	}
	return
}

func (l *licenseService) GetCount(
	r interface{},
	contextModel *dto.UserContext,
) (
	result dto.Response,
	err error,
) {
	var inputStruct dto.GetListDTO
	err = services.MapToStructGetList(r, &inputStruct)
	if err != nil {
		return
	}

	result.Result, err = l.doCount(inputStruct)
	if err != nil {
		return
	}
	result.Message = "Success Count License"

	return
}

func (l *licenseService) doCount(inputStruct dto.GetListDTO) (result int64, err error) {
	result, err = l.licenseDao.CountLicense(inputStruct.Filter)
	if err != nil {
		return
	}
	return
}
