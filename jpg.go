package mogrify

import (
	"bytes"
	"errors"
	"io"
)

var (
	BlobEmpty = errors.New("blob was empty")
)

type Jpg struct {
	// Import GdImage and all it's methods
	GdImage
}

func DecodeJpg(reader io.Reader) Image {
	var image Jpg

	gd, err := readFromJpg(reader)

	if err != nil {
		return nil
	}

	image.gd = gd

	return &image
}

// Take an io.Reader and read a gdImage from it
func readFromJpg(reader io.Reader) (gd *gdImage, err error) {
	var buffer bytes.Buffer

	_, err = buffer.ReadFrom(reader)
	if err != nil {
		return
	}

	gd = gdCreateFromJpeg(buffer.Bytes())
	if gd == nil {
		return nil, loadError
	}

	return gd, nil
}

func writeAsJpg(img Image, writer io.Writer) (n int64, err error) {
	slice, err := img.image().gdImageJpeg()
	if err != nil {
		return 0, err
	}

	written, err := writer.Write(slice)

	// todo: return actual len of write
	return (int64)(written), err
}