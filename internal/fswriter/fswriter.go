package fswriter

import (
	"io/fs"
)

type Input interface {

}

type FSWriter interface {
	Write(input Input, fs fs.FS) error

}

func New() FSWriter {
	return &fsWriter{}
}

type fsWriter struct {

}

func (f *fsWriter) Write(input Input, fs fs.FS) error {
	return nil
}


