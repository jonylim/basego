package model

// CstAccountSession contains details of an account's session.
type CstAccountSession struct {
	RedisNil    bool   `redis:"redisNil"`
	ID          int64  `redis:"id"`
	AccountID   int64  `redis:"accountID"`
	Platform    string `redis:"platform"`
	DeviceModel string `redis:"deviceModel"`
	DeviceID    string `redis:"deviceID"`
	UserAgent   string `redis:"userAgent"`
	IPAddress   string `redis:"ipAddr"`
	LogoutTime  int64  `redis:"logoutTime"`
	CreatedTime int64  `redis:"createdTime"`
	UpdatedTime int64  `redis:"updatedTime"`
	DeletedTime int64  `redis:"deletedTime"`
}

// CstAccountSessionToken contains details of an access token & refresh token.
type CstAccountSessionToken struct {
	RedisNil           bool   `redis:"redisNil"`
	ID                 int64  `redis:"id"`
	SessionID          int64  `redis:"sessionID"`
	AccessToken        string `redis:"accessToken"`
	AccessTokenExpiry  int64  `redis:"accessTokenExpiry"`
	RefreshToken       string `redis:"refreshToken"`
	RefreshTokenExpiry int64  `redis:"refreshTokenExpiry"`
	CreatedTime        int64  `redis:"createdTime"`
	DeletedTime        int64  `redis:"deletedTime"`
}
