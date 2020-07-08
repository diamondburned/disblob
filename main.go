package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Bios-Marcel/discordemojimap"
	"github.com/diamondburned/disblob/inline/png"
	"github.com/diamondburned/disblob/inline/svg"
	"github.com/diamondburned/disblob/process"
	"github.com/pkg/errors"
)

const f_CSS = "" +
	`img[alt=":%[1]s:"],` +
	`li[aria-label^=":%[1]s:"]>div` + "{" +

	"content:var(--a);" +
	"background-image:var(--a)!important;" +
	`--a:url("%[2]s")` +

	"}\n"

func main() {
	// Write the Firefox CSS fix where both versions of an emoji would appear
	// (gk).
	fmt.Println(`.emoji[src$=".svg"] { object-position: -9999px -9999px }`)

	// Write the global list rule that fixes emoji scaling.
	fmt.Println(`img[class^="emoji"], li[id^="emoji"] > div {
		background-position: 0 !important; background-size: contain !important;
	}`)

	// Make a new processor and use it for root variables.
	var para = process.NewParallel()

	for name, emoji := range discordemojimap.EmojiMap {
		para.DoAsync(name, emoji, func(name, emoji string) error {
			p, err := BlobExists(emoji)
			if err != nil {
				if err != ErrNotFound {
					return errors.Wrap(err, "Failed to check exist")
				}
				logfail("Skipping non-existent", name, p, nil)
				return nil
			}

			// Read the emoji data. Read errors can be fatal.
			d, err := ioutil.ReadFile(p)
			if err != nil {
				return errfail("Failed to open", name, p, err)
			}

			// Try and inline it.
			var bytes []byte
			switch {
			case strings.HasSuffix(p, ".png"):
				bytes, err = png.Inline(d)
			case strings.HasSuffix(p, ".svg"):
				bytes, err = svg.Inline(d)
			default:
				// Skip unknown formats.
				logfail("Unknown format of", name, p, nil)
				return nil
			}

			if err != nil {
				logfail("Failed to inline", name, p, err)
				return nil
			}

			fmt.Printf(f_CSS, name, string(bytes))
			return nil
		})
	}

	// Wait for all root variables to finish writing.
	para.Wait()

	// Error check.
	if err := para.Err(); err != nil {
		log.Fatalln("Failed to write emoji data:", err)
	}
}

func logfail(desc, name, path string, err error) {
	log.Println(errfail(desc, name, path, err))
}

func errfail(desc, name, path string, err error) error {
	var errstring string
	if err != nil {
		errstring = ": " + err.Error()
	}

	return fmt.Errorf("%s emoji %s (path: %s)%s", desc, name, path, errstring)
}
