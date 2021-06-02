package model

// CstAccountTOS contains a customer account's Terms of Service status.
type CstAccountTOS struct {
	RedisNil    bool  `json:"-"           redis:"redisNil"`
	ID          int64 `json:"id"          redis:"id"`
	AccountID   int64 `json:"accountID"   redis:"accountID"`
	CreatedTime int64 `json:"createdTime" redis:"createdTime"`
	UpdatedTime int64 `json:"updatedTime" redis:"updatedTime"`
	DeletedTime int64 `json:"deletedTime" redis:"deletedTime"`
}
