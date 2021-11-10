package extractor

import (
	"github.com/joscha-alisch/go-cgen/internal/extractor/ast"
	"modernc.org/cc/v3"
)

type handler struct {
	ast.BaseHandler

	Definition Definition
}

func (h *handler) OnExternalDeclarationStart(n *cc.ExternalDeclaration) {
	switch n.Case {
	case cc.ExternalDeclarationDecl:
		spec := ast.SpecFrom(n.Declaration)
		ast.EachInitDeclarator(n.Declaration.InitDeclaratorList, func(d *cc.InitDeclarator) {
			if enum, ok := spec.Type.Enum(); ok {
				h.addEnum(enum, d.Declarator)
			}
		})
	}
}

func (h *handler) addEnum(enumSpec ast.SpecTypeEnum, decl *cc.Declarator) {
	enum := Enum{
		Identifier: h.id(decl.DirectDeclarator.Name().String()),
	}

	for _, option := range enumSpec.Options {
		enum.Options = append(enum.Options, EnumOption{
			Identifier: h.id(option.Identifier),
			Value: Value{
				Type:  valueTypeFromAst(option.Value.Type),
				Value: option.Value.Value,
			},
		})
	}

	h.Definition.Enums = append(h.Definition.Enums, enum)
}

func (h *handler) id(identifier string) string {
	if h.Definition.Identifiers == nil {
		h.Definition.Identifiers = make(Identifiers)
	}

	h.Definition.Identifiers[identifier] = identifier
	return identifier
}
