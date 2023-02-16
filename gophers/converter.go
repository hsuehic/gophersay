package gopher

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/qeesung/image2ascii/convert"
)

func Convert(path string) string {
	// create convert options
	convertOptions := convert.DefaultOptions

	convertOptions.FixedHeight = 70
	convertOptions.FixedWidth = 70
	convertOptions.Reversed = true

	// Create the image converter
	converter := convert.NewImageConverter()
	return converter.ImageFile2ASCIIString(path, &convertOptions)
}
