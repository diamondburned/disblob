package svg

import (
	"bytes"
	"net/url"
	"regexp"

	"github.com/pkg/errors"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/svg"
)

const inlinePrefix = "data:image/svg+xml,"

var minifier *minify.M

func init() {
	minifier = minify.New()
	minifier.AddFunc("image/svg+xml", svg.Minify)
	minifier.AddFunc("text/css", css.Minify)
}

// Inline transforms an SVG into its inline data equivalent.
func Inline(svg []byte) ([]byte, error) {
	// Minify before validating and inlining. Allocate a stage 1 buffer.
	stage1 := bytes.Buffer{}
	stage1.Grow(len(svg))

	if err := minifier.Minify("image/svg+xml", &stage1, bytes.NewReader(svg)); err != nil {
		return nil, errors.Wrap(err, "Failed to minify")
	}

	// This breaks some SVGs, so we're not doing it.

	// // Allocate a stage 2 buffer to write the prefix as well as minify the SVG
	// // tag's attribute.
	// stage2 := bytes.Buffer{}
	// stage2.Grow(stage1.Len())

	// if err := minifyOpeningTag(&stage2, &stage1); err != nil {
	// 	return nil, errors.Wrap(err, "Failed to minify the SVG attribute tag")
	// }

	// Finally, URL escape everything.
	stage3 := bytes.Buffer{}
	stage3.WriteString(inlinePrefix)
	stage3.Write(escape(stage1.Bytes()))

	return stage3.Bytes(), nil
}

var (
	// regex to escape symbols
	symbolRegex = regexp.MustCompile(`(?m)[\r\n#?\[\\\]^\x60{|}'"]`)
	// regex to strip invalid style attributes
	invstyleRegex = regexp.MustCompile(`(?mU) style=".*'.*'.*"`)
	// regex to fix ugly font
	comicNeueRegex = regexp.MustCompile(`(?mUi)(['"])Comic Neue( \w+)?['"]`)

	sansserif = []byte("${1}sans-serif${1}")
)

func escape(src []byte) []byte {
	// replace all Comit Neues w/ normal Sans
	src = comicNeueRegex.ReplaceAll(src, sansserif)

	return []byte(url.PathEscape(string(src)))
}