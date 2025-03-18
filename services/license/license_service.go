package license

import (
	"database/sql"
	"encoding/json"
	"errors"
	"go-api-frame/dao"
	"go-api-frame/dto"
	"go-api-frame/dto/in"
	"go-api-frame/services"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type LicenseService interface {
	InsertLicense(r interface{}, contextModel *dto.UserContext) (result dto.Response, err error)
	InsertLicenseWithFile(r interface{}, contextModel *dto.UserContext) (result dto.Response, err error)
	ViewDetail(r interface{}, contextModel *dto.UserContext) (result dto.Response, err error)
}

type licenseService struct {
	licenseDao dao.LicenseDao
	db         *sql.DB
}

func NewLicenseService(db *sql.DB, dao dao.LicenseDao) LicenseService {
	return &licenseService{
		db:         db,
		licenseDao: dao,
	}
}

func mapToStructDTO(data interface{}, result *in.LicenseDtoRequest) (err error) {
	err = services.MapToStruct(data, &result)
	if err != nil {
		return
	}
	return
}

func mapToStructMultipartDTO(data interface{}, result *in.LicenseFileDTORequest) (err error) {
	err = services.MapToStructMultipart(data, &result.LicenseDTO, &result.FileLicense)
	if err != nil {
		return
	}
	return
}

func (l *licenseService) ReadBody(r *http.Request) (result in.LicenseDtoRequest, err error) {
	//read body json to struct
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	if len(body) != 0 {
		err = json.NewDecoder(r.Body).Decode(&result)
		if err != nil {
			err = errors.New("Invalid Json Format")
			return
		}
	}
	// for case path param
	pathParam := mux.Vars(r)
	tempID, _ := strconv.Atoi(pathParam["ID"])
	result.ID = int64(tempID)

	return
}

func (l *licenseService) ValidateForInsert(tx *sql.Tx, inputStruct in.LicenseDtoRequest) (err error) {
	var result int64
	result, err = l.licenseDao.CheckExistingLicense(tx, inputStruct.MachineUUID)
	if err != nil {
		return
	}
	if result != 0 {
		err = errors.New("This Machine already registered")
		return
	}
	result, err = l.licenseDao.CheckExistingLicense(tx, inputStruct.PublicKey)
	if err != nil {
		return
	}
	if result != 0 {
		err = errors.New(("License already used, please contact Admin!!!"))
		return
	}
	return
}
