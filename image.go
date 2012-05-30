package mogrify

import (
  "errors"
	"fmt"
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
