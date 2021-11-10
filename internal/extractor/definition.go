package extractor

import "github.com/joscha-alisch/go-cgen/internal/extractor/ast"

type Definition struct {
	Identifiers Identifiers
	Enums       []Enum
}

type Identifiers map[string]string

type Enum struct {
	Identifier string
	Options    []EnumOption
}

type EnumOption struct {
	Identifier string
	Value      Value
}

type Value struct {
	Type  ValueType
	Value interface{}
}

type ValueType int

const (
	ValueTypeUnknown ValueType = iota
	ValueTypeIdentifier
	ValueTypeLiteral
)

func valueTypeFromAst(t ast.ValueType) ValueType {
	switch t {
	case ast.ValueTypeIdentifier:
		return ValueTypeIdentifier
	case ast.ValueTypeConstant:
		return ValueTypeLiteral
	default:
		return ValueTypeUnknown
	}
}
