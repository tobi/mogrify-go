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

type Format int

const (
  PNG Format = iota
  GIF
  JPG
)

var (
  JpgLoadError  = errors.New("Jpeg cannot be loaded")
  resampleError = errors.New("Resampling failed")
)

type Image struct {
  format Format
  gd     *gdImage
  //  ReadFrom func(r Reader) (n int64, err error)
}

func NewImage(format Format) *Image {
  return &Image{}
}

func (img *Image) ReadFromJpeg(reader io.Reader) (n int64, err error) {
  var buffer bytes.Buffer
  n, err = buffer.ReadFrom(reader)

  if err != nil {
    return
  }

  gd := gdCreateFromJpeg(buffer.Bytes())
  if gd == nil {
    return n, JpgLoadError
  }
  img.Destroy()
  img.format = JPG
  img.gd = gd

  return
}

func (img *Image) WriteTo(writer io.Writer) (n int64, err error) {
  return 0, nil
}

func (img *Image) Width() int {
  return img.gd.width()
}

func (img *Image) Height() int {
  return img.gd.height()
}

func (img *Image) Dimensions() string {
  return fmt.Sprintf("%dx%d", img.Width(), img.Height())
}

func (img *Image) CopyResized(width, height int) (*Image, error) {
  dst := gdCreate(width, height)
  img.gd.gdCopyResized(dst, 0, 0, 0, 0, width, height, img.gd.width(), img.gd.height())

  return &Image{img.format, dst}, nil
}

func (img *Image) CopyResampled(width, height int) (*Image, error) {
  dst := gdCreate(width, height)
  img.gd.gdCopyResampled(dst, 0, 0, 0, 0, width, height, img.gd.width(), img.gd.height())

  if dst.invalid() {
    return nil, resampleError
  }

  return &Image{img.format, dst}, nil
}

func (img *Image) Destroy() {
  img.gd.gdDestroy()
}
