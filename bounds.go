package mogrify

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Bounds struct {
	Width, Height int
}

func BoundsFromString(bounds string) Bounds {
	var x, y int

	finder, err := regexp.Compile("([0-9]*)x([0-9]*)")
	if err != nil {
		return Bounds{0, 0}
	}

	dimensions := finder.FindStringSubmatch(bounds)

	x, _ = strconv.Atoi(dimensions[1])
	y, _ = strconv.Atoi(dimensions[2])

	return Bounds{x, y}
}

func (b Bounds) String() string {
	return fmt.Sprintf("%dx%d", b.Width, b.Height)
}

func (b Bounds) ScaleProportionally(targetWidth, targetHeight int) Bounds {
	scalex := float64(targetWidth) / float64(b.Width)
	scaley := float64(targetHeight) / float64(b.Height)
	scale := math.Min(scalex, scaley)

	return Bounds{
		Width:  int(math.Floor(float64(b.Width) * scale)),
		Height: int(math.Floor(float64(b.Height) * scale)),
	}
}

func (b Bounds) ShrinkProportionally(targetWidth, targetHeight int) Bounds {
	// Make sure there is work to be done
	if b.Width < targetWidth || b.Height < targetHeight {
		return b
	}

	return b.ScaleProportionally(targetWidth, targetHeight)
}

func (b Bounds) GrowProportionally(targetWidth, targetHeight int) Bounds {
	// Make sure there is work to be done
	if b.Width > targetWidth || b.Height > targetHeight {
		return b
	}

	return b.ScaleProportionally(targetWidth, targetHeight)
}
