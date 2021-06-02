package model

// File contains a file's information.
type File struct {
	RedisNil         bool   `redis:"redisNil"`
	ID               int64  `redis:"id"`
	OwnerType        string `redis:"ownerType"`
	OwnerID          int64  `redis:"ownerID"`
	Category         string `redis:"category"`
	Filename         string `redis:"filename"`
	OriginalFilename string `redis:"oriFilename"`
	MediaType        string `redis:"mediaType"`
	FileExt          string `redis:"fileExt"`
	FileSize         int64  `redis:"fileSize"`
	Width            int    `redis:"width"`
	Height           int    `redis:"height"`
	ThumbMediaType   string `redis:"thumbMediaType"`
	ThumbFileExt     string `redis:"thumbFileExt"`
	ThumbFileSize    int64  `redis:"thumbFileSize"`
	ThumbWidth       int    `redis:"thumbWidth"`
	ThumbHeight      int    `redis:"thumbHeight"`
	Storage          string `redis:"storage"`
	IsEncrypted      bool   `redis:"isEncrypted"`
	EncryptKey       string `redis:"encryptKey"`
	Uploader         string `redis:"uploader"`
	CreatedTime      int64  `redis:"createdTime"`
	UpdatedTime      int64  `redis:"updatedTime"`
	DeletedTime      int64  `redis:"deletedTime"`
}
