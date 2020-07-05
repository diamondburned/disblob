package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Bios-Marcel/discordemojimap"
	"github.com/diamondburned/disblob/minify/png"
	"github.com/diamondburned/disblob/minify/svg"
	"github.com/diamondburned/disblob/process"
)

const f_CSS = `img[alt=":%s:"]{content:url('%s');}
`

func main() {
	p, err := process.NewProcessor("./state.db")
	if err != nil {
		log.Fatalln("Failed to make a new state")
	}

	for name := range discordemojimap.EmojiMap {
		p.DoAsync(name, func(name string) bool {
			var emoji = discordemojimap.EmojiMap[name]
			bl, ok := BlobExists(emoji)
			if !ok {
				logfail("Skipping non-existent", name, bl.Path, nil)
				return false
			}

			f, err := ioutil.ReadFile(bl.Path)
			if err != nil {
				logfail("Failed to open", name, bl.Path, err)
				return false
			}

			var b []byte
			if bl.SVG {
				b, err = svg.Inline(f)
			} else {
				b, err = png.Inline(f)
			}

			if err != nil {
				logfail("Failed to inline", name, bl.Path, err)
				return false
			}

			fmt.Printf(f_CSS, name, string(b))
			return true
		})
	}

	if err := p.Finalize(); err != nil {
		log.Fatalln("Failed to finalize:", err)
	}
}

func logfail(desc, name, path string, err error) {
	var errstring string
	if err != nil {
		errstring = ": " + err.Error()
	}

	log.Printf("%s emoji %s (path: %s)%s", desc, name, path, errstring)
}
