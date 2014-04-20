package mogrify

import (
	"bytes"
	"fmt"
	"io"
)

// Jpeg image that can be transformed.
type Jpeg struct {
	// Embed GdImage and all it's methods
	GdImage
}

// DecodeJpeg decodes a JPEG image from a reader.
func DecodeJpeg(reader io.Reader) (Image, error) {
	var image Jpeg

	image.gd = gdCreateFromJpeg(drain(reader))
	if image.gd == nil {
		return nil, fmt.Errorf("couldn't create JPEG decoder")
	}

	return &image, nil
}

// EncodeJpeg encodes the image onto the writer as a JPEG.
func EncodeJpeg(w io.Writer, img Image) (int64, error) {
	slice, err := img.image().gdImageJpeg()
	if err != nil {
		return 0, err
	}

	return bytes.NewBuffer(slice).WriteTo(w)
}
