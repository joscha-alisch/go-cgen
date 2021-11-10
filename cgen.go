package cgen

import (
	"os"

	"github.com/joscha-alisch/go-cgen/internal/extractor"
	"github.com/joscha-alisch/go-cgen/internal/fswriter"
	"github.com/joscha-alisch/go-cgen/internal/generator"
	"github.com/joscha-alisch/go-cgen/internal/organiser"
	"github.com/joscha-alisch/go-cgen/internal/parser"
	"github.com/joscha-alisch/go-cgen/internal/ruleset"
	"github.com/joscha-alisch/go-cgen/internal/translator"
)

type Config struct {
	Headers []string
	OutDir  string
}

type CGen interface {
	Run() error
}

func New() CGen {
	return &cGen{}
}

type cGen struct {
	config Config
}

func (c *cGen) Run() error {
	p := parser.New(parser.Config{})
	e := extractor.New()
	r := ruleset.New()
	o := organiser.New()
	t := translator.New()
	g := generator.New()
	w := fswriter.New()

	ast, err := p.Parse(c.config.Headers...)
	if err != nil {
		return err
	}

	def, err := e.Extract(ast)
	if err != nil {
		return err
	}

	def, err = r.Apply(def)
	if err != nil {
		return err
	}

	lib, err := o.Organise(def)
	if err != nil {
		return err
	}

	gen, err := t.Translate(lib, g)
	if err != nil {
		return err
	}

	err = w.Write(gen, os.DirFS(c.config.OutDir))
	if err != nil {
		return err
	}

	return nil
}
