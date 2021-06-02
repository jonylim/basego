package imagehelper

import (
	"bytes"
	"image"
	"io"

	// Init image decoder.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	// Init image decoder.
	_ "golang.org/x/image/webp"
)

// GetImageAndConfig returns an image and its config decoded from the reader.
func GetImageAndConfig(reader io.Reader) (img image.Image, cfg image.Config, err error) {
	var buf bytes.Buffer
	r := io.TeeReader(reader, &buf)

	tmpImg, _, err1 := image.Decode(r)
	tmpCfg, _, err2 := image.DecodeConfig(&buf)
	if err1 != nil {
		err = err1
	} else if err2 != nil {
		err = err2
	} else {
		img = tmpImg
		cfg = tmpCfg
		// contentType := fileType + "/" + format
	}
	return
}
