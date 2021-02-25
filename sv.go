package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/pkg/errors"
)

var (
	svgoPath string
	svgoOnce sync.Once
)

func hasSvgo() bool {
	svgoOnce.Do(func() {
		p, _ := exec.LookPath("svgo")
		svgoPath = p

		if svgoPath != "" {
			log.Println("svgo detected, using it to read svg.")
		}
	})

	return svgoPath != ""
}

func svgo(path string) (io.ReadCloser, error) {
	if !hasSvgo() {
		return nil, errors.New("no svgo")
	}

	cmd := exec.Command(svgoPath, "-i", path, "-o", "-")
	cmd.Stderr = os.Stderr

	o, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stdout pipe")
	}

	if err := cmd.Start(); err != nil {
		o.Close()
		return nil, errors.Wrap(err, "failed to start")
	}

	return cmdCloser{o, cmd}, nil
}

type cmdCloser struct {
	io.ReadCloser
	cmd *exec.Cmd
}

func (cc cmdCloser) Close() error {
	closeErr := cc.ReadCloser.Close()

	if err := cc.cmd.Wait(); err != nil {
		return errors.Wrap(err, "svgo cannot finish")
	}

	return closeErr
}
