package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func Drop(req *http.Request) error {
	formFile, formFileHeaders, err := req.FormFile("file")
	if err != nil {
		return err
	}
	defer formFile.Close()

	filename := path.Base(formFileHeaders.Filename)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot create file: %s", err)
	}

	fmt.Printf("downloading: %s\n", filename)
	_, err = io.Copy(file, formFile)
	if err != nil {
		return fmt.Errorf("cannot download file %s", err)
	}
	fmt.Printf("download complete: %s\n", filename)

	return nil
}
