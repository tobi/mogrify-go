package mogrify

import (
  "math"
  "fmt"
)

type Bounds struct {
  Width, Height int
}

func BoundsFromString(bounds string) Bounds {
  var x,y int
  fmt.Sscanf(bounds, "%dx%d", &x, &y)
  return Bounds{x, y}
}

func (b Bounds) String() string {
  return fmt.Sprintf("%dx%d", b.Width, b.Height)
}

func (b Bounds) ScaleProportionally(targetWidth, targetHeight int) Bounds {
  scalex := (float64)(targetWidth) / (float64)(b.Width)
  scaley := (float64)(targetHeight) / (float64)(b.Height)
  scale := math.Min(scalex, scaley)

  return Bounds{ (int)(math.Floor((float64)(b.Width) * scale)), (int)(math.Floor((float64)(b.Height) * scale)) }
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
