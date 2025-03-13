package dao

import (
	"database/sql"
	"fmt"
	"go-api-frame/repository"
)

type LicenseDao interface {
	InsertLicense(tx *sql.Tx, model repository.LicenseModel) (result int64, err error)
	CheckExistingMachine(tx *sql.Tx, machineUUID string) (result int64, err error)
	CheckExistingLicense(tx *sql.Tx, licenseKey string) (result int64, err error)
	CheckLicenseAndMachine(model repository.LicenseModel) (result bool, err error)
	ViewDetail(id int64) (result repository.LicenseModel, err error)
}

type licenseDAO struct {
	db        *sql.DB
	tableName string
}

func NewLicenseDAO(db *sql.DB) *licenseDAO {
	return &licenseDAO{
		db:        db,
		tableName: "tbl_license",
	}
}

func (d *licenseDAO) InsertLicense(tx *sql.Tx, model repository.LicenseModel) (result int64, err error) {
	query := fmt.Sprintf(`INSERT INTO %s(license_key,machine_uuid,store_id,created_at) 
	VALUES ($1,$2,$3,$4) RETURNING id`, d.tableName)
	param := []interface{}{model.LicenseKey.String, model.MachineUUID.String, model.StoreID.Int64, model.CreatedAt.Time}

	err = tx.QueryRow(query, param...).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	return
}

func (d *licenseDAO) CheckExistingMachine(tx *sql.Tx, machineUUID string) (result int64, err error) {
	query := fmt.Sprintf(`SELECT id FROM %s 
	WHERE machine_uuid=$1`, d.tableName)
	param := []interface{}{machineUUID}
	err = tx.QueryRow(query, param...).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	return
}

func (d *licenseDAO) CheckExistingLicense(tx *sql.Tx, licenseKey string) (result int64, err error) {
	query := fmt.Sprintf(`SELECT id FROM %s 
	WHERE license_key=$1`, d.tableName)
	param := []interface{}{licenseKey}
	err = tx.QueryRow(query, param...).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	return
}

func (d *licenseDAO) CheckLicenseAndMachine(model repository.LicenseModel) (result bool, err error) {
	query := fmt.Sprintf(`SELECT EXISTS(
	SELECT 1 FROM %s WHERE machine_uuid=$1, license_key=$2, store_id=$3)`, d.tableName)
	param := []interface{}{model.MachineUUID.String, model.LicenseKey.String, model.StoreID.Int64}
	err = d.db.QueryRow(query, param...).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}

	}
	return
}

func (d *licenseDAO) ViewDetail(id int64) (result repository.LicenseModel, err error) {
	query := fmt.Sprintf(`SELECT id, machine_uuid, license_key, store_id,created_at FROM %s WHERE id=$1`, d.tableName)
	param := []interface{}{id}
	err = d.db.QueryRow(query, param...).Scan(&result.ID, &result.MachineUUID, &result.LicenseKey, &result.StoreID, &result.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	return
}
