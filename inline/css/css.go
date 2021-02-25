package css

import (
	"io"

	"github.com/diamondburned/disblob/inline"
)

func Inline(w io.Writer, css io.Reader) error {
	return inline.Minifier.Minify("text/css", w, css)
}
