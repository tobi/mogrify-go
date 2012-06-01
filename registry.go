package mogrify

import (
	"io"
	"errors"
)

type decodeFunc (func(io.Reader) Image)
type encodeFunc (func(io.Writer, Image) (int64, error))

var (
	noEncoder				= errors.New("no encoder for mime type")
	encoders	map[string]encodeFunc	= make(map[string]encodeFunc)
	decoders	map[string]decodeFunc	= make(map[string]decodeFunc)
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

func Encode(mime string, w io.Writer, i Image) (int64, error) {
	encoder := encoders[mime]

	if encoder == nil {
		return 0, noEncoder
	}

	return encoder(w, i)
}

func Decode(mime string, r io.Reader) Image {
	decoder := decoders[mime]

	if decoder == nil {
		return nil
	}

	return decoder(r)
}
