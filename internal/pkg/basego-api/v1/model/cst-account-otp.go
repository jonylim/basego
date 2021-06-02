package model

// CstAccountOTP contains an OTP's information.
type CstAccountOTP struct {
	RedisNil           bool   `json:"-" redis:"redisNil"`
	ID                 int64  `json:"-" redis:"id"`
	AccountID          int64  `json:"-" redis:"accountID"`
	Key                string `json:"-" redis:"key"`
	Code               string `json:"-" redis:"code"`
	Action             string `json:"-" redis:"action"`
	Method             string `json:"-" redis:"method"`
	Email              string `json:"-" redis:"email"`
	CountryID          int32  `json:"-" redis:"countryID"`
	CountryCallingCode string `json:"-" redis:"countryCallingCode"`
	Phone              string `json:"-" redis:"phone"`
	PhoneWithCode      string `json:"-" redis:"phoneWithCode"`
	ExpiryTime         int64  `json:"-" redis:"expiryTime"`
	SendCount          int    `json:"-" redis:"sendCount"`
	AttemptCount       int    `json:"-" redis:"attemptCount"`
	IsVerified         bool   `json:"-" redis:"isVerified"`
	CreatedTime        int64  `json:"-" redis:"createdTime"`
	UpdatedTime        int64  `json:"-" redis:"updatedTime"`
	DeletedTime        int64  `json:"-" redis:"deletedTime"`
}
