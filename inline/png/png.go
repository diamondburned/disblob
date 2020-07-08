package png

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
)

// Size to reduce to.
const Size = 48

const inlinePrefix = "data:image/png;base64,"

// encoder with padding, REQUIRED
var b64encoder = base64.StdEncoding.WithPadding(base64.StdPadding)

func Inline(pngdata []byte) ([]byte, error) {
	// Compress first.
	b, err := compress(pngdata)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to compress PNG")
	}

	var b64len = base64.StdEncoding.EncodedLen(len(b))
	var b64dat = make([]byte, len(inlinePrefix)+b64len)

	// Write the header.
	copy(b64dat, []byte(inlinePrefix))

	// Base64 encode.
	b64encoder.Encode(b64dat[len(inlinePrefix):], b)

	return b64dat, nil
}

var pngencoder = png.Encoder{
	CompressionLevel: png.BestCompression,
}

func compress(pngdata []byte) ([]byte, error) {
	i, err := png.Decode(bytes.NewReader(pngdata))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode PNG")
	}

	i = imaging.Fit(i, Size, Size, imaging.Lanczos)

	var buf bytes.Buffer

	if err := pngencoder.Encode(&buf, i); err != nil {
		return nil, errors.Wrap(err, "Failed to re-encode PNG")
	}

	return buf.Bytes(), nil
}
