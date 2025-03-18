package out

import "time"

type LicenseDTOOut struct {
	ID          int64     `json:"id"`
	MachineUUID string    `json:"machine_uuid" `
	PublicKey   string    `json:"public_key"`
	StoreID     int64     `json:"store_id"`
	CreatedAt   time.Time `json:"created_at"`
}
