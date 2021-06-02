package model

// XAPIKey contains details of an API key for validating API requests.
type XAPIKey struct {
	ID            int32  `redis:"id"`
	APIKeyID      string `redis:"keyID"`
	APIKeySecret  string `redis:"keySec"`
	Domain        string `redis:"domain"`
	AppPlatform   string `redis:"appPlatform"`
	AppIdentifier string `redis:"appIdentifier"`
	ExpiryTime    int64  `redis:"expiryTime"`
	IsEnabled     bool   `redis:"isEnabled"`
	CreatedTime   int64  `redis:"createdTime"`
	UpdatedTime   int64  `redis:"updatedTime"`
	DeletedTime   int64  `redis:"deletedTime"`
}
