package mogrify

import (
  "bytes"
  "errors"
	"fmt"
  "io"
)

var (
  resampleError = errors.New("Resampling failed")
  resizeError   = errors.New("Resampling failed")
  loadError     = errors.New("Image cannot be loaded")
)

type Image interface {
	Width() int
	Height() int
  Destroy()

  NewResampled(width, height int) (*GdImage, error)
  NewResized(width, height int) (*GdImage, error)

  image() *gdImage
}

func Dimensions(img Image) string {
	return fmt.Sprintf("%dx%d", img.Width(), img.Height())
}

func EncodeJpg(w io.Writer, img Image) (int64, error) {
  slice, err := img.image().gdImageJpeg()
  if err != nil {
    return 0, err
  }

  return bytes.NewBuffer(slice).WriteTo(w)
}
