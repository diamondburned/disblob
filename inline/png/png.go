package png

import (
	"encoding/base64"
	"image/png"
	"io"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
)

// Size to reduce to.
var Size = 64

const inlinePrefix = "data:image/png;base64,"

// encoder with padding, REQUIRED
var b64encoder = base64.StdEncoding.WithPadding(base64.StdPadding)

func Inline(w io.Writer, png io.Reader) error {
	_, err := w.Write([]byte(inlinePrefix))
	if err != nil {
		return errors.Wrap(err, "failed to write inline prefix")
	}

	enc := base64.NewEncoder(base64.StdEncoding, w)

	if err := compress(png, enc); err != nil {
		return err
	}

	return enc.Close()
}

var pngencoder = png.Encoder{
	CompressionLevel: png.BestCompression,
}

func compress(r io.Reader, w io.Writer) error {
	i, err := png.Decode(r)
	if err != nil {
		return errors.Wrap(err, "failed to decode PNG")
	}

	i = imaging.Fit(i, Size, Size, imaging.CatmullRom)

	if err := pngencoder.Encode(w, i); err != nil {
		return errors.Wrap(err, "failed to re-encode PNG")
	}

	return nil
}
