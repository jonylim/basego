package model

// CstAccount contains a customer account's information.
type CstAccount struct {
	ID                    int64    `json:"id"`
	FullName              string   `json:"fullName"`
	Email                 string   `json:"email"`
	IsEmailVerified       bool     `json:"isEmailVerified"`
	CountryID             int32    `json:"countryID"`
	CountryCallingCode    string   `json:"countryCallingCode"`
	Phone                 string   `json:"phone"`
	PhoneWithCode         string   `json:"phoneWithCode"`
	IsPhoneVerified       bool     `json:"isPhoneVerified"`
	Password              string   `json:"-"`
	PasswordSalt          string   `json:"-"`
	Use2FA                bool     `json:"-"`
	ImageURL              ImageURL `json:"imageURL"`
	LastLoginTime         int64    `json:"lastLoginTime"`
	LastActivityTime      int64    `json:"lastActivityTime"`
	RequireChangePassword bool     `json:"requireChangePassword"`
	CreatedTime           int64    `json:"createdTime"`
	UpdatedTime           int64    `json:"updatedTime"`
	DeletedTime           int64    `json:"deletedTime"`
}

// ClearPersonalInfo clears the account's personal information, such as email address and phone number.
func (account *CstAccount) ClearPersonalInfo() {
	account.ClearEmailInfo()
	account.ClearPhoneInfo()
	account.ClearPasswordInfo()
}

// ClearEmailInfo clears the account's email address information.
func (account *CstAccount) ClearEmailInfo() {
	account.Email = ""
	account.IsEmailVerified = false
}

// ClearPhoneInfo clears the account's phone number information.
func (account *CstAccount) ClearPhoneInfo() {
	account.CountryCallingCode = ""
	account.Phone = ""
	account.PhoneWithCode = ""
	account.IsPhoneVerified = false
}

// ClearPasswordInfo clears the account's password information.
func (account *CstAccount) ClearPasswordInfo() {
	account.Password = ""
	account.PasswordSalt = ""
	account.Use2FA = false
	account.RequireChangePassword = false
}
