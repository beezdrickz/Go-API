package dto

import (
	"time"
)

type PayloadResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp`
	Data      Response  `json:"data"`
}

type Response struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
