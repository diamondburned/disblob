package svg

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/url"
	"regexp"

	"github.com/diamondburned/disblob/inline"
	"github.com/pkg/errors"
)

const inlinePrefix = "data:image/svg+xml,"

// PreferBase64 prefers base64 over URL encoding.
var PreferBase64 = true

// ForceSansSerif forces the font to be sans-serif instead of Comic Neue.
var ForceSansSerif = true

// regex to fix ugly font
var comicNeueRegex = regexp.MustCompile(`(?mUi)(['"])Comic Neue( \w+)?['"]`)

var (
	urlEncodePrefix = "data:image/svg+xml,"
	b64EncodePrefix = "data:image/svg+xml;base64,"
)

// Inline transforms an SVG into its inline data equivalent.
func Inline(w io.Writer, svg io.Reader) (err error) {
	if PreferBase64 {
		_, err = w.Write([]byte(b64EncodePrefix))
	} else {
		_, err = w.Write([]byte(urlEncodePrefix))
	}

	if err != nil {
		return errors.Wrap(err, "failed to write prefix")
	}

	if !ForceSansSerif {
		return doInline(svg, w)
	}

	// If we want to use a sans-serif font, then we must allocate a buffer.
	buf := bytes.Buffer{}
	buf.Grow(4096)

	if err := doInline(svg, &buf); err != nil {
		return err
	}

	b := comicNeueRegex.ReplaceAll(buf.Bytes(), []byte("${1}sans-serif${1}"))

	_, err = w.Write(b)
	return err
}

func doInline(svg io.Reader, w io.Writer) (err error) {
	// Either encode in Base64 or URL.
	if PreferBase64 {
		enc := base64.NewEncoder(base64.StdEncoding, w)
		defer enc.Close()

		w = enc

	} else {
		w = urlEncoder{w}
	}

	if err := inline.Minifier.Minify("image/svg+xml", w, svg); err != nil {
		return errors.Wrap(err, "failed to minify")
	}

	return nil
}

type urlEncoder struct{ w io.Writer }

func (enc urlEncoder) Write(b []byte) (int, error) {
	_, err := enc.w.Write([]byte(url.QueryEscape(string(b))))
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
