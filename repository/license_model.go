package repository

import "database/sql"

type LicenseModel struct {
	ID          sql.NullInt64
	StoreID     sql.NullInt64
	MachineUUID sql.NullString
	LicenseKey  sql.NullString
	CreatedAt   sql.NullTime
}
