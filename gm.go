package mogrify

// #cgo CFLAGS: -I/usr/local/include/GraphicsMagick
// #cgo LDFLAGS: -lGraphicsMagickWand -lGraphicsMagick
// #include <wand/magick_wand.h>
import "C"

import (
  "errors"
  "fmt"
  "io"
  "unsafe"
)

var (
  BlobEmpty = errors.New("blob was empty")
)

type Image struct {
  wand     *C.MagickWand
  filename string
}

type ImageError struct {
  message  string
  severity int
}

func init() {
  C.InitializeMagick(nil)
}

func (e *ImageError) Error() string {
  return fmt.Sprintf("GraphicsMagick: %s severity: %d", e.message, e.severity)
}

func Open(filename string) *Image {
  image := NewImage()

  if image.OpenFile(filename) == nil {
    return image
  }
  return nil
}

func (img *Image) error() error {
  var ex C.ExceptionType

  char_ptr := C.MagickGetException(img.wand, &ex)
  defer C.free(unsafe.Pointer(char_ptr))

  return &ImageError{C.GoString(char_ptr), int(ex)}
}

func NewImage() *Image {
  image := new(Image)
  image.wand = C.NewMagickWand()
  return image
}

func (img *Image) OpenFile(filename string) error {
  cfilename := C.CString(filename)
  defer C.free(unsafe.Pointer(cfilename))

  status := C.MagickReadImage(img.wand, cfilename)
  if status == C.MagickFalse {
    return img.error()
  }
  return nil
}

func (img *Image) OpenBlob(bytes []byte) error {
  if len(bytes) < 1 {
    return BlobEmpty
  }

  status := C.MagickReadImageBlob(img.wand, (*C.uchar)(&bytes[0]), C.size_t(len(bytes)))

  if status == C.MagickFalse {
    return img.error()
  }
  return nil
}

func (img *Image) SaveBlob() ([]byte, error) {
  var len C.size_t
  char_ptr := C.MagickWriteImageBlob(img.wand, &len)

  if char_ptr == nil {
    return nil, img.error()
  }

  defer C.free(unsafe.Pointer(char_ptr))

  return C.GoBytes(unsafe.Pointer(char_ptr), C.int(len)), nil
}

func (img *Image) Write(writer io.Writer) (int, error) {
  bytes, err := img.SaveBlob()

  if err != nil {
    return 0, img.error()
  }

  return writer.Write(bytes)
}

func (img *Image) Resize(width, height uint) error {
  res := C.MagickResizeImage(img.wand, C.ulong(width), C.ulong(height), C.GaussianFilter, 1)

  if res == C.MagickFalse {
    return img.error()
  }
  return nil
}

func (img *Image) SaveFile(filename string) error {
  cfilename := C.CString(filename)
  defer C.free(unsafe.Pointer(cfilename))

  status := C.MagickWriteImage(img.wand, cfilename)
  if status == C.MagickFalse {
    return img.error()
  }
  return nil
}

func (img *Image) Destroy() {
  if img.wand != nil {
    C.DestroyMagickWand(img.wand)
  }
}