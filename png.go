package mogrify

import (
	"bytes"
	"fmt"
	"io"
)

type Png struct {
	// Import GdImage and all it's methods
	GdImage
}

func DecodePng(reader io.Reader) (Image, error) {
	var image Png

	image.gd = gdCreateFromPng(drain(reader))
	if image.gd == nil {
		return nil, fmt.Errorf("couldn't create PNG decoder")
	}

	return &image, nil
}

func EncodePng(w io.Writer, img Image) (int64, error) {
	slice, err := img.image().gdImagePng()
	if err != nil {
		return 0, err
	}

	return bytes.NewBuffer(slice).WriteTo(w)
}
