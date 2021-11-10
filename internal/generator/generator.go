package generator

type Generator interface{}

func New() Generator {
	return &generator{}
}

type generator struct{}
