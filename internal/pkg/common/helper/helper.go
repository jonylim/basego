package helper

var mapContentTypeFileExts = map[string]string{
	"image/bmp":       ".bmp",
	"image/jpeg":      ".jpg",
	"image/png":       ".png",
	"image/webp":      ".webp",
	"video/quicktime": ".mov",
	"video/mp4":       ".mp4",
	"application/pdf": ".pdf",
	"application/zip": ".zip",
}

// GetFileExtensionByMediaType returns file extension for a media type.
func GetFileExtensionByMediaType(mediaType string) string {
	if fileExt, ok := mapContentTypeFileExts[mediaType]; ok {
		return fileExt
	}
	return ""
}

var mapFileExtContentTypes = map[string]string{
	".bmp":  "image/bmp",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".webp": "image/webp",
	".mov":  "video/quicktime",
	".mp4":  "video/mp4",
	".pdf":  "application/pdf",
	".zip":  "application/zip",
}

// GetMediaTypeFromFilename returns media type by filename.
func GetMediaTypeFromFilename(filename string) string {
	return GetMediaTypeFromFileExt(GetFileExtension(filename))
}

// GetMediaTypeFromFileExt returns media type for a file extension.
func GetMediaTypeFromFileExt(ext string) string {
	if mediaType, ok := mapFileExtContentTypes[ext]; ok {
		return mediaType
	}
	return "application/octet-stream"
}

// CalculateThumbnailImageDimension calculates image dimension for thumbnail while keeping the width-height ratio.
func CalculateThumbnailImageDimension(originalWidth, originalHeight, thumbnailMaxWidth, thumbnailMaxHeight int) (newWidth, newHeight int) {
	newWidth, newHeight = originalWidth, originalHeight
	if newWidth > thumbnailMaxWidth {
		newHeight = int(newHeight * thumbnailMaxWidth / newWidth)
		if newHeight < 1 {
			newHeight = 1
		}
		newWidth = thumbnailMaxWidth
	}
	if newHeight > thumbnailMaxHeight {
		newWidth = int(newWidth * thumbnailMaxHeight / newHeight)
		if newWidth < 1 {
			newWidth = 1
		}
		newHeight = thumbnailMaxHeight
	}
	return
}
