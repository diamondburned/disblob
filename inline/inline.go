package inline

import (
	"io"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/svg"
	"golang.org/x/sync/errgroup"
)

var Minifier *minify.M

func init() {
	Minifier = minify.New()
	Minifier.AddFunc("image/svg+xml", svg.Minify)
	Minifier.AddFunc("text/css", css.Minify)
}

// Pipeline is a pipeline callback to start a writer that processes a reader. It
// is called in its own goroutine.
type Pipeline func(io.Writer, io.Reader) error

// WritePipeline creates a new write pipeline. It returns nil if no pipelines
// are given. The caller must close the WriteCloser to flush and wait for the
// pipeline to finish.
func WritePipeline(dst io.Writer, pipelines ...Pipeline) io.WriteCloser {
	if len(pipelines) == 0 {
		return nil
	}

	var wg errgroup.Group

	topReader, topWriter := io.Pipe()

	for i, pipe := range pipelines {
		// Read from the copy, not from range.
		pipe := pipe

		if i == len(pipelines)-1 {
			wg.Go(func() error {
				err := pipe(dst, topReader)
				topReader.CloseWithError(err)
				return err
			})

			break
		}

		newSrc, newDst := io.Pipe()
		passDownReader := topReader

		wg.Go(func() error {
			err := pipe(newDst, passDownReader)
			newDst.CloseWithError(err)
			passDownReader.Close()
			return err
		})

		topReader = newSrc
	}

	return wgClose{topWriter, &wg}
}

type wgClose struct {
	*io.PipeWriter
	wg *errgroup.Group
}

func (wg wgClose) Close() error {
	wg.PipeWriter.Close()

	// Always wait for goroutines to finish.
	return wg.wg.Wait()
}
