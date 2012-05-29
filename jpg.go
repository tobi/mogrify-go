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
	GdImage
}

func DecodeJpg(reader io.Reader) Image {
	var image Jpg
	if _, err := readFromJpg(reader); err != nil {
		return nil
	}
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

	_, err = writer.Write(slice) 
	// todo: return actual len of write
	return 0, err
}

func (img *Jpg) Encode(w io.Writer) (int64, error) {

	slice, err := img.image().gdImageJpeg() 
	if err != nil {
		return 0, err
	}

	return bytes.NewBuffer(slice).WriteTo(w)	
}