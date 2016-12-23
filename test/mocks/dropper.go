package mocks

import "io"

type Dropper struct {
	Recieved struct {
		Filename string
		File     io.Reader
	}
	Returns struct {
	}
}

func (d *Dropper) Drop(filename string, file io.Reader) {
	d.Recieved.Filename = filename
	d.Recieved.File = file
}
