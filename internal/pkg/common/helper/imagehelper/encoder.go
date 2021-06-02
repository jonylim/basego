package imagehelper

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

// EncodeImage encodes the given image as the specified media type.
func EncodeImage(img image.Image, mediaType string) ([]byte, error) {
	var buf bytes.Buffer
	var err error
	switch mediaType {
	case "image/jpeg":
		err = jpeg.Encode(&buf, img, nil)
	case "image/png":
		err = png.Encode(&buf, img)
	case "image/gif":
		err = gif.Encode(&buf, img, nil)
	default:
		err = fmt.Errorf("Media type '%s' is not invalid", mediaType)
	}
	return buf.Bytes(), err
}
