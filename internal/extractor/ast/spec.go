package ast

import (
	"modernc.org/cc/v3"
)

type Spec struct {
	Type     SpecType
	TypeDef  bool
	Static   bool
	Extern   bool
	Auto     bool
	Register bool
	Local    bool
	Const    bool
	Restrict bool
	Volatile bool
	Atomic   bool
	Inline   bool
	NoReturn bool
}

type SpecType struct {
	enum     bool
	enumSpec SpecTypeEnum
}

func (s SpecType) Enum() (SpecTypeEnum, bool) {
	return s.enumSpec, s.enum
}

type SpecTypeEnum struct {
	Options []SpecTypeEnumOption
}

type SpecTypeEnumOption struct {
	Identifier string
	Value      Value
}

type Value struct {
	Type  ValueType
	Value interface{}
}

type ValueType int

const (
	ValueTypeUnsupported ValueType = iota
	ValueTypeIdentifier
	ValueTypeConstant
)

func SpecFrom(d *cc.Declaration) Spec {
	spec := Spec{}

	EachDeclarationSpecifier(d.DeclarationSpecifiers, func(item *cc.DeclarationSpecifiers) {
		switch item.Case {
		case cc.DeclarationSpecifiersStorage:
			switch item.StorageClassSpecifier.Case {
			case cc.StorageClassSpecifierTypedef:
				spec.TypeDef = true
			case cc.StorageClassSpecifierExtern:
				spec.Extern = true
			case cc.StorageClassSpecifierStatic:
				spec.Static = true
			case cc.StorageClassSpecifierAuto:
				spec.Auto = true
			case cc.StorageClassSpecifierRegister:
				spec.Register = true
			case cc.StorageClassSpecifierThreadLocal:
				spec.Local = true
			}
		case cc.DeclarationSpecifiersTypeSpec:
			spec.Type = specTypeFrom(item.TypeSpecifier)
		case cc.DeclarationSpecifiersTypeQual:
			switch item.TypeQualifier.Case {
			case cc.TypeQualifierConst:
				spec.Const = true
			case cc.TypeQualifierRestrict:
				spec.Restrict = true
			case cc.TypeQualifierVolatile:
				spec.Volatile = true
			case cc.TypeQualifierAtomic:
				spec.Atomic = true
			}
		case cc.DeclarationSpecifiersFunc:
			switch item.FunctionSpecifier.Case {
			case cc.FunctionSpecifierInline:
				spec.Inline = true
			case cc.FunctionSpecifierNoreturn:
				spec.NoReturn = true
			}
		case cc.DeclarationSpecifiersAlignSpec:
		case cc.DeclarationSpecifiersAttribute:
		}
	})

	return spec
}

func specTypeFrom(t *cc.TypeSpecifier) SpecType {
	spec := SpecType{}

	switch t.Case {
	case cc.TypeSpecifierEnum:
		spec.enum = true
		spec.enumSpec = specTypeEnumFrom(t.EnumSpecifier)
	}

	return spec
}

func specTypeEnumFrom(t *cc.EnumSpecifier) SpecTypeEnum {
	spec := SpecTypeEnum{}

	var previous *SpecTypeEnumOption
	EachEnumerator(t.EnumeratorList, func(e *cc.Enumerator) {
		option := specTypeEnumOptionFrom(e, previous)
		spec.Options = append(spec.Options, option)
		previous = &option
	})
	return spec
}

func specTypeEnumOptionFrom(e *cc.Enumerator, prev *SpecTypeEnumOption) SpecTypeEnumOption {
	o := SpecTypeEnumOption{
		Identifier: e.Token.String(),
	}

	switch e.Case {
	case cc.EnumeratorIdent:
		if prev == nil {
			o.Value = Value{
				Type:  ValueTypeConstant,
				Value: 0,
			}
		} else if prev.Value.Type == ValueTypeConstant {
			o.Value = Value{
				Type:  ValueTypeConstant,
				Value: increment(prev.Value.Value),
			}
		}
	case cc.EnumeratorExpr:
		o.Value = FindValueConstantExpression(e.ConstantExpression)
	}
	return o
}

func increment(number interface{}) interface{} {
	if number == nil {
		return 0
	}
	if n, ok := number.(int); ok {
		return n + 1
	}

	if n, ok := number.(int32); ok {
		return n + 1
	}

	if n, ok := number.(int64); ok {
		return n + 1
	}

	if n, ok := number.(uint); ok {
		return n + 1
	}
	if n, ok := number.(uint32); ok {
		return n + 1
	}
	if n, ok := number.(uint64); ok {
		return n + 1
	}
	return 0
}
