package extractor

import (
	"github.com/joscha-alisch/go-cgen/internal/extractor/ast"
	"modernc.org/cc/v3"
)

/**
An Extractor extracts everything from the given C headers that is relevant to generating bindings.

This currently includes:
- A list of all Identifiers
- Enums
*/
type Extractor interface {
	Extract(ast *cc.AST) (Definition, error)
}

func New() Extractor {
	return &extractor{}
}

type extractor struct {
}

func (e *extractor) Extract(t *cc.AST) (Definition, error) {
	h := handler{}

	v := ast.NewVisitor(&h)
	v.Visit(t.TranslationUnit)

	return h.Definition, nil
}
