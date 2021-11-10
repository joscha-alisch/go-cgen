package extractor

import "modernc.org/cc/v3"

type Extractor interface {
	Extract(ast *cc.AST) (*Definition, error)
}

func New() Extractor {
	return &extractor{}
}

type extractor struct {

}

func (e *extractor) Extract(ast *cc.AST) (*Definition, error) {
	return nil, nil
}

