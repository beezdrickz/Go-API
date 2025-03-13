package in

import "errors"

type LicenseDtoRequest struct {
	ID          int64  `json:"id"`
	MachineUUID string `json:"machine_uuid"`
	PublicKey   string `json:"public_key"`
	StoreID     int64  `json:"store_id"`
}

func (l *LicenseDtoRequest) ValidateInsert() (err error) {
	if l.MachineUUID == "" {
		err = errors.New("Machine UUID mustn't empty")
		return
	}

	if l.PublicKey == "" {
		err = errors.New("Public Key mustn't empty")
		return
	}

	if l.StoreID == 0 {
		err = errors.New("Store ID mustn't empty")
		return
	}
	return
}
