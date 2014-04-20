package mogrify

import (
	"fmt"
)

// GdImage implements a wrapper around libgd.
type GdImage struct {
	gd *gdImage
}

// NewImage creates an image that can be modified.
func NewImage(width, height int) Image {
	var image GdImage
	image.gd = gdCreate(width, height)
	if image.gd == nil {
		return nil
	}

	return &image
}

// NewResized builds a copy of this image, resized to fit the new
// bounds.
func (img *GdImage) NewResized(bounds Bounds) (*GdImage, error) {
	bounds, err := calculateBounds(bounds, img)
	if err != nil {
		return nil, err
	}

	resized := img.image().gdCopyResized(0, 0, 0, 0, bounds.Width, bounds.Height, img.image().width(), img.image().height())
	if resized == nil {
		return nil, fmt.Errorf("cgo call to gdCopyResized failed")
	}

	return &GdImage{resized}, nil
}

// NewResampled builds a copy of this image, resampled within the new
// bounds.
func (img *GdImage) NewResampled(bounds Bounds) (*GdImage, error) {
	bounds, err := calculateBounds(bounds, img)
	if err != nil {
		return nil, err
	}

	resized := img.image().gdCopyResampled(0, 0, 0, 0, bounds.Width, bounds.Height, img.image().width(), img.image().height())
	if resized == nil {
		return nil, fmt.Errorf("cgo call to gdCopyResampled failed")
	}

	return &GdImage{resized}, nil
}

// NewCropped builds a copy of this image, cropped by the bounds.
func (img *GdImage) NewCropped(x int, y int, bounds Bounds) (*GdImage, error) {
	bounds, err := calculateBounds(bounds, img)
	if err != nil {
		return nil, err
	}

	cropped := img.image().gdCopy(0, 0, x, y, bounds.Width, bounds.Height)
	if cropped == nil {
		return nil, fmt.Errorf("cgo call to gdCopy failed")
	}

	return &GdImage{cropped}, nil
}

// Bounds within which this image holds.
func (img *GdImage) Bounds() Bounds {
	return Bounds{img.image().width(), img.image().height()}
}

func (img *GdImage) image() *gdImage {
	return img.gd
}

// Destroy cleans up the resources used by this image.
func (img *GdImage) Destroy() {
	img.image().gdDestroy()
}

func calculateBounds(bounds Bounds, img *GdImage) (Bounds, error) {
	if bounds == (Bounds{0, 0}) {
		return bounds, fmt.Errorf("both sides can't be of length 0")
	}

	if bounds.Width == 0 {
		bounds.Width = img.image().width() * bounds.Height / img.image().height()
	}

	if bounds.Height == 0 {
		bounds.Height = img.image().height() * bounds.Width / img.image().width()
	}

	return bounds, nil
}
