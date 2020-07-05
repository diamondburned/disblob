package css

import (
	"io"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

var minifier *minify.M

func init() {
	minifier = minify.New()
	minifier.AddFunc("text/css", css.Minify)
}

type MinifyStream struct {
	errch  chan error
	writes chan []byte
}

var _ io.WriteCloser = (*MinifyStream)(nil)

func NewMinifyStream(output io.Writer) *MinifyStream {
	ms := &MinifyStream{
		errch:  make(chan error),
		writes: make(chan []byte),
	}

	// Pipe to transfer between the read and write goroutines.
	r, w := io.Pipe()

	// Start the read goroutine.
	go func() {
		ms.errch <- minifier.Minify("text/css", output, r)
		r.Close()
	}()

	// Start the write goroutine, which writes to minifier.
	go func() {
		for bytes := range ms.writes {
			w.Write(bytes)
		}
		w.Close()
	}()

	return ms
}

// Write must write complete CSS blocks.
func (ms *MinifyStream) Write(b []byte) (int, error) {
	ms.writes <- b
	return len(b), nil
}

// Close closes the writer and waits for the background task to complete.
func (ms *MinifyStream) Close() error {
	close(ms.writes)
	return <-ms.errch
}
