package mogrify

import ()

type GdImage struct {
	gd *gdImage
}

func NewImage(width, height int) Image {
	var image GdImage
	image.gd = gdCreate(width, height)
	if image.gd == nil {
		return nil
	}

	return &image
}

func (img *GdImage) NewResized(bounds Bounds) (*GdImage, error) {

	if bounds == (Bounds{0, 0}) {
		return nil, resampleError
	}

	if bounds.Width == 0 {
		bounds.Width = img.image().width() * bounds.Height / img.image().height()
	}
	if bounds.Height == 0 {
		bounds.Height = img.image().height() * bounds.Width / img.image().width()
	}

	resized := img.image().gdCopyResized(0, 0, 0, 0, bounds.Width, bounds.Height, img.image().width(), img.image().height())
	if resized == nil {
		return nil, resampleError
	}

	return &GdImage{resized}, nil
}

func (img *GdImage) NewResampled(bounds Bounds) (*GdImage, error) {
	resized := img.image().gdCopyResampled(0, 0, 0, 0, bounds.Width, bounds.Height, img.image().width(), img.image().height())
	if resized == nil {
		return nil, resampleError
	}

	return &GdImage{resized}, nil
}

func (img *GdImage) Bounds() Bounds {
	return Bounds{img.image().width(), img.image().height()}
}

func (img *GdImage) Height() int {
	return img.image().height()
}

func (img *GdImage) image() *gdImage {
	return img.gd
}

func (img *GdImage) Destroy() {
	img.image().gdDestroy()
}
