package mogrify

import (
  "bytes"
  "io"
)

type Png struct {
  // Import GdImage and all it's methods
  GdImage
}

func DecodePng(reader io.Reader) Image {
  var image Png

  image.gd = gdCreateFromPng(drain(reader))
  if image.gd == nil {
    return nil
  }

  return &image
}

func EncodePng(w io.Writer, img Image) (int64, error) {
  slice, err := img.image().gdImagePng()
  if err != nil {
    return 0, err
  }

  return bytes.NewBuffer(slice).WriteTo(w)
}
