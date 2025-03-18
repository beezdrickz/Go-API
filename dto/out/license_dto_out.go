package out

import "time"

type LicenseDTOOut struct {
	ID          int64     `json:"id"`
	MachineUUID string    `json:"machine_uuid" validate:"required" `
	PublicKey   string    `json:"public_key" validate:"required"`
	StoreID     int64     `json:"store_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}
