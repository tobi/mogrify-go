package mogrify

import (
	"bytes"
	"io"
)

type Gif struct {
	// Import GdImage and all it's methods
	GdImage
}

func DecodeGif(reader io.Reader) Image {
	var image Gif

	image.gd = gdCreateFromGif(drain(reader))
	if image.gd == nil {
		return nil
	}

	return &image
}

func EncodeGif(w io.Writer, img Image) (int64, error) {
	slice, err := img.image().gdImageGif()
	if err != nil {
		return 0, err
	}

	return bytes.NewBuffer(slice).WriteTo(w)
}
