package mogrify

import (
	"fmt"
	"io"
)

type decodeFunc func(io.Reader) (Image, error)
type encodeFunc func(io.Writer, Image) (int64, error)

var (
	encoders = make(map[string]encodeFunc)
	decoders = make(map[string]decodeFunc)
)

func registerFormat(mime string, e encodeFunc, d decodeFunc) {
	encoders[mime] = e
	decoders[mime] = d
}

func init() {
	registerFormat("image/png", EncodePng, DecodePng)
	registerFormat("image/jpeg", EncodeJpeg, DecodeJpeg)
	registerFormat("image/jpg", EncodeJpeg, DecodeJpeg)
	registerFormat("image/gif", EncodeGif, DecodeGif)
}

// Encode an image onto a writer using an encoder appropriate for the
// given mimetype, if one exists.
func Encode(mime string, w io.Writer, i Image) (int64, error) {
	encoder, ok := encoders[mime]
	if !ok {
		return 0, fmt.Errorf("no encoder for mime type '%s'", mime)
	}
	return encoder(w, i)
}

// Decode an image from the reader using a decoder appropriate for the
// given mimetype, if one exists.
func Decode(mime string, r io.Reader) (Image, error) {
	decoder, ok := decoders[mime]
	if !ok {
		return nil, fmt.Errorf("no decoder for mime type '%s'", mime)
	}
	return decoder(r)
}
