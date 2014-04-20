package mogrify

import (
	"fmt"
)

// Image can be used to transform existing images.
type Image interface {
	Bounds() Bounds
	Destroy()
	NewResampled(bounds Bounds) (*GdImage, error)
	NewResized(bounds Bounds) (*GdImage, error)
	NewCropped(x int, y int, bounds Bounds) (*GdImage, error)

	image() *gdImage
}

// Dimensions of an image, as a string of the form:
//		NxM
// where N is the width, M the height, of the image.
func Dimensions(img Image) string {
	return fmt.Sprintf("%dx%d", img.Bounds().Width, img.Bounds().Height)
}
