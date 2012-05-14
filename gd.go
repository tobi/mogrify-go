package mogrify

// #cgo LDFLAGS: -lgd
// #include <gd.h>
import "C"

import (
	//	"bytes"
	// "fmt"
	"errors"
	"unsafe"
)

var (
	imageError = errors.New("image is nil")
	writeError = errors.New("image cannot be accessed")
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
	return p == nil
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

func (p *gdImage) gdCopyResampled(dst *gdImage, dstX, dstY, srcX, srcY, dstW, dstH, srcW, srcH int) {
	if p == nil || p.img == nil {
		panic(imageError)
	}

	C.gdImageCopyResampled(dst.img, p.img, C.int(dstX), C.int(dstY), C.int(srcX), C.int(srcY),
		C.int(dstW), C.int(dstH), C.int(srcW), C.int(srcH))
}

func (p *gdImage) gdCopyResized(dst *gdImage, dstX, dstY, srcX, srcY, dstW, dstH, srcW, srcH int) {
	if p == nil {
		panic(imageError)
	}
	C.gdImageCopyResized(dst.img, p.img, C.int(dstX), C.int(dstY), C.int(srcX), C.int(srcY),
		C.int(dstW), C.int(dstH), C.int(srcW), C.int(srcH))
}

// func (p *gdImage) gdImagePng() ([]byte, error) {
// 	var size int

// 	data := C.gdImagePngPtr(p.img, &size)
// 	if data == nil {
// 		return []byte{}, writeError
// 	}

// 	defer C.gdFree(unsafe.Pointer(data))	

// 	return C.GoBytes(data)
// }

func (p *gdImage) gdImageJpeg() ([]byte, error) {
	var size C.int

	// use -1 as quality, this will mean to use standard Jpeg quality
	data := C.gdImageJpegPtr(p.img, &size, -1)
	if data == nil || int(size) == 0 {
		return []byte{}, writeError
	}
	
	defer C.gdFree(unsafe.Pointer(data))	

	return C.GoBytes(data, size), nil
}
