package parser

import "modernc.org/cc/v3"

type Config struct {

}

type Parser interface {
	Parse() (*cc.AST, error)
}

func New() Parser {
	return &parser{}
}

type parser struct {

}

func (p *parser) Parse() (*cc.AST, error) {
	panic("implement me")
}
