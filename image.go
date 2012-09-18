package mogrify

import (
	"errors"
)

var (
	resampleError	= errors.New("Resampling failed")
	resizeError	= errors.New("Resampling failed")
	loadError	= errors.New("Image cannot be loaded")
)

type Image interface {
	Bounds() Bounds

	Destroy()

	NewResampled(bounds Bounds) (*GdImage, error)
	NewResized(bounds Bounds) (*GdImage, error)
  NewCropped(x int, y int, bounds Bounds) (*GdImage, error)

	image() *gdImage
}

func Dimensions(img Image) string {
	return img.Bounds().String()
}
