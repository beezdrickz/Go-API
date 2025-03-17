package license

import (
	"errors"
	"go-api-frame/dto"
	"go-api-frame/dto/in"
	"go-api-frame/dto/out"
	"go-api-frame/repository"
	"net/http"
)

func (l *licenseService) ViewDetail(r *http.Request, ctx *dto.UserContext) (result dto.Response, err error) {
	var inputStruct in.LicenseDtoRequest
	inputStruct, err = l.ReadBody(r)
	if err != nil {
		return
	}
	result.Result, err = l.doViewDetail(inputStruct.ID)
	if err != nil {
		return
	}
	result.Message = "View Detail Success"
	return

}

func (l *licenseService) doViewDetail(id int64) (result out.LicenseDTOOut, err error) {

	resultDb, err := l.licenseDao.ViewDetail(id)
	if err != nil {
		return
	}
	if resultDb.ID.Int64 == 0 {
		err = errors.New("Data Not Exists")
		return
	}

	result = convertToView(resultDb)
	return
}

func convertToView(resultDb repository.LicenseModel) (result out.LicenseDTOOut) {
	result.ID = resultDb.ID.Int64
	result.MachineUUID = resultDb.MachineUUID.String
	result.StoreID = resultDb.StoreID.Int64
	result.PublicKey = resultDb.LicenseKey.String
	result.CreatedAt = resultDb.CreatedAt.Time
	return
}
