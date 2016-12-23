package localdrop

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalDrop struct {
	Out     io.Writer
	BaseDir string
}

func (l LocalDrop) Drop(filename string, file io.Reader) {
	err := l.dropErr(filename, file)
	if err != nil {
		fmt.Fprintf(l.Out, "download failed: %s: %s\n", filename, err)
		return
	}
}

func (l LocalDrop) dropErr(filename string, file io.Reader) error {
	localFilename := filepath.Join(l.BaseDir, filename)

	_, err := os.Stat(localFilename)
	if !os.IsNotExist(err) {
		return errors.New("file already exists")
	}

	localFile, err := os.Create(localFilename)
	if err != nil {
		return err
	}

	fmt.Fprintf(l.Out, "downloading: %s\n", filename)
	_, err = io.Copy(localFile, file)
	if err != nil {
		return err
	}

	fmt.Fprintf(l.Out, "download complete: %s\n", filename)
	return nil
}
