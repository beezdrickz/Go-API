package license

import (
	"database/sql"
	"go-api-frame/dto"
	"go-api-frame/dto/in"
	"go-api-frame/repository"
	"net/http"
	"time"
)

func (l *licenseService) InsertLicense(r *http.Request) (result dto.Response, err error) {
	var inputStruct in.LicenseDtoRequest
	var resultInsert int64
	inputStruct, err = l.ReadBody(r)
	if err != nil {
		return
	}
	resultInsert, err = l.doInsertLicense(inputStruct)
	if err != nil {
		return
	}
	result.Message = "Success Insert Data License"
	result.Result = resultInsert
	return
}

func (l *licenseService) doInsertLicense(inputStruct in.LicenseDtoRequest) (result int64, err error) {
	var tx *sql.Tx

	tx, err = l.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // Rollback in case of panic
			panic(p)          // Re-throw the panic
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	err = l.ValidateForInsert(tx, inputStruct)
	if err != nil {
		return
	}

	result, err = l.licenseDao.InsertLicense(tx, repository.LicenseModel{
		LicenseKey:  sql.NullString{String: inputStruct.PublicKey},
		MachineUUID: sql.NullString{String: inputStruct.MachineUUID},
		StoreID:     sql.NullInt64{Int64: inputStruct.StoreID},
		CreatedAt:   sql.NullTime{Time: time.Now()},
	})

	if err != nil {
		return
	}
	return
}
