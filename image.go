package mogrify

import (
	"fmt"
)

type Image interface {
	Width() int
	Height() int
}

func Dimensions(img Image) string {
	return fmt.Sprintf("%dx%d", img.Width(), img.Height())
}
