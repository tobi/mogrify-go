package mogrify

import (
	"bytes"
	"fmt"
	"io"
)

type Gif struct {
	// Import GdImage and all it's methods
	GdImage
}

func DecodeGif(reader io.Reader) (Image, error) {
	var image Gif

	image.gd = gdCreateFromGif(drain(reader))
	if image.gd == nil {
		return nil, fmt.Errorf("couldn't create GIF decoder")
	}

	return &image, nil
}

func EncodeGif(w io.Writer, img Image) (int64, error) {
	slice, err := img.image().gdImageGif()
	if err != nil {
		return 0, err
	}

	return bytes.NewBuffer(slice).WriteTo(w)
}
