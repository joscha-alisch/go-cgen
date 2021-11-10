package organiser

import (
	"github.com/joscha-alisch/go-cgen/internal/extractor"
)

type Organiser interface {
	Organise(def extractor.Definition) (*Library, error)
}

func New() Organiser {
	return &organiser{}
}

type organiser struct{}

func (o *organiser) Organise(extractor.Definition) (*Library, error) {
	return nil, nil
}
