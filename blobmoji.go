package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode"

	"github.com/pkg/errors"
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

	// Trim trailing separators.
	return strings.TrimRight(builder.String(), sep)
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

// knownFiles maps an emoji prefix to the file.
var knownFiles = map[string]string{}
var knownNames = []string{}
var knownOnce = sync.Once{}

func ensureKnownFiles() (toperr error) {
	knownOnce.Do(func() {
		for _, formats := range tryformats {
			var dirname = filepath.Join(BlobmojiPath, formats.path)

			d, err := os.Open(dirname)
			if err != nil {
				toperr = errors.Wrap(err, "Failed to open formats dir")
				return
			}

			f, err := d.Readdir(-1)
			if err != nil {
				toperr = errors.Wrap(err, "Failed to read dir")
				return
			}

			for _, file := range f {
				// Strip the format.
				name := strings.Split(file.Name(), ".")[0]

				// Skip if we already have this emoji.
				if _, ok := knownFiles[name]; ok {
					continue
				}

				// Assign to map and slice.
				knownFiles[name] = filepath.Join(dirname, file.Name())
				knownNames = append(knownNames, name)
			}
		}
	})
	return
}

type Blob struct {
	Path string // path to the blog
	SVG  bool
}

var stopWalk = errors.New("stop walk")
var ErrNotFound = errors.New("emoji not found")

// BlobExists returns the filename and true if the given emoji exists in
// Blobmoji. It returns the attempted filename in Path and false otherwise.
func BlobExists(emoji string) (string, error) {
	if err := ensureKnownFiles(); err != nil {
		return "", err
	}

	var prefix = blobPrefixName(emoji)

	if path, ok := knownFiles[prefix]; ok {
		return path, nil
	}

	// Try the slow path. Search the prefix and grab the first match.
	for _, name := range knownNames {
		if strings.HasPrefix(name, prefix) {
			return knownFiles[name], nil
		}
	}

	return prefix, ErrNotFound
}
