package inline

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestWritePipeline(t *testing.T) {
	dst := bytes.Buffer{}

	w := WritePipeline(
		&dst,
		pipeCopy(0, false),
		pipeCopy(1, false),
		pipeCopy(2, false),
		pipeCopy(3, false),
		pipeCopy(4, false),
	)

	_, err := io.WriteString(w, "hello, world")
	if err != nil {
		t.Fatal("failed to write to pipeline:", err)
	}

	if err := w.Close(); err != nil {
		t.Fatal("failed to close pipeline:", err)
	}

	if dst.String() != "hello, world 0 1 2 3 4" {
		t.Fatalf("unexpected string in buffer: %q", dst.String())
	}
}

func TestWritePipelineErroneous(t *testing.T) {
	dst := bytes.Buffer{}

	w := WritePipeline(
		&dst,
		pipeCopy(0, false),
		pipeCopy(1, false),
		pipeCopy(2, true),
		pipeCopy(3, false),
	)

	_, err := io.WriteString(w, "hello, world")
	if err != nil {
		t.Fatal("failed to write to pipeline:", err)
	}

	if err := w.Close(); err != nil {
		if !errors.Is(err, fakeErr) {
			t.Fatal("unexpected error while closing pipeline:", err)
		}
		return
	}

	t.Fatal("expected close error, got none")
}

var fakeErr = errors.New("fake err")

func pipeCopy(i int, mustErr bool) Pipeline {
	return func(w io.Writer, r io.Reader) error {
		_, err := io.Copy(w, r)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(w, " %d", i)
		if err != nil {
			return err
		}

		if mustErr {
			return fakeErr
		}

		return nil
	}
}
