package mogrify

// #cgo LDFLAGS: -lgd
// #include <gd.h>
import "C"

import (
	//	"bytes"
	//	"fmt"
	"errors"
	"unsafe"
)

var (
	imageError  = errors.New("[GD] image is nil")
	createError = errors.New("[GD] cannot create new image")
	writeError  = errors.New("[GD] image cannot be written")
)

type gdImage struct {
	img *C.gdImage
}

func img(img *C.gdImage) *gdImage {
	if img == nil {
		return nil
	}

	image := &gdImage{img}
	if isInvalid(image) {
		return nil
	}
	return image
}

func gdCreate(sx, sy int) *gdImage {
	return img(C.gdImageCreateTrueColor(C.int(sx), C.int(sy)))
}

func gdCreateFromJpeg(buffer []byte) *gdImage {
	return img(C.gdImageCreateFromJpegPtr(C.int(len(buffer)), unsafe.Pointer(&buffer[0])))
}

func gdCreateFromGif(buffer []byte) *gdImage {
	return img(C.gdImageCreateFromGifPtr(C.int(len(buffer)), unsafe.Pointer(&buffer[0])))
}

func gdCreateFromPng(buffer []byte) *gdImage {
	return img(C.gdImageCreateFromGifPtr(C.int(len(buffer)), unsafe.Pointer(&buffer[0])))
}

func (p *gdImage) gdDestroy() {
	if p != nil && p.img != nil {
		C.gdImageDestroy(p.img)
	}
}

func isInvalid(p *gdImage) bool {
	return p == nil || p.img.pixels != nil
}

func (p *gdImage) width() int {
	if p == nil {
		panic(imageError)
	}
	return int((*p.img).sx)
}

func (p *gdImage) height() int {
	if p == nil {
		panic(imageError)
	}
	return int((*p.img).sy)
}

func (p *gdImage) gdCopyResampled(dstX, dstY, srcX, srcY, dstW, dstH, srcW, srcH int) *gdImage {
	if p == nil || p.img == nil {
		panic(imageError)
	}

	dst := gdCreate(dstW, dstH)

	if dst == nil {
		return nil
	}

	C.gdImageCopyResampled(dst.img, p.img, C.int(dstX), C.int(dstY), C.int(srcX), C.int(srcY),
		C.int(dstW), C.int(dstH), C.int(srcW), C.int(srcH))

	if isInvalid(dst) {
		dst.gdDestroy()
		return nil
	}

	return dst
}

func (p *gdImage) gdCopyResized(dstX, dstY, srcX, srcY, dstW, dstH, srcW, srcH int) *gdImage {
	if p == nil || p.img == nil {
		panic(imageError)
	}

	dst := gdCreate(dstW, dstH)

	if dst == nil {
		return nil
	}

	C.gdImageCopyResized(dst.img, p.img, C.int(dstX), C.int(dstY), C.int(srcX), C.int(srcY),
		C.int(dstW), C.int(dstH), C.int(srcW), C.int(srcH))

	if isInvalid(dst) {
		dst.gdDestroy()
		return nil
	}

	return dst
}

func (p *gdImage) gdImagePng() ([]byte, error) {
	if p == nil {
		panic(imageError)
	}

	var size C.int

	data := C.gdImagePngPtr(p.img, &size)
	if data == nil || int(size) == 0 {
		return []byte{}, writeError
	}

	defer C.gdFree(unsafe.Pointer(data))

	return C.GoBytes(data, size), nil
}

func (p *gdImage) gdImageJpeg() ([]byte, error) {
	if p == nil {
		panic(imageError)
	}

	var size C.int

	// use -1 as quality, this will mean to use standard Jpeg quality
	data := C.gdImageJpegPtr(p.img, &size, -1)
	if data == nil || int(size) == 0 {
		return []byte{}, writeError
	}

	defer C.gdFree(unsafe.Pointer(data))

	return C.GoBytes(data, size), nil
}
