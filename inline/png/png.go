package png

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
)

// Size to reduce to.
const Size = 56

const inlinePrefix = "data:image/png;base64,"

func Inline(pngdata []byte) ([]byte, error) {
	// Compress first.
	b, err := compress(pngdata)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to compress PNG")
	}

	// Base64 encode.
	var b64dat = make([]byte, base64.URLEncoding.EncodedLen(len(b)))
	base64.URLEncoding.Encode(b64dat, b)

	return b64dat, nil
}

var encoder = png.Encoder{
	CompressionLevel: png.BestCompression,
}

func compress(pngdata []byte) ([]byte, error) {
	i, err := png.Decode(bytes.NewReader(pngdata))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode PNG")
	}

	i = imaging.Fit(i, Size, Size, imaging.Lanczos)

	var buf bytes.Buffer

	if err := encoder.Encode(&buf, i); err != nil {
		return nil, errors.Wrap(err, "Failed to re-encode PNG")
	}

	return buf.Bytes(), nil
}
