package mogrify

import (
  "bytes"
  "io"
)

func EncodeJpg(w io.Writer, img Image) (int64, error) {
  slice, err := img.image().gdImageJpeg()
  if err != nil {
    return 0, err
  }

  return bytes.NewBuffer(slice).WriteTo(w)
}

func EncodePng(w io.Writer, img Image) (int64, error) {
  slice, err := img.image().gdImagePng()
  if err != nil {
    return 0, err
  }

  return bytes.NewBuffer(slice).WriteTo(w)
}

func EncodeGif(w io.Writer, img Image) (int64, error) {
  slice, err := img.image().gdImageGif()
  if err != nil {
    return 0, err
  }

  return bytes.NewBuffer(slice).WriteTo(w)
}
