package model

// ImageURL contains image URLs for both thumbnail and fullsize images.
type ImageURL struct {
	Thumbnail string `json:"thumbnail"`
	Fullsize  string `json:"fullsize"`
}
