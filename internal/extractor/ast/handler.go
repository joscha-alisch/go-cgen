package ast

import "modernc.org/cc/v3"

type Handler interface {
	Start()
	End()

	OnAbstractDeclaratorStart(n *cc.AbstractDeclarator)
	OnAbstractDeclaratorEnd(n *cc.AbstractDeclarator)

	OnAdditiveExpressionStart(n *cc.AdditiveExpression)
	OnAdditiveExpressionEnd(n *cc.AdditiveExpression)

	OnAlignmentSpecifierStart(n *cc.AlignmentSpecifier)
	OnAlignmentSpecifierEnd(n *cc.AlignmentSpecifier)

	OnAndExpressionStart(n *cc.AndExpression)
	OnAndExpressionEnd(n *cc.AndExpression)

	OnArgumentExpressionListStart(n *cc.ArgumentExpressionList)
	OnArgumentExpressionListEnd(n *cc.ArgumentExpressionList)

	OnAsmStart(n *cc.Asm)
	OnAsmEnd(n *cc.Asm)

	OnAsmArgListStart(n *cc.AsmArgList)
	OnAsmArgListEnd(n *cc.AsmArgList)

	OnAsmExpressionListStart(n *cc.AsmExpressionList)
	OnAsmExpressionListEnd(n *cc.AsmExpressionList)

	OnAsmFunctionDefinitionStart(n *cc.AsmFunctionDefinition)
	OnAsmFunctionDefinitionEnd(n *cc.AsmFunctionDefinition)

	OnAsmIndexStart(n *cc.AsmIndex)
	OnAsmIndexEnd(n *cc.AsmIndex)

	OnAsmQualifierStart(n *cc.AsmQualifier)
	OnAsmQualifierEnd(n *cc.AsmQualifier)

	OnAsmQualifierListStart(n *cc.AsmQualifierList)
	OnAsmQualifierListEnd(n *cc.AsmQualifierList)

	OnAsmStatementStart(n *cc.AsmStatement)
	OnAsmStatementEnd(n *cc.AsmStatement)

	OnAssignmentExpressionStart(n *cc.AssignmentExpression)
	OnAssignmentExpressionEnd(n *cc.AssignmentExpression)

	OnAtomicTypeSpecifierStart(n *cc.AtomicTypeSpecifier)
	OnAtomicTypeSpecifierEnd(n *cc.AtomicTypeSpecifier)

	OnAttributeSpecifierStart(n *cc.AttributeSpecifier)
	OnAttributeSpecifierEnd(n *cc.AttributeSpecifier)

	OnAttributeSpecifierListStart(n *cc.AttributeSpecifierList)
	OnAttributeSpecifierListEnd(n *cc.AttributeSpecifierList)

	OnAttributeValueStart(n *cc.AttributeValue)
	OnAttributeValueEnd(n *cc.AttributeValue)

	OnAttributeValueListStart(n *cc.AttributeValueList)
	OnAttributeValueListEnd(n *cc.AttributeValueList)

	OnBlockItemStart(n *cc.BlockItem)
	OnBlockItemEnd(n *cc.BlockItem)

	OnBlockItemListStart(n *cc.BlockItemList)
	OnBlockItemListEnd(n *cc.BlockItemList)

	OnCastExpressionStart(n *cc.CastExpression)
	OnCastExpressionEnd(n *cc.CastExpression)

	OnCompoundStatementStart(n *cc.CompoundStatement)
	OnCompoundStatementEnd(n *cc.CompoundStatement)

	OnConditionalExpressionStart(n *cc.ConditionalExpression)
	OnConditionalExpressionEnd(n *cc.ConditionalExpression)

	OnConstantExpressionStart(n *cc.ConstantExpression)
	OnConstantExpressionEnd(n *cc.ConstantExpression)

	OnDeclarationStart(n *cc.Declaration)
	OnDeclarationEnd(n *cc.Declaration)

	OnDeclarationListStart(n *cc.DeclarationList)
	OnDeclarationListEnd(n *cc.DeclarationList)

	OnDeclarationSpecifiersStart(n *cc.DeclarationSpecifiers)
	OnDeclarationSpecifiersEnd(n *cc.DeclarationSpecifiers)

	OnDeclaratorStart(n *cc.Declarator)
	OnDeclaratorEnd(n *cc.Declarator)

	OnDesignationStart(n *cc.Designation)
	OnDesignationEnd(n *cc.Designation)

	OnDesignatorStart(n *cc.Designator)
	OnDesignatorEnd(n *cc.Designator)

	OnDesignatorListStart(n *cc.DesignatorList)
	OnDesignatorListEnd(n *cc.DesignatorList)

	OnDirectAbstractDeclaratorStart(n *cc.DirectAbstractDeclarator)
	OnDirectAbstractDeclaratorEnd(n *cc.DirectAbstractDeclarator)

	OnDirectDeclaratorStart(n *cc.DirectDeclarator)
	OnDirectDeclaratorEnd(n *cc.DirectDeclarator)

	OnEnumSpecifierStart(n *cc.EnumSpecifier)
	OnEnumSpecifierEnd(n *cc.EnumSpecifier)

	OnEnumeratorStart(n *cc.Enumerator)
	OnEnumeratorEnd(n *cc.Enumerator)

	OnEnumeratorListStart(n *cc.EnumeratorList)
	OnEnumeratorListEnd(n *cc.EnumeratorList)

	OnEqualityExpressionStart(n *cc.EqualityExpression)
	OnEqualityExpressionEnd(n *cc.EqualityExpression)

	OnExclusiveOrExpressionStart(n *cc.ExclusiveOrExpression)
	OnExclusiveOrExpressionEnd(n *cc.ExclusiveOrExpression)

	OnExpressionStart(n *cc.Expression)
	OnExpressionEnd(n *cc.Expression)

	OnExpressionListStart(n *cc.ExpressionList)
	OnExpressionListEnd(n *cc.ExpressionList)

	OnExpressionStatementStart(n *cc.ExpressionStatement)
	OnExpressionStatementEnd(n *cc.ExpressionStatement)

	OnExternalDeclarationStart(n *cc.ExternalDeclaration)
	OnExternalDeclarationEnd(n *cc.ExternalDeclaration)

	OnFunctionDefinitionStart(n *cc.FunctionDefinition)
	OnFunctionDefinitionEnd(n *cc.FunctionDefinition)

	OnFunctionSpecifierStart(n *cc.FunctionSpecifier)
	OnFunctionSpecifierEnd(n *cc.FunctionSpecifier)

	OnIdentifierListStart(n *cc.IdentifierList)
	OnIdentifierListEnd(n *cc.IdentifierList)

	OnInclusiveOrExpressionStart(n *cc.InclusiveOrExpression)
	OnInclusiveOrExpressionEnd(n *cc.InclusiveOrExpression)

	OnInitDeclaratorStart(n *cc.InitDeclarator)
	OnInitDeclaratorEnd(n *cc.InitDeclarator)

	OnInitDeclaratorListStart(n *cc.InitDeclaratorList)
	OnInitDeclaratorListEnd(n *cc.InitDeclaratorList)

	OnInitializerStart(n *cc.Initializer)
	OnInitializerEnd(n *cc.Initializer)

	OnInitializerListStart(n *cc.InitializerList)
	OnInitializerListEnd(n *cc.InitializerList)

	OnIterationStatementStart(n *cc.IterationStatement)
	OnIterationStatementEnd(n *cc.IterationStatement)

	OnJumpStatementStart(n *cc.JumpStatement)
	OnJumpStatementEnd(n *cc.JumpStatement)

	OnLabelDeclarationStart(n *cc.LabelDeclaration)
	OnLabelDeclarationEnd(n *cc.LabelDeclaration)

	OnLabeledStatementStart(n *cc.LabeledStatement)
	OnLabeledStatementEnd(n *cc.LabeledStatement)

	OnLogicalAndExpressionStart(n *cc.LogicalAndExpression)
	OnLogicalAndExpressionEnd(n *cc.LogicalAndExpression)

	OnLogicalOrExpressionStart(n *cc.LogicalOrExpression)
	OnLogicalOrExpressionEnd(n *cc.LogicalOrExpression)

	OnMultiplicativeExpressionStart(n *cc.MultiplicativeExpression)
	OnMultiplicativeExpressionEnd(n *cc.MultiplicativeExpression)

	OnParameterDeclarationStart(n *cc.ParameterDeclaration)
	OnParameterDeclarationEnd(n *cc.ParameterDeclaration)

	OnParameterListStart(n *cc.ParameterList)
	OnParameterListEnd(n *cc.ParameterList)

	OnParameterTypeListStart(n *cc.ParameterTypeList)
	OnParameterTypeListEnd(n *cc.ParameterTypeList)

	OnPointerStart(n *cc.Pointer)
	OnPointerEnd(n *cc.Pointer)

	OnPostfixExpressionStart(n *cc.PostfixExpression)
	OnPostfixExpressionEnd(n *cc.PostfixExpression)

	OnPragmaSTDCStart(n *cc.PragmaSTDC)
	OnPragmaSTDCEnd(n *cc.PragmaSTDC)

	OnPrimaryExpressionStart(n *cc.PrimaryExpression)
	OnPrimaryExpressionEnd(n *cc.PrimaryExpression)

	OnRelationalExpressionStart(n *cc.RelationalExpression)
	OnRelationalExpressionEnd(n *cc.RelationalExpression)

	OnSelectionStatementStart(n *cc.SelectionStatement)
	OnSelectionStatementEnd(n *cc.SelectionStatement)

	OnShiftExpressionStart(n *cc.ShiftExpression)
	OnShiftExpressionEnd(n *cc.ShiftExpression)

	OnSpecifierQualifierListStart(n *cc.SpecifierQualifierList)
	OnSpecifierQualifierListEnd(n *cc.SpecifierQualifierList)

	OnStatementStart(n *cc.Statement)
	OnStatementEnd(n *cc.Statement)

	OnStorageClassSpecifierStart(n *cc.StorageClassSpecifier)
	OnStorageClassSpecifierEnd(n *cc.StorageClassSpecifier)

	OnStructDeclarationStart(n *cc.StructDeclaration)
	OnStructDeclarationEnd(n *cc.StructDeclaration)

	OnStructDeclarationListStart(n *cc.StructDeclarationList)
	OnStructDeclarationListEnd(n *cc.StructDeclarationList)

	OnStructDeclaratorStart(n *cc.StructDeclarator)
	OnStructDeclaratorEnd(n *cc.StructDeclarator)

	OnStructDeclaratorListStart(n *cc.StructDeclaratorList)
	OnStructDeclaratorListEnd(n *cc.StructDeclaratorList)

	OnStructOrUnionStart(n *cc.StructOrUnion)
	OnStructOrUnionEnd(n *cc.StructOrUnion)

	OnStructOrUnionSpecifierStart(n *cc.StructOrUnionSpecifier)
	OnStructOrUnionSpecifierEnd(n *cc.StructOrUnionSpecifier)

	OnTranslationUnitStart(n *cc.TranslationUnit)
	OnTranslationUnitEnd(n *cc.TranslationUnit)

	OnTypeNameStart(n *cc.TypeName)
	OnTypeNameEnd(n *cc.TypeName)

	OnTypeQualifierStart(n *cc.TypeQualifier)
	OnTypeQualifierEnd(n *cc.TypeQualifier)

	OnTypeQualifiersStart(n *cc.TypeQualifiers)
	OnTypeQualifiersEnd(n *cc.TypeQualifiers)

	OnTypeSpecifierStart(n *cc.TypeSpecifier)
	OnTypeSpecifierEnd(n *cc.TypeSpecifier)

	OnUnaryExpressionStart(n *cc.UnaryExpression)
	OnUnaryExpressionEnd(n *cc.UnaryExpression)
}

type BaseHandler struct{}

func (h *BaseHandler) Start() {}
func (h *BaseHandler) End()   {}

func (h *BaseHandler) OnAbstractDeclaratorStart(*cc.AbstractDeclarator) {}
func (h *BaseHandler) OnAbstractDeclaratorEnd(*cc.AbstractDeclarator)   {}

func (h *BaseHandler) OnAdditiveExpressionStart(*cc.AdditiveExpression) {}
func (h *BaseHandler) OnAdditiveExpressionEnd(*cc.AdditiveExpression)   {}

func (h *BaseHandler) OnAlignmentSpecifierStart(*cc.AlignmentSpecifier) {}
func (h *BaseHandler) OnAlignmentSpecifierEnd(*cc.AlignmentSpecifier)   {}

func (h *BaseHandler) OnAndExpressionStart(*cc.AndExpression) {}
func (h *BaseHandler) OnAndExpressionEnd(*cc.AndExpression)   {}

func (h *BaseHandler) OnArgumentExpressionListStart(*cc.ArgumentExpressionList) {}
func (h *BaseHandler) OnArgumentExpressionListEnd(*cc.ArgumentExpressionList)   {}

func (h *BaseHandler) OnAsmStart(*cc.Asm) {}
func (h *BaseHandler) OnAsmEnd(*cc.Asm)   {}

func (h *BaseHandler) OnAsmArgListStart(*cc.AsmArgList) {}
func (h *BaseHandler) OnAsmArgListEnd(*cc.AsmArgList)   {}

func (h *BaseHandler) OnAsmExpressionListStart(*cc.AsmExpressionList) {}
func (h *BaseHandler) OnAsmExpressionListEnd(*cc.AsmExpressionList)   {}

func (h *BaseHandler) OnAsmFunctionDefinitionStart(*cc.AsmFunctionDefinition) {}
func (h *BaseHandler) OnAsmFunctionDefinitionEnd(*cc.AsmFunctionDefinition)   {}

func (h *BaseHandler) OnAsmIndexStart(*cc.AsmIndex) {}
func (h *BaseHandler) OnAsmIndexEnd(*cc.AsmIndex)   {}

func (h *BaseHandler) OnAsmQualifierStart(*cc.AsmQualifier) {}
func (h *BaseHandler) OnAsmQualifierEnd(*cc.AsmQualifier)   {}

func (h *BaseHandler) OnAsmQualifierListStart(*cc.AsmQualifierList) {}
func (h *BaseHandler) OnAsmQualifierListEnd(*cc.AsmQualifierList)   {}

func (h *BaseHandler) OnAsmStatementStart(*cc.AsmStatement) {}
func (h *BaseHandler) OnAsmStatementEnd(*cc.AsmStatement)   {}

func (h *BaseHandler) OnAssignmentExpressionStart(*cc.AssignmentExpression) {}
func (h *BaseHandler) OnAssignmentExpressionEnd(*cc.AssignmentExpression)   {}

func (h *BaseHandler) OnAtomicTypeSpecifierStart(*cc.AtomicTypeSpecifier) {}
func (h *BaseHandler) OnAtomicTypeSpecifierEnd(*cc.AtomicTypeSpecifier)   {}

func (h *BaseHandler) OnAttributeSpecifierStart(*cc.AttributeSpecifier) {}
func (h *BaseHandler) OnAttributeSpecifierEnd(*cc.AttributeSpecifier)   {}

func (h *BaseHandler) OnAttributeSpecifierListStart(*cc.AttributeSpecifierList) {}
func (h *BaseHandler) OnAttributeSpecifierListEnd(*cc.AttributeSpecifierList)   {}

func (h *BaseHandler) OnAttributeValueStart(*cc.AttributeValue) {}
func (h *BaseHandler) OnAttributeValueEnd(*cc.AttributeValue)   {}

func (h *BaseHandler) OnAttributeValueListStart(*cc.AttributeValueList) {}
func (h *BaseHandler) OnAttributeValueListEnd(*cc.AttributeValueList)   {}

func (h *BaseHandler) OnBlockItemStart(*cc.BlockItem) {}
func (h *BaseHandler) OnBlockItemEnd(*cc.BlockItem)   {}

func (h *BaseHandler) OnBlockItemListStart(*cc.BlockItemList) {}
func (h *BaseHandler) OnBlockItemListEnd(*cc.BlockItemList)   {}

func (h *BaseHandler) OnCastExpressionStart(*cc.CastExpression) {}
func (h *BaseHandler) OnCastExpressionEnd(*cc.CastExpression)   {}

func (h *BaseHandler) OnCompoundStatementStart(*cc.CompoundStatement) {}
func (h *BaseHandler) OnCompoundStatementEnd(*cc.CompoundStatement)   {}

func (h *BaseHandler) OnConditionalExpressionStart(*cc.ConditionalExpression) {}
func (h *BaseHandler) OnConditionalExpressionEnd(*cc.ConditionalExpression)   {}

func (h *BaseHandler) OnConstantExpressionStart(*cc.ConstantExpression) {}
func (h *BaseHandler) OnConstantExpressionEnd(*cc.ConstantExpression)   {}

func (h *BaseHandler) OnDeclarationStart(*cc.Declaration) {}
func (h *BaseHandler) OnDeclarationEnd(*cc.Declaration)   {}

func (h *BaseHandler) OnDeclarationListStart(*cc.DeclarationList) {}
func (h *BaseHandler) OnDeclarationListEnd(*cc.DeclarationList)   {}

func (h *BaseHandler) OnDeclarationSpecifiersStart(*cc.DeclarationSpecifiers) {}
func (h *BaseHandler) OnDeclarationSpecifiersEnd(*cc.DeclarationSpecifiers)   {}

func (h *BaseHandler) OnDeclaratorStart(*cc.Declarator) {}
func (h *BaseHandler) OnDeclaratorEnd(*cc.Declarator)   {}

func (h *BaseHandler) OnDesignationStart(*cc.Designation) {}
func (h *BaseHandler) OnDesignationEnd(*cc.Designation)   {}

func (h *BaseHandler) OnDesignatorStart(*cc.Designator) {}
func (h *BaseHandler) OnDesignatorEnd(*cc.Designator)   {}

func (h *BaseHandler) OnDesignatorListStart(*cc.DesignatorList) {}
func (h *BaseHandler) OnDesignatorListEnd(*cc.DesignatorList)   {}

func (h *BaseHandler) OnDirectAbstractDeclaratorStart(*cc.DirectAbstractDeclarator) {}
func (h *BaseHandler) OnDirectAbstractDeclaratorEnd(*cc.DirectAbstractDeclarator)   {}

func (h *BaseHandler) OnDirectDeclaratorStart(*cc.DirectDeclarator) {}
func (h *BaseHandler) OnDirectDeclaratorEnd(*cc.DirectDeclarator)   {}

func (h *BaseHandler) OnEnumSpecifierStart(*cc.EnumSpecifier) {}
func (h *BaseHandler) OnEnumSpecifierEnd(*cc.EnumSpecifier)   {}

func (h *BaseHandler) OnEnumeratorStart(*cc.Enumerator) {}
func (h *BaseHandler) OnEnumeratorEnd(*cc.Enumerator)   {}

func (h *BaseHandler) OnEnumeratorListStart(*cc.EnumeratorList) {}
func (h *BaseHandler) OnEnumeratorListEnd(*cc.EnumeratorList)   {}

func (h *BaseHandler) OnEqualityExpressionStart(*cc.EqualityExpression) {}
func (h *BaseHandler) OnEqualityExpressionEnd(*cc.EqualityExpression)   {}

func (h *BaseHandler) OnExclusiveOrExpressionStart(*cc.ExclusiveOrExpression) {}
func (h *BaseHandler) OnExclusiveOrExpressionEnd(*cc.ExclusiveOrExpression)   {}

func (h *BaseHandler) OnExpressionStart(*cc.Expression) {}
func (h *BaseHandler) OnExpressionEnd(*cc.Expression)   {}

func (h *BaseHandler) OnExpressionListStart(*cc.ExpressionList) {}
func (h *BaseHandler) OnExpressionListEnd(*cc.ExpressionList)   {}

func (h *BaseHandler) OnExpressionStatementStart(*cc.ExpressionStatement) {}
func (h *BaseHandler) OnExpressionStatementEnd(*cc.ExpressionStatement)   {}

func (h *BaseHandler) OnExternalDeclarationStart(*cc.ExternalDeclaration) {}
func (h *BaseHandler) OnExternalDeclarationEnd(*cc.ExternalDeclaration)   {}

func (h *BaseHandler) OnFunctionDefinitionStart(*cc.FunctionDefinition) {}
func (h *BaseHandler) OnFunctionDefinitionEnd(*cc.FunctionDefinition)   {}

func (h *BaseHandler) OnFunctionSpecifierStart(*cc.FunctionSpecifier) {}
func (h *BaseHandler) OnFunctionSpecifierEnd(*cc.FunctionSpecifier)   {}

func (h *BaseHandler) OnIdentifierListStart(*cc.IdentifierList) {}
func (h *BaseHandler) OnIdentifierListEnd(*cc.IdentifierList)   {}

func (h *BaseHandler) OnInclusiveOrExpressionStart(*cc.InclusiveOrExpression) {}
func (h *BaseHandler) OnInclusiveOrExpressionEnd(*cc.InclusiveOrExpression)   {}

func (h *BaseHandler) OnInitDeclaratorStart(*cc.InitDeclarator) {}
func (h *BaseHandler) OnInitDeclaratorEnd(*cc.InitDeclarator)   {}

func (h *BaseHandler) OnInitDeclaratorListStart(*cc.InitDeclaratorList) {}
func (h *BaseHandler) OnInitDeclaratorListEnd(*cc.InitDeclaratorList)   {}

func (h *BaseHandler) OnInitializerStart(*cc.Initializer) {}
func (h *BaseHandler) OnInitializerEnd(*cc.Initializer)   {}

func (h *BaseHandler) OnInitializerListStart(*cc.InitializerList) {}
func (h *BaseHandler) OnInitializerListEnd(*cc.InitializerList)   {}

func (h *BaseHandler) OnIterationStatementStart(*cc.IterationStatement) {}
func (h *BaseHandler) OnIterationStatementEnd(*cc.IterationStatement)   {}

func (h *BaseHandler) OnJumpStatementStart(*cc.JumpStatement) {}
func (h *BaseHandler) OnJumpStatementEnd(*cc.JumpStatement)   {}

func (h *BaseHandler) OnLabelDeclarationStart(*cc.LabelDeclaration) {}
func (h *BaseHandler) OnLabelDeclarationEnd(*cc.LabelDeclaration)   {}

func (h *BaseHandler) OnLabeledStatementStart(*cc.LabeledStatement) {}
func (h *BaseHandler) OnLabeledStatementEnd(*cc.LabeledStatement)   {}

func (h *BaseHandler) OnLogicalAndExpressionStart(*cc.LogicalAndExpression) {}
func (h *BaseHandler) OnLogicalAndExpressionEnd(*cc.LogicalAndExpression)   {}

func (h *BaseHandler) OnLogicalOrExpressionStart(*cc.LogicalOrExpression) {}
func (h *BaseHandler) OnLogicalOrExpressionEnd(*cc.LogicalOrExpression)   {}

func (h *BaseHandler) OnMultiplicativeExpressionStart(*cc.MultiplicativeExpression) {}
func (h *BaseHandler) OnMultiplicativeExpressionEnd(*cc.MultiplicativeExpression)   {}

func (h *BaseHandler) OnParameterDeclarationStart(*cc.ParameterDeclaration) {}
func (h *BaseHandler) OnParameterDeclarationEnd(*cc.ParameterDeclaration)   {}

func (h *BaseHandler) OnParameterListStart(*cc.ParameterList) {}
func (h *BaseHandler) OnParameterListEnd(*cc.ParameterList)   {}

func (h *BaseHandler) OnParameterTypeListStart(*cc.ParameterTypeList) {}
func (h *BaseHandler) OnParameterTypeListEnd(*cc.ParameterTypeList)   {}

func (h *BaseHandler) OnPointerStart(*cc.Pointer) {}
func (h *BaseHandler) OnPointerEnd(*cc.Pointer)   {}

func (h *BaseHandler) OnPostfixExpressionStart(*cc.PostfixExpression) {}
func (h *BaseHandler) OnPostfixExpressionEnd(*cc.PostfixExpression)   {}

func (h *BaseHandler) OnPragmaSTDCStart(*cc.PragmaSTDC) {}
func (h *BaseHandler) OnPragmaSTDCEnd(*cc.PragmaSTDC)   {}

func (h *BaseHandler) OnPrimaryExpressionStart(*cc.PrimaryExpression) {}
func (h *BaseHandler) OnPrimaryExpressionEnd(*cc.PrimaryExpression)   {}

func (h *BaseHandler) OnRelationalExpressionStart(*cc.RelationalExpression) {}
func (h *BaseHandler) OnRelationalExpressionEnd(*cc.RelationalExpression)   {}

func (h *BaseHandler) OnSelectionStatementStart(*cc.SelectionStatement) {}
func (h *BaseHandler) OnSelectionStatementEnd(*cc.SelectionStatement)   {}

func (h *BaseHandler) OnShiftExpressionStart(*cc.ShiftExpression) {}
func (h *BaseHandler) OnShiftExpressionEnd(*cc.ShiftExpression)   {}

func (h *BaseHandler) OnSpecifierQualifierListStart(*cc.SpecifierQualifierList) {}
func (h *BaseHandler) OnSpecifierQualifierListEnd(*cc.SpecifierQualifierList)   {}

func (h *BaseHandler) OnStatementStart(*cc.Statement) {}
func (h *BaseHandler) OnStatementEnd(*cc.Statement)   {}

func (h *BaseHandler) OnStorageClassSpecifierStart(*cc.StorageClassSpecifier) {}
func (h *BaseHandler) OnStorageClassSpecifierEnd(*cc.StorageClassSpecifier)   {}

func (h *BaseHandler) OnStructDeclarationStart(*cc.StructDeclaration) {}
func (h *BaseHandler) OnStructDeclarationEnd(*cc.StructDeclaration)   {}

func (h *BaseHandler) OnStructDeclarationListStart(*cc.StructDeclarationList) {}
func (h *BaseHandler) OnStructDeclarationListEnd(*cc.StructDeclarationList)   {}

func (h *BaseHandler) OnStructDeclaratorStart(*cc.StructDeclarator) {}
func (h *BaseHandler) OnStructDeclaratorEnd(*cc.StructDeclarator)   {}

func (h *BaseHandler) OnStructDeclaratorListStart(*cc.StructDeclaratorList) {}
func (h *BaseHandler) OnStructDeclaratorListEnd(*cc.StructDeclaratorList)   {}

func (h *BaseHandler) OnStructOrUnionStart(*cc.StructOrUnion) {}
func (h *BaseHandler) OnStructOrUnionEnd(*cc.StructOrUnion)   {}

func (h *BaseHandler) OnStructOrUnionSpecifierStart(*cc.StructOrUnionSpecifier) {}
func (h *BaseHandler) OnStructOrUnionSpecifierEnd(*cc.StructOrUnionSpecifier)   {}

func (h *BaseHandler) OnTranslationUnitStart(*cc.TranslationUnit) {}
func (h *BaseHandler) OnTranslationUnitEnd(*cc.TranslationUnit)   {}

func (h *BaseHandler) OnTypeNameStart(*cc.TypeName) {}
func (h *BaseHandler) OnTypeNameEnd(*cc.TypeName)   {}

func (h *BaseHandler) OnTypeQualifierStart(*cc.TypeQualifier) {}
func (h *BaseHandler) OnTypeQualifierEnd(*cc.TypeQualifier)   {}

func (h *BaseHandler) OnTypeQualifiersStart(*cc.TypeQualifiers) {}
func (h *BaseHandler) OnTypeQualifiersEnd(*cc.TypeQualifiers)   {}

func (h *BaseHandler) OnTypeSpecifierStart(*cc.TypeSpecifier) {}
func (h *BaseHandler) OnTypeSpecifierEnd(*cc.TypeSpecifier)   {}

func (h *BaseHandler) OnUnaryExpressionStart(*cc.UnaryExpression) {}
func (h *BaseHandler) OnUnaryExpressionEnd(*cc.UnaryExpression)   {}
