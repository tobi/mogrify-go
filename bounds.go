package mogrify

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var boundFinder = regexp.MustCompile("([0-9]*)x([0-9]*)")

// Bounds of an image.
type Bounds struct {
	Width, Height int
}

// BoundsFromString creates a Bounds from strings of the form:
//	"100x150"   ->	Width 100, Height 150
//	"x150"      ->	Width 0, Height 150
//	"100x"      ->	Width 100, Height 0
//	"x"         ->	Width 0, Height 0
//	"no match"  ->	error
func BoundsFromString(bounds string) (*Bounds, error) {

	dimensions := boundFinder.FindStringSubmatch(bounds)
	if len(dimensions) != 3 {
		return nil, fmt.Errorf("malformed bound string")
	}

	atoiOrZero := func(str string) int {
		if str == "" {
			return 0
		}
		val, _ := strconv.Atoi(str)
		return val
	}

	return &Bounds{
		Width:  atoiOrZero(dimensions[1]),
		Height: atoiOrZero(dimensions[2]),
	}, nil
}

// ScaleProportionally the bounds to the smallest side.
func (b Bounds) ScaleProportionally(targetWidth, targetHeight int) Bounds {
	scalex := float64(targetWidth) / float64(b.Width)
	scaley := float64(targetHeight) / float64(b.Height)
	scale := math.Min(scalex, scaley)

	return Bounds{
		Width:  int(math.Floor(float64(b.Width) * scale)),
		Height: int(math.Floor(float64(b.Height) * scale)),
	}
}

// ShrinkProportionally the bounds only if both sides are larger than
// the target.
func (b Bounds) ShrinkProportionally(targetWidth, targetHeight int) Bounds {
	// Make sure there is work to be done
	if b.Width < targetWidth || b.Height < targetHeight {
		return b
	}

	return b.ScaleProportionally(targetWidth, targetHeight)
}

// GrowProportionally the bounds only if both sides are smaller than
// the target.
func (b Bounds) GrowProportionally(targetWidth, targetHeight int) Bounds {
	// Make sure there is work to be done
	if b.Width > targetWidth || b.Height > targetHeight {
		return b
	}

	return b.ScaleProportionally(targetWidth, targetHeight)
}
