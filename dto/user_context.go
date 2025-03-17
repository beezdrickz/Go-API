package dto

import "time"

type UserContext struct {
	ID            int64     `json:"id"`
	UserID        string    `json:"userId"`
	Username      string    `json:"username"`
	UserCompany   string    `json:"usercompany"`
	UserLogintime time.Time `json:"userlogintime"`
	UserIpAddr2   string    `json:"useripaddr2"`
	Role          string    `json:"role"`
	ConsignmentID int64     `json:"consignmentId"`
	Store         int64     `json:"store"`
	Permission    string    `json:"permission"`
	IsOtp         bool      `json:"istotp"`
	SessionID     string    `json:"sessionid"`
}
