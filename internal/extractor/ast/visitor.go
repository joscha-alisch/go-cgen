package ast

import "modernc.org/cc/v3"

type Visitor interface {
	Visit(root cc.Node)
}

func NewVisitor(handler Handler) Visitor {
	return &astVisitor{
		handler: handler,
	}
}

type astVisitor struct {
	handler Handler
}

//gocyclo:ignore
func (a *astVisitor) Visit(root cc.Node) {
	cc.Inspect(root, func(n cc.Node, start bool) bool {
		switch n.(type) {
		case *cc.AbstractDeclarator:
			if start {
				a.handler.OnAbstractDeclaratorStart(n.(*cc.AbstractDeclarator))
			} else {
				a.handler.OnAbstractDeclaratorEnd(n.(*cc.AbstractDeclarator))
			}
		case *cc.AdditiveExpression:
			if start {
				a.handler.OnAdditiveExpressionStart(n.(*cc.AdditiveExpression))
			} else {
				a.handler.OnAdditiveExpressionEnd(n.(*cc.AdditiveExpression))
			}
		case *cc.AlignmentSpecifier:
			if start {
				a.handler.OnAlignmentSpecifierStart(n.(*cc.AlignmentSpecifier))
			} else {
				a.handler.OnAlignmentSpecifierEnd(n.(*cc.AlignmentSpecifier))
			}
		case *cc.AndExpression:
			if start {
				a.handler.OnAndExpressionStart(n.(*cc.AndExpression))
			} else {
				a.handler.OnAndExpressionEnd(n.(*cc.AndExpression))
			}
		case *cc.ArgumentExpressionList:
			if start {
				a.handler.OnArgumentExpressionListStart(n.(*cc.ArgumentExpressionList))
			} else {
				a.handler.OnArgumentExpressionListEnd(n.(*cc.ArgumentExpressionList))
			}
		case *cc.Asm:
			if start {
				a.handler.OnAsmStart(n.(*cc.Asm))
			} else {
				a.handler.OnAsmEnd(n.(*cc.Asm))
			}
		case *cc.AsmArgList:
			if start {
				a.handler.OnAsmArgListStart(n.(*cc.AsmArgList))
			} else {
				a.handler.OnAsmArgListEnd(n.(*cc.AsmArgList))
			}
		case *cc.AsmExpressionList:
			if start {
				a.handler.OnAsmExpressionListStart(n.(*cc.AsmExpressionList))
			} else {
				a.handler.OnAsmExpressionListEnd(n.(*cc.AsmExpressionList))
			}
		case *cc.AsmFunctionDefinition:
			if start {
				a.handler.OnAsmFunctionDefinitionStart(n.(*cc.AsmFunctionDefinition))
			} else {
				a.handler.OnAsmFunctionDefinitionEnd(n.(*cc.AsmFunctionDefinition))
			}
		case *cc.AsmIndex:
			if start {
				a.handler.OnAsmIndexStart(n.(*cc.AsmIndex))
			} else {
				a.handler.OnAsmIndexEnd(n.(*cc.AsmIndex))
			}
		case *cc.AsmQualifier:
			if start {
				a.handler.OnAsmQualifierStart(n.(*cc.AsmQualifier))
			} else {
				a.handler.OnAsmQualifierEnd(n.(*cc.AsmQualifier))
			}
		case *cc.AsmQualifierList:
			if start {
				a.handler.OnAsmQualifierListStart(n.(*cc.AsmQualifierList))
			} else {
				a.handler.OnAsmQualifierListEnd(n.(*cc.AsmQualifierList))
			}
		case *cc.AsmStatement:
			if start {
				a.handler.OnAsmStatementStart(n.(*cc.AsmStatement))
			} else {
				a.handler.OnAsmStatementEnd(n.(*cc.AsmStatement))
			}
		case *cc.AssignmentExpression:
			if start {
				a.handler.OnAssignmentExpressionStart(n.(*cc.AssignmentExpression))
			} else {
				a.handler.OnAssignmentExpressionEnd(n.(*cc.AssignmentExpression))
			}
		case *cc.AtomicTypeSpecifier:
			if start {
				a.handler.OnAtomicTypeSpecifierStart(n.(*cc.AtomicTypeSpecifier))
			} else {
				a.handler.OnAtomicTypeSpecifierEnd(n.(*cc.AtomicTypeSpecifier))
			}
		case *cc.AttributeSpecifier:
			if start {
				a.handler.OnAttributeSpecifierStart(n.(*cc.AttributeSpecifier))
			} else {
				a.handler.OnAttributeSpecifierEnd(n.(*cc.AttributeSpecifier))
			}
		case *cc.AttributeSpecifierList:
			if start {
				a.handler.OnAttributeSpecifierListStart(n.(*cc.AttributeSpecifierList))
			} else {
				a.handler.OnAttributeSpecifierListEnd(n.(*cc.AttributeSpecifierList))
			}
		case *cc.AttributeValue:
			if start {
				a.handler.OnAttributeValueStart(n.(*cc.AttributeValue))
			} else {
				a.handler.OnAttributeValueEnd(n.(*cc.AttributeValue))
			}
		case *cc.AttributeValueList:
			if start {
				a.handler.OnAttributeValueListStart(n.(*cc.AttributeValueList))
			} else {
				a.handler.OnAttributeValueListEnd(n.(*cc.AttributeValueList))
			}
		case *cc.BlockItem:
			if start {
				a.handler.OnBlockItemStart(n.(*cc.BlockItem))
			} else {
				a.handler.OnBlockItemEnd(n.(*cc.BlockItem))
			}
		case *cc.BlockItemList:
			if start {
				a.handler.OnBlockItemListStart(n.(*cc.BlockItemList))
			} else {
				a.handler.OnBlockItemListEnd(n.(*cc.BlockItemList))
			}
		case *cc.CastExpression:
			if start {
				a.handler.OnCastExpressionStart(n.(*cc.CastExpression))
			} else {
				a.handler.OnCastExpressionEnd(n.(*cc.CastExpression))
			}
		case *cc.CompoundStatement:
			if start {
				a.handler.OnCompoundStatementStart(n.(*cc.CompoundStatement))
			} else {
				a.handler.OnCompoundStatementEnd(n.(*cc.CompoundStatement))
			}
		case *cc.ConditionalExpression:
			if start {
				a.handler.OnConditionalExpressionStart(n.(*cc.ConditionalExpression))
			} else {
				a.handler.OnConditionalExpressionEnd(n.(*cc.ConditionalExpression))
			}
		case *cc.ConstantExpression:
			if start {
				a.handler.OnConstantExpressionStart(n.(*cc.ConstantExpression))
			} else {
				a.handler.OnConstantExpressionEnd(n.(*cc.ConstantExpression))
			}
		case *cc.Declaration:
			if start {
				a.handler.OnDeclarationStart(n.(*cc.Declaration))
			} else {
				a.handler.OnDeclarationEnd(n.(*cc.Declaration))
			}
		case *cc.DeclarationList:
			if start {
				a.handler.OnDeclarationListStart(n.(*cc.DeclarationList))
			} else {
				a.handler.OnDeclarationListEnd(n.(*cc.DeclarationList))
			}
		case *cc.DeclarationSpecifiers:
			if start {
				a.handler.OnDeclarationSpecifiersStart(n.(*cc.DeclarationSpecifiers))
			} else {
				a.handler.OnDeclarationSpecifiersEnd(n.(*cc.DeclarationSpecifiers))
			}
		case *cc.Declarator:
			if start {
				a.handler.OnDeclaratorStart(n.(*cc.Declarator))
			} else {
				a.handler.OnDeclaratorEnd(n.(*cc.Declarator))
			}
		case *cc.Designation:
			if start {
				a.handler.OnDesignationStart(n.(*cc.Designation))
			} else {
				a.handler.OnDesignationEnd(n.(*cc.Designation))
			}
		case *cc.Designator:
			if start {
				a.handler.OnDesignatorStart(n.(*cc.Designator))
			} else {
				a.handler.OnDesignatorEnd(n.(*cc.Designator))
			}
		case *cc.DesignatorList:
			if start {
				a.handler.OnDesignatorListStart(n.(*cc.DesignatorList))
			} else {
				a.handler.OnDesignatorListEnd(n.(*cc.DesignatorList))
			}
		case *cc.DirectAbstractDeclarator:
			if start {
				a.handler.OnDirectAbstractDeclaratorStart(n.(*cc.DirectAbstractDeclarator))
			} else {
				a.handler.OnDirectAbstractDeclaratorEnd(n.(*cc.DirectAbstractDeclarator))
			}
		case *cc.DirectDeclarator:
			if start {
				a.handler.OnDirectDeclaratorStart(n.(*cc.DirectDeclarator))
			} else {
				a.handler.OnDirectDeclaratorEnd(n.(*cc.DirectDeclarator))
			}
		case *cc.EnumSpecifier:
			if start {
				a.handler.OnEnumSpecifierStart(n.(*cc.EnumSpecifier))
			} else {
				a.handler.OnEnumSpecifierEnd(n.(*cc.EnumSpecifier))
			}
		case *cc.Enumerator:
			if start {
				a.handler.OnEnumeratorStart(n.(*cc.Enumerator))
			} else {
				a.handler.OnEnumeratorEnd(n.(*cc.Enumerator))
			}
		case *cc.EnumeratorList:
			if start {
				a.handler.OnEnumeratorListStart(n.(*cc.EnumeratorList))
			} else {
				a.handler.OnEnumeratorListEnd(n.(*cc.EnumeratorList))
			}
		case *cc.EqualityExpression:
			if start {
				a.handler.OnEqualityExpressionStart(n.(*cc.EqualityExpression))
			} else {
				a.handler.OnEqualityExpressionEnd(n.(*cc.EqualityExpression))
			}
		case *cc.ExclusiveOrExpression:
			if start {
				a.handler.OnExclusiveOrExpressionStart(n.(*cc.ExclusiveOrExpression))
			} else {
				a.handler.OnExclusiveOrExpressionEnd(n.(*cc.ExclusiveOrExpression))
			}
		case *cc.Expression:
			if start {
				a.handler.OnExpressionStart(n.(*cc.Expression))
			} else {
				a.handler.OnExpressionEnd(n.(*cc.Expression))
			}
		case *cc.ExpressionList:
			if start {
				a.handler.OnExpressionListStart(n.(*cc.ExpressionList))
			} else {
				a.handler.OnExpressionListEnd(n.(*cc.ExpressionList))
			}
		case *cc.ExpressionStatement:
			if start {
				a.handler.OnExpressionStatementStart(n.(*cc.ExpressionStatement))
			} else {
				a.handler.OnExpressionStatementEnd(n.(*cc.ExpressionStatement))
			}
		case *cc.ExternalDeclaration:
			if start {
				a.handler.OnExternalDeclarationStart(n.(*cc.ExternalDeclaration))
			} else {
				a.handler.OnExternalDeclarationEnd(n.(*cc.ExternalDeclaration))
			}
		case *cc.FunctionDefinition:
			if start {
				a.handler.OnFunctionDefinitionStart(n.(*cc.FunctionDefinition))
			} else {
				a.handler.OnFunctionDefinitionEnd(n.(*cc.FunctionDefinition))
			}
		case *cc.FunctionSpecifier:
			if start {
				a.handler.OnFunctionSpecifierStart(n.(*cc.FunctionSpecifier))
			} else {
				a.handler.OnFunctionSpecifierEnd(n.(*cc.FunctionSpecifier))
			}
		case *cc.IdentifierList:
			if start {
				a.handler.OnIdentifierListStart(n.(*cc.IdentifierList))
			} else {
				a.handler.OnIdentifierListEnd(n.(*cc.IdentifierList))
			}
		case *cc.InclusiveOrExpression:
			if start {
				a.handler.OnInclusiveOrExpressionStart(n.(*cc.InclusiveOrExpression))
			} else {
				a.handler.OnInclusiveOrExpressionEnd(n.(*cc.InclusiveOrExpression))
			}
		case *cc.InitDeclarator:
			if start {
				a.handler.OnInitDeclaratorStart(n.(*cc.InitDeclarator))
			} else {
				a.handler.OnInitDeclaratorEnd(n.(*cc.InitDeclarator))
			}
		case *cc.InitDeclaratorList:
			if start {
				a.handler.OnInitDeclaratorListStart(n.(*cc.InitDeclaratorList))
			} else {
				a.handler.OnInitDeclaratorListEnd(n.(*cc.InitDeclaratorList))
			}
		case *cc.Initializer:
			if start {
				a.handler.OnInitializerStart(n.(*cc.Initializer))
			} else {
				a.handler.OnInitializerEnd(n.(*cc.Initializer))
			}
		case *cc.InitializerList:
			if start {
				a.handler.OnInitializerListStart(n.(*cc.InitializerList))
			} else {
				a.handler.OnInitializerListEnd(n.(*cc.InitializerList))
			}
		case *cc.IterationStatement:
			if start {
				a.handler.OnIterationStatementStart(n.(*cc.IterationStatement))
			} else {
				a.handler.OnIterationStatementEnd(n.(*cc.IterationStatement))
			}
		case *cc.JumpStatement:
			if start {
				a.handler.OnJumpStatementStart(n.(*cc.JumpStatement))
			} else {
				a.handler.OnJumpStatementEnd(n.(*cc.JumpStatement))
			}
		case *cc.LabelDeclaration:
			if start {
				a.handler.OnLabelDeclarationStart(n.(*cc.LabelDeclaration))
			} else {
				a.handler.OnLabelDeclarationEnd(n.(*cc.LabelDeclaration))
			}
		case *cc.LabeledStatement:
			if start {
				a.handler.OnLabeledStatementStart(n.(*cc.LabeledStatement))
			} else {
				a.handler.OnLabeledStatementEnd(n.(*cc.LabeledStatement))
			}
		case *cc.LogicalAndExpression:
			if start {
				a.handler.OnLogicalAndExpressionStart(n.(*cc.LogicalAndExpression))
			} else {
				a.handler.OnLogicalAndExpressionEnd(n.(*cc.LogicalAndExpression))
			}
		case *cc.LogicalOrExpression:
			if start {
				a.handler.OnLogicalOrExpressionStart(n.(*cc.LogicalOrExpression))
			} else {
				a.handler.OnLogicalOrExpressionEnd(n.(*cc.LogicalOrExpression))
			}
		case *cc.MultiplicativeExpression:
			if start {
				a.handler.OnMultiplicativeExpressionStart(n.(*cc.MultiplicativeExpression))
			} else {
				a.handler.OnMultiplicativeExpressionEnd(n.(*cc.MultiplicativeExpression))
			}
		case *cc.ParameterDeclaration:
			if start {
				a.handler.OnParameterDeclarationStart(n.(*cc.ParameterDeclaration))
			} else {
				a.handler.OnParameterDeclarationEnd(n.(*cc.ParameterDeclaration))
			}
		case *cc.ParameterList:
			if start {
				a.handler.OnParameterListStart(n.(*cc.ParameterList))
			} else {
				a.handler.OnParameterListEnd(n.(*cc.ParameterList))
			}
		case *cc.ParameterTypeList:
			if start {
				a.handler.OnParameterTypeListStart(n.(*cc.ParameterTypeList))
			} else {
				a.handler.OnParameterTypeListEnd(n.(*cc.ParameterTypeList))
			}
		case *cc.Pointer:
			if start {
				a.handler.OnPointerStart(n.(*cc.Pointer))
			} else {
				a.handler.OnPointerEnd(n.(*cc.Pointer))
			}
		case *cc.PostfixExpression:
			if start {
				a.handler.OnPostfixExpressionStart(n.(*cc.PostfixExpression))
			} else {
				a.handler.OnPostfixExpressionEnd(n.(*cc.PostfixExpression))
			}
		case *cc.PragmaSTDC:
			if start {
				a.handler.OnPragmaSTDCStart(n.(*cc.PragmaSTDC))
			} else {
				a.handler.OnPragmaSTDCEnd(n.(*cc.PragmaSTDC))
			}
		case *cc.PrimaryExpression:
			if start {
				a.handler.OnPrimaryExpressionStart(n.(*cc.PrimaryExpression))
			} else {
				a.handler.OnPrimaryExpressionEnd(n.(*cc.PrimaryExpression))
			}
		case *cc.RelationalExpression:
			if start {
				a.handler.OnRelationalExpressionStart(n.(*cc.RelationalExpression))
			} else {
				a.handler.OnRelationalExpressionEnd(n.(*cc.RelationalExpression))
			}
		case *cc.SelectionStatement:
			if start {
				a.handler.OnSelectionStatementStart(n.(*cc.SelectionStatement))
			} else {
				a.handler.OnSelectionStatementEnd(n.(*cc.SelectionStatement))
			}
		case *cc.ShiftExpression:
			if start {
				a.handler.OnShiftExpressionStart(n.(*cc.ShiftExpression))
			} else {
				a.handler.OnShiftExpressionEnd(n.(*cc.ShiftExpression))
			}
		case *cc.SpecifierQualifierList:
			if start {
				a.handler.OnSpecifierQualifierListStart(n.(*cc.SpecifierQualifierList))
			} else {
				a.handler.OnSpecifierQualifierListEnd(n.(*cc.SpecifierQualifierList))
			}
		case *cc.Statement:
			if start {
				a.handler.OnStatementStart(n.(*cc.Statement))
			} else {
				a.handler.OnStatementEnd(n.(*cc.Statement))
			}
		case *cc.StorageClassSpecifier:
			if start {
				a.handler.OnStorageClassSpecifierStart(n.(*cc.StorageClassSpecifier))
			} else {
				a.handler.OnStorageClassSpecifierEnd(n.(*cc.StorageClassSpecifier))
			}
		case *cc.StructDeclaration:
			if start {
				a.handler.OnStructDeclarationStart(n.(*cc.StructDeclaration))
			} else {
				a.handler.OnStructDeclarationEnd(n.(*cc.StructDeclaration))
			}
		case *cc.StructDeclarationList:
			if start {
				a.handler.OnStructDeclarationListStart(n.(*cc.StructDeclarationList))
			} else {
				a.handler.OnStructDeclarationListEnd(n.(*cc.StructDeclarationList))
			}
		case *cc.StructDeclarator:
			if start {
				a.handler.OnStructDeclaratorStart(n.(*cc.StructDeclarator))
			} else {
				a.handler.OnStructDeclaratorEnd(n.(*cc.StructDeclarator))
			}
		case *cc.StructDeclaratorList:
			if start {
				a.handler.OnStructDeclaratorListStart(n.(*cc.StructDeclaratorList))
			} else {
				a.handler.OnStructDeclaratorListEnd(n.(*cc.StructDeclaratorList))
			}
		case *cc.StructOrUnion:
			if start {
				a.handler.OnStructOrUnionStart(n.(*cc.StructOrUnion))
			} else {
				a.handler.OnStructOrUnionEnd(n.(*cc.StructOrUnion))
			}
		case *cc.StructOrUnionSpecifier:
			if start {
				a.handler.OnStructOrUnionSpecifierStart(n.(*cc.StructOrUnionSpecifier))
			} else {
				a.handler.OnStructOrUnionSpecifierEnd(n.(*cc.StructOrUnionSpecifier))
			}
		case *cc.TranslationUnit:
			if start {
				a.handler.OnTranslationUnitStart(n.(*cc.TranslationUnit))
			} else {
				a.handler.OnTranslationUnitEnd(n.(*cc.TranslationUnit))
			}
		case *cc.TypeName:
			if start {
				a.handler.OnTypeNameStart(n.(*cc.TypeName))
			} else {
				a.handler.OnTypeNameEnd(n.(*cc.TypeName))
			}
		case *cc.TypeQualifier:
			if start {
				a.handler.OnTypeQualifierStart(n.(*cc.TypeQualifier))
			} else {
				a.handler.OnTypeQualifierEnd(n.(*cc.TypeQualifier))
			}
		case *cc.TypeQualifiers:
			if start {
				a.handler.OnTypeQualifiersStart(n.(*cc.TypeQualifiers))
			} else {
				a.handler.OnTypeQualifiersEnd(n.(*cc.TypeQualifiers))
			}
		case *cc.TypeSpecifier:
			if start {
				a.handler.OnTypeSpecifierStart(n.(*cc.TypeSpecifier))
			} else {
				a.handler.OnTypeSpecifierEnd(n.(*cc.TypeSpecifier))
			}
		case *cc.UnaryExpression:
			if start {
				a.handler.OnUnaryExpressionStart(n.(*cc.UnaryExpression))
			} else {
				a.handler.OnUnaryExpressionEnd(n.(*cc.UnaryExpression))
			}
		}

		return true
	})
}
