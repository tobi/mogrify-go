package mogrify

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

var (
	BlobEmpty = errors.New("blob was empty")
)

var (
	jpgLoadError  = errors.New("Jpeg cannot be loaded")
	resampleError = errors.New("Resampling failed")
)

type Jpg struct {
	gd *gdImage
}

func NewJpg(reader io.Reader) *Jpg {
	var image Jpg
	if _, err := image.ReadFrom(reader); err != nil {
		fmt.Printf("%s", err)
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
		return n, jpgLoadError
	}

	if img.gd != nil {
		img.Destroy()
	}

	img.gd = gd
	return
}

func (img *Jpg) WriteTo(writer io.Writer) (n int64, err error) {
	return 0, nil
}

func (img *Jpg) Width() int {
	return img.gd.width()
}

func (img *Jpg) Height() int {
	return img.gd.height()
}

func (img *Jpg) NewResized(width, height int) (*Jpg, error) {
	resized := NewBlankJpg(width, height)
	img.gd.gdCopyResized(resized.gd, 0, 0, 0, 0, width, height, img.Width(), img.Height())

	return resized, nil
}

func (img *Jpg) NewResampled(width, height int) (*Jpg, error) {
	resized := NewBlankJpg(width, height)
	if resized == nil {
		return nil, resampleError
	}

	img.gd.gdCopyResampled(resized.gd, 0, 0, 0, 0, width, height, img.gd.width(), img.gd.height())

	if isInvalid(resized.gd) {
		return nil, resampleError
	}

	return resized, nil
}

func (img *Jpg) Destroy() {
	img.gd.gdDestroy()
}
