package mogrify

import (
	"bytes"
	"fmt"
	"io"
)

type Jpeg struct {
	// Import GdImage and all it's methods
	GdImage
}

func DecodeJpeg(reader io.Reader) (Image, error) {
	var image Jpeg

	image.gd = gdCreateFromJpeg(drain(reader))
	if image.gd == nil {
		return nil, fmt.Errorf("couldn't create JPEG decoder")
	}

	return &image, nil
}

func EncodeJpeg(w io.Writer, img Image) (int64, error) {
	slice, err := img.image().gdImageJpeg()
	if err != nil {
		return 0, err
	}

	return bytes.NewBuffer(slice).WriteTo(w)
}
