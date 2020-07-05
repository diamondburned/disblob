package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Bios-Marcel/discordemojimap"
	"github.com/diamondburned/disblob/minify/png"
	"github.com/diamondburned/disblob/minify/svg"
	"github.com/diamondburned/disblob/process"
)

const f_CSS = `:root{--%[1]s:url('%[2]s')}
img[alt=":%[1]s:"]{content:var(--%[1]s)}
li[aria-label^=":%[1]s:"]>div{background-image:var(--%[1]s)!important;background-position:0!important;background-size:contain!important}
`

func main() {
	p, err := process.NewProcessor("./state.db")
	if err != nil {
		log.Fatalln("Failed to make a new state")
	}

	for name := range discordemojimap.EmojiMap {
		p.DoAsync(name, func(name string) bool {
			var emoji = discordemojimap.EmojiMap[name]
			p, err := BlobExists(emoji)
			if err != nil {
				if err != ErrNotFound {
					log.Fatalln("Failed to cache emojis:", err)
				}
				logfail("Skipping non-existent", name, p, nil)
				return false
			}

			f, err := ioutil.ReadFile(p)
			if err != nil {
				logfail("Failed to open", name, p, err)
				return false
			}

			var bytes []byte
			if strings.HasSuffix(p, ".svg") {
				bytes, err = svg.Inline(f)
			} else {
				bytes, err = png.Inline(f)
			}

			if err != nil {
				logfail("Failed to inline", name, p, err)
				return false
			}

			fmt.Printf(f_CSS, name, string(bytes))
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
