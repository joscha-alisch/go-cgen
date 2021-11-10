package parser

import (
	"modernc.org/cc/v3"
	"strings"
)

// Parser parses C-headers and returns the corresponding abstract syntax tree.
type Parser interface {
	Parse(headers ...string) (*cc.AST, error)
}

/**
New returns a new Parser for the given configuration.
*/
func New(config Config) Parser {
	return &parser{
		config: config,
	}
}

type parser struct {
	config Config
}

const (
	predefinedSourceName = "predefined"
	definesSourceName    = "defines"
)

var parse = cc.Parse
var hostConfig = cc.HostConfig

func (p *parser) Parse(headers ...string) (*cc.AST, error) {
	cfg := p.config.OverrideConfig
	if cfg == nil {
		cfg = &cc.Config{}
	}

	var predefines string
	var includes, sysIncludes []string
	var err error
	if !p.config.SkipHostConfig {
		predefines, includes, sysIncludes, err = hostConfig(p.config.CppCommand)
		if err != nil {
			return nil, err
		}
	}

	includes = append(includes, p.config.IncludePaths...)
	sysIncludes = append(sysIncludes)

	var sources []cc.Source

	if predefines != "" {
		sources = append(sources, cc.Source{Name: predefinedSourceName, Value: predefines})
	}

	if p.config.Defines != nil {
		sources = append(sources, cc.Source{Name: definesSourceName, Value: strings.Join(p.config.Defines, "\n")})
	}

	for i := range headers {
		sources = append(sources, cc.Source{Name: headers[i]})
	}

	return parse(cfg, includes, sysIncludes, sources)
}
