package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func JoinUnicodeStringify(ustr, sep string) string {
	var runes = []rune(ustr)

	builder := strings.Builder{}
	// Pre-grow to a worst case FFFFF (length 5) plus length of separator, all
	// multiplied by the length of total runes.
	builder.Grow((5 + len(sep)) * len(runes))

	for i, r := range runes {
		// Ignore all variant selectors.
		if unicode.Is(unicode.Variation_Selector, r) {
			continue
		}

		// Print the hex value at least 4 numbers wide.
		fmt.Fprintf(&builder, "%04x", r)

		if i < len(runes)-1 {
			builder.WriteString(sep)
		}
	}

	return builder.String()
}

var BlobmojiPath = filepath.Join(".", "blobmoji")

func blobPrefixName(emoji string) string {
	return "emoji_u" + JoinUnicodeStringify(emoji, "_")
}

// Try SVG first, then PNG.
var tryformats = []struct{ path, format string }{
	{"svg", ".svg"},
	{filepath.Join("png", "128"), ".png"},
}

type Blob struct {
	Path string // path to the blog
	SVG  bool
}

// BlobExists returns the filename and true if the given emoji exists in
// Blobmoji. It returns an empty string and false otherwise.
func BlobExists(emoji string) (Blob, bool) {
	var prefix = blobPrefixName(emoji)
	var name string

	for _, format := range tryformats {
		name = prefix + format.format
		path := filepath.Join(BlobmojiPath, format.path, name)

		if _, err := os.Stat(path); err == nil {
			return Blob{path, format.format == ".svg"}, true
		}
	}

	return Blob{}, false
}
