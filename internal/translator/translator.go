package translator

type Generator interface {

}

type Provider interface {

}

type Result struct {

}

type Translator interface {
	Translate(lib Provider, g Generator) (*Result, error)
}

func New() Translator {
	return &translator{}
}

type translator struct {

}

func (t *translator) Translate(lib Provider, g Generator) (*Result, error) {
	return nil, nil
}
