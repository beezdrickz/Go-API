package initiate_dependency

import (
	"database/sql"
	"go-api-frame/dao"
)

type DAOContainer struct {
	LicenseDAO dao.LicenseDao
}

func InitDAO(db *sql.DB) *DAOContainer {
	return &DAOContainer{
		LicenseDAO: dao.NewLicenseDAO(db),
	}
}
