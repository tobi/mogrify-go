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

var (
  JpgLoadError  = errors.New("Jpeg cannot be loaded")
  resampleError = errors.New("Resampling failed")
)

type Image interface {
  Width() int
  Height() int  
}

func Dimensions(img Image) string {
  return fmt.Sprintf("%dx%d", img.Width(), img.Height())
}


type Jpg struct {  
  gd     *gdImage  
}

func NewJpg(reader io.Reader) *Jpg {
  var image Jpg
  image.ReadFrom(reader)
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
    return n, JpgLoadError
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
  dst := gdCreate(width, height)
  img.gd.gdCopyResized(dst, 0, 0, 0, 0, width, height, img.gd.width(), img.gd.height())

  return &Jpg{dst}, nil
}

func (img *Jpg) NewResampled(width, height int) (*Jpg, error) {
  dst := gdCreate(width, height)
  img.gd.gdCopyResampled(dst, 0, 0, 0, 0, width, height, img.gd.width(), img.gd.height())

  if dst.invalid() {
    return nil, resampleError
  }

  return &Jpg{dst}, nil
}

func (img *Jpg) Destroy() {
  img.gd.gdDestroy()
}
