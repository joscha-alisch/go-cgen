package ruleset

import (
	"github.com/joscha-alisch/go-cgen/internal/extractor"
)

type RuleSet interface {
	Apply(def *extractor.Definition) (*extractor.Definition, error)
}

func New() RuleSet {
	return &ruleSet{}
}

type ruleSet struct {

}

func (r *ruleSet) Apply(def *extractor.Definition) (*extractor.Definition, error) {
	return nil, nil
}
