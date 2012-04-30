package mogrify

// #cgo CFLAGS: -I/usr/local/include/GraphicsMagick
// #cgo LDFLAGS: -lGraphicsMagickWand -lGraphicsMagick
// #include <wand/magick_wand.h>
import "C"

import (
  "errors"
  "fmt"
  "unsafe"
)

var (
  CannotOpen   = errors.New("Cannot open file")
  ResizeFailed = errors.New("Resize operation failed")
)

func Print(s string) {
  cs := C.CString(s)
  C.fputs(cs, (*C.FILE)(C.stdout))
  C.free(unsafe.Pointer(cs))
}

type Image struct {
  wand     *C.MagickWand
  filename string
}

type ImageError struct {
  message  string
  severity int
}

func (e *ImageError) Error() string {
  return fmt.Sprintf("GraphicsMagick: %s severity: %d", e.message, e.severity)
}

func init() {
  C.InitializeMagick(nil)
}

func (img *Image) exception() error {
  var ex C.ExceptionType

  message := C.GoString(C.MagickGetException(img.wand, &ex))

  return &ImageError{message, int(ex)}
}

func NewImage() *Image {
  image := new(Image)
  image.wand = C.NewMagickWand()
  return image
}

func Open(filename string) *Image {
  image := NewImage()

  if image.OpenFile(filename) == nil {
    return image
  }
  return nil
}

func (img *Image) OpenFile(filename string) error {
  status := C.MagickReadImage(img.wand, C.CString(filename))
  if status == C.MagickFalse {
    return CannotOpen
  }
  return nil
}

func (img *Image) Resize(width, height uint64) error {
  res := C.MagickResizeImage(img.wand, C.ulong(width), C.ulong(height), C.GaussianFilter, 1)

  if res != 1 {
    return img.exception()
  }
  return nil
}

func (img *Image) SaveFile(filename string) bool {
  status := C.MagickWriteImage(img.wand, C.CString(filename))
  if status == C.MagickFalse {
    return false
  }
  return true
}

func (img *Image) Destroy() {
  if img.wand != nil {
    C.DestroyMagickWand(img.wand)
  }
}
