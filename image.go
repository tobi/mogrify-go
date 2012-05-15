package mogrify

import (
	"fmt"
  "errors"
)

var (
  resampleError = errors.New("Resampling failed")
  resizeError = errors.New("Resampling failed")
  loadError  = errors.New("Image cannot be loaded")
)

type Image interface {
	Width() int
	Height() int
}

func Dimensions(img Image) string {
	return fmt.Sprintf("%dx%d", img.Width(), img.Height())
}