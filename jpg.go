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
	gd *gdImage
}

func NewJpg(reader io.Reader) *Jpg {
	var image Jpg
	if _, err := image.ReadFrom(reader); err != nil {
		return nil
	}
	return &image
}

func NewBlankJpg(width, height int) *Jpg {
	var image Jpg
	image.gd = gdCreate(width, height)
	if image.gd == nil {
		return nil
	}

	return &image
}

func (img *Jpg) ReadFrom(reader io.Reader) (n int64, err error) {
	var buffer bytes.Buffer

	n, err = buffer.ReadFrom(reader)
	if err != nil {
		return
	}

	gd := gdCreateFromJpeg(buffer.Bytes())
	if gd == nil {
		return n, loadError
	}

	img.Destroy()
	img.gd = gd
	return
}

func (img *Jpg) WriteTo(writer io.Writer) (n int64, err error) {
	slice, err := img.gd.gdImageJpeg() 
	if err != nil {
		return 0, err
	}


	_, err = writer.Write(slice) 

	return 0, err
}

func (img *Jpg) Width() int {
	return img.gd.width()
}

func (img *Jpg) Height() int {
	return img.gd.height()
}

func (img *Jpg) NewResized(width, height int) (*Jpg, error) {
	resized := img.gd.gdCopyResized(0, 0, 0, 0, width, height, img.gd.width(), img.gd.height())
	if resized == nil {
		return nil, resampleError
	}

	return &Jpg{resized}, nil
}

func (img *Jpg) NewResampled(width, height int) (*Jpg, error) {
	resized := img.gd.gdCopyResampled(0, 0, 0, 0, width, height, img.gd.width(), img.gd.height())
	if resized == nil {
		return nil, resampleError
	}

	return &Jpg{resized}, nil
}

func (img *Jpg) Destroy() {
	img.gd.gdDestroy()
}
