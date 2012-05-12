package mogrify

// #cgo LDFLAGS: -lgd
// #include <gd.h>
import "C"

import (
	//	"bytes"
	"unsafe"
)

type gdImage struct {
	img *C.gdImage
}

func img(img *C.gdImage) *gdImage {
	image := &gdImage{img}
	if image.invalid() {
		return nil
	}
	return image
}

func assertGd(img *C.gdImage) bool {
	return (*img).pixels != nil
}

func gdCreate(sx, sy int) *gdImage {
	return img(C.gdImageCreate(C.int(sx), C.int(sy)))
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

func (p *gdImage) invalid() bool {
	return p == nil || (*p.img).pixels == nil
}

func (p *gdImage) width() int {
	return int((*p.img).sx)
}

func (p *gdImage) height() int {
	return int((*p.img).sy)
}

func (p *gdImage) gdCopyResampled(dst *gdImage, dstX, dstY, srcX, srcY, dstW, dstH, srcW, srcH int) {
	C.gdImageCopyResampled(dst.img, p.img, C.int(dstX), C.int(dstY), C.int(srcX), C.int(srcY),
		C.int(dstW), C.int(dstH), C.int(srcW), C.int(srcH))
}

func (p *gdImage) gdCopyResized(dst *gdImage, dstX, dstY, srcX, srcY, dstW, dstH, srcW, srcH int) {
	C.gdImageCopyResized(dst.img, p.img, C.int(dstX), C.int(dstY), C.int(srcX), C.int(srcY),
		C.int(dstW), C.int(dstH), C.int(srcW), C.int(srcH))
}
