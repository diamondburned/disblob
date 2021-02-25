package main

import (
	"fmt"
	"io/fs"
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
	return strings.TrimSuffix(builder.String(), sep)
}

var BlobmojiPath = filepath.Join(".", "blobmoji")

func blobPrefixName(emoji string) string {
	return "emoji_u" + JoinUnicodeStringify(emoji, "_")
}

func blobPopLastRune(emoji string) (string, bool) {
	parts := strings.Split(emoji, "_")
	if len(parts) < 3 { // never trim the base rune
		return "", false
	}

	return strings.Join(parts[:len(parts)-1], "_"), true
}

// knownFiles maps an emoji prefix to the file.
var (
	knownFiles = map[string]string{}
	knownNames []string
	knownOnce  sync.Once
	knownErr   error
)

func ensureKnownFiles(formats []emojiFormat) error {
	knownOnce.Do(func() {
		dirFS := os.DirFS(BlobmojiPath)

		for _, format := range formats {
			d, err := fs.ReadDir(dirFS, format.Path)
			if err != nil {
				knownErr = errors.Wrap(err, "failed to open formats dir")
				return
			}

			for _, file := range d {
				// Strip the format.
				name := strings.TrimSuffix(file.Name(), format.Ext)

				// Skip if we already have this emoji.
				if _, ok := knownFiles[name]; ok {
					continue
				}

				// Assign to map and slice.
				knownFiles[name] = filepath.Join(BlobmojiPath, format.Path, file.Name())
				knownNames = append(knownNames, name)
			}
		}
	})

	return knownErr
}

var ErrNotFound = errors.New("emoji not found")

// BlobExists returns the filename and true if the given emoji exists in
// Blobmoji. It returns the attempted filename in Path and false otherwise.
func BlobExists(emoji string, formats []emojiFormat) (string, error) {
	if err := ensureKnownFiles(formats); err != nil {
		return "", err
	}

	prefix := blobPrefixName(emoji)

	path, ok := knownFiles[prefix]
	if ok {
		return path, nil
	}

	var shorten = prefix
	for {
		shorten, ok = blobPopLastRune(shorten)
		if !ok {
			break
		}

		path, ok := knownFiles[shorten]
		if ok {
			return path, nil
		}
	}

	// Try the slow path. Search the prefix and grab the first match.
	for _, name := range knownNames {
		if strings.HasPrefix(name, prefix) {
			return knownFiles[name], nil
		}
	}

	return prefix, errors.Wrapf(ErrNotFound, "error with %q", prefix)
}
