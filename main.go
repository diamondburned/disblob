package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/Bios-Marcel/discordemojimap"
	"github.com/cheggaaa/pb/v3"
	"github.com/diamondburned/disblob/inline"
	"github.com/diamondburned/disblob/inline/css"
	"github.com/diamondburned/disblob/inline/png"
	"github.com/diamondburned/disblob/inline/svg"
)

// Spawn 2 times more workers than thread count.
var jobs = runtime.GOMAXPROCS(-1) * 2

// Broken selectors:
//
//    `li[aria-label^=":%[1]s:"]>div` + "{"

const f_CSS = "" +
	`img[src^="/assets"][data-name=":%[1]s:"], ` + // used for emoji names
	`img[src^="/assets"][alt="%[2]s"] ` + // used for unicodes

	"{ " +

	// Write the Firefox CSS fix where both versions of an emoji would appear
	// (gk).
	`object-position: var(--op); ` + // only for available emojis
	`background-image: url("%[3]s"); ` +

	"} "

type emojiFormat struct {
	Path string
	Ext  string
}

var (
	// Known formats.
	svgFormat = emojiFormat{"svg", ".svg"}
	pngFormat = emojiFormat{"png/128", ".png"}

	// Flags.
	preferPNG = false
	noSvgo    = false
	verbose   = false
)

func main() {
	flag.IntVar(&png.Size, "size", png.Size, "PNG image size")
	flag.BoolVar(&preferPNG, "png", preferPNG, "prefer PNG over SVG")
	flag.BoolVar(&noSvgo, "nosvgo", noSvgo, "force no svgo for SVG")
	flag.BoolVar(&verbose, "v", verbose, "show unfetched emojis")
	flag.IntVar(&jobs, "j", jobs, "nubmer of jobs, default is thread count * 2")

	flag.Parse()

	var formats = []emojiFormat{svgFormat, pngFormat}
	if preferPNG {
		formats = []emojiFormat{pngFormat, svgFormat}
	}

	if !noSvgo && !hasSvgo() {
		log.Fatalln("missing svgo (use -nosvgo or -png or install svgo)")
	}

	bufOut := bufio.NewWriterSize(os.Stdout, 40960) // 40KB
	w := inline.WritePipeline(bufOut, css.Inline)

	fmt.Fprintln(w, `
		div[class^="reactionInner"] img.emoji,
		span[class^="emojiContainer"] img.emoji {
			/* value for object-position */
			--op: -9999px -9999px;

			/* background fix */
			background-position: 0 !important;
			background-size: contain !important;
			background-repeat: no-repeat;
		}
	`)

	type emojiFile struct {
		Name  string
		Emoji string
	}

	emojiCh := make(chan emojiFile, jobs)

	var wg sync.WaitGroup
	var failed failTripper

	progress := pb.New(len(discordemojimap.EmojiMap))
	progress.SetWriter(os.Stderr)

	if !verbose {
		progress.Start()
	}

	for i := 0; i < jobs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Dedicate each goroutine its own bytes buffers.
			var reader io.ReadCloser
			readBuf := bufio.NewReaderSize(reader, 20480) // reader buffer, 20KB

			outBuf := bytes.Buffer{}
			outBuf.Grow(10240) // pre-allocate 10KB

			doEmoji := func(emoji emojiFile) {
				p, err := BlobExists(emoji.Emoji, formats)
				if err != nil {
					if err != ErrNotFound {
						failed.log("failed to open", emoji.Name, "", err)
						return
					}

					logfail("skipping non-existent", emoji.Name, "", err)
					return
				}

				fileExt := filepath.Ext(p)
				if fileExt == ".svg" && !noSvgo && hasSvgo() {
					reader, err = svgo(p)
				} else {
					reader, err = os.Open(p)
				}

				if err != nil {
					failed.log("failed to open", emoji.Name, p, err)
					return
				}

				// Close reader when done.
				defer reader.Close()

				// Reset the buffers before using.
				readBuf.Reset(reader)
				outBuf.Reset()

				switch fileExt {
				case ".png":
					err = png.Inline(&outBuf, readBuf)
				case ".svg":
					err = svg.Inline(&outBuf, readBuf)
				default:
					// Skip unknown formats.
					failed.log("unknown format of", emoji.Name, p, nil)
					return
				}

				if err != nil {
					failed.log("failed to inline", emoji.Name, p, err)
					return
				}

				fmt.Fprintf(w, f_CSS, emoji.Name, emoji.Emoji, outBuf.Bytes())
			}

			for emoji := range emojiCh {
				doEmoji(emoji)
				progress.Increment()
			}
		}()
	}

	for name, emoji := range discordemojimap.EmojiMap {
		emojiCh <- emojiFile{
			Name:  name,
			Emoji: emoji,
		}
	}

	// Mark that jobs are all distributed and wait for them to finish.
	close(emojiCh)
	wg.Wait()

	progress.Finish()

	if err := w.Close(); err != nil {
		log.Fatalln("pipeline error:", err)
	}

	if err := bufOut.Flush(); err != nil {
		log.Fatalln("flush error:", err)
	}

	if failed.isFailed() {
		os.Exit(1)
	}
}

type failTripper uint32

func (f *failTripper) fail() { atomic.StoreUint32((*uint32)(f), 1) }

func (f *failTripper) isFailed() bool { return atomic.LoadUint32((*uint32)(f)) > 0 }

func (f *failTripper) log(desc, name, path string, err error) {
	f.fail()
	logfail(desc, name, path, err)
}

func logfail(desc, name, path string, err error) {
	if verbose {
		log.Println(errfail(desc, name, path, err))
	}
}

func errfail(desc, name, path string, err error) error {
	var errstring string
	if path != "" {
		errstring += fmt.Sprintf(" (path: %s)", path)
	}
	if err != nil {
		errstring = ": " + err.Error()
	}

	return fmt.Errorf("%s emoji %s%s", desc, name, errstring)
}
