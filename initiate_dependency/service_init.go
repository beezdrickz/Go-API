package initiate_dependency

import (
	"database/sql"
	"go-api-frame/services/license"
)

type ServiceContainer struct {
	LicenseService license.LicenseService
}

func InitService(db *sql.DB, dao *DAOContainer) *ServiceContainer {
	return &ServiceContainer{
		LicenseService: license.NewLicenseService(db, dao.LicenseDAO),
	}
}
