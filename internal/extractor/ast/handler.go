package ast

import "modernc.org/cc/v3"

// Handler is a callback handler used in Visitor.  During traversal of the cc.AST, the methods on Handler are called on
// start and end of node traversal.
type Handler interface {
	Start()
	End()

	// OnAbstractDeclaratorStart is called during the traversal of the AST, whenever a node of type *cc.AbstractDeclarator is entered.
	OnAbstractDeclaratorStart(n *cc.AbstractDeclarator)
	// OnAbstractDeclaratorEnd is called during the traversal of the AST, whenever a node of type *cc.AbstractDeclarator is exited.
	OnAbstractDeclaratorEnd(n *cc.AbstractDeclarator)

	// OnAdditiveExpressionStart is called during the traversal of the AST, whenever a node of type *cc.AdditiveExpression is entered.
	OnAdditiveExpressionStart(n *cc.AdditiveExpression)
	// OnAdditiveExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.AdditiveExpression is exited.
	OnAdditiveExpressionEnd(n *cc.AdditiveExpression)

	// OnAlignmentSpecifierStart is called during the traversal of the AST, whenever a node of type *cc.AlignmentSpecifier is entered.
	OnAlignmentSpecifierStart(n *cc.AlignmentSpecifier)
	// OnAlignmentSpecifierEnd is called during the traversal of the AST, whenever a node of type *cc.AlignmentSpecifier is exited.
	OnAlignmentSpecifierEnd(n *cc.AlignmentSpecifier)

	// OnAndExpressionStart is called during the traversal of the AST, whenever a node of type *cc.AndExpression is entered.
	OnAndExpressionStart(n *cc.AndExpression)
	// OnAndExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.AndExpression is exited.
	OnAndExpressionEnd(n *cc.AndExpression)

	// OnArgumentExpressionListStart is called during the traversal of the AST, whenever a node of type *cc.ArgumentExpressionList is entered.
	OnArgumentExpressionListStart(n *cc.ArgumentExpressionList)
	// OnArgumentExpressionListEnd is called during the traversal of the AST, whenever a node of type *cc.ArgumentExpressionList is exited.
	OnArgumentExpressionListEnd(n *cc.ArgumentExpressionList)

	// OnAsmStart is called during the traversal of the AST, whenever a node of type *cc.Asm is entered.
	OnAsmStart(n *cc.Asm)
	// OnAsmEnd is called during the traversal of the AST, whenever a node of type *cc.Asm is exited.
	OnAsmEnd(n *cc.Asm)

	// OnAsmArgListStart is called during the traversal of the AST, whenever a node of type *cc.AsmArgList is entered.
	OnAsmArgListStart(n *cc.AsmArgList)
	// OnAsmArgListEnd is called during the traversal of the AST, whenever a node of type *cc.AsmArgList is exited.
	OnAsmArgListEnd(n *cc.AsmArgList)

	// OnAsmExpressionListStart is called during the traversal of the AST, whenever a node of type *cc.AsmExpressionList is entered.
	OnAsmExpressionListStart(n *cc.AsmExpressionList)
	// OnAsmExpressionListEnd is called during the traversal of the AST, whenever a node of type *cc.AsmExpressionList is exited.
	OnAsmExpressionListEnd(n *cc.AsmExpressionList)

	// OnAsmFunctionDefinitionStart is called during the traversal of the AST, whenever a node of type *cc.AsmFunctionDefinition is entered.
	OnAsmFunctionDefinitionStart(n *cc.AsmFunctionDefinition)
	// OnAsmFunctionDefinitionEnd is called during the traversal of the AST, whenever a node of type *cc.AsmFunctionDefinition is exited.
	OnAsmFunctionDefinitionEnd(n *cc.AsmFunctionDefinition)

	// OnAsmIndexStart is called during the traversal of the AST, whenever a node of type *cc.AsmIndex is entered.
	OnAsmIndexStart(n *cc.AsmIndex)
	// OnAsmIndexEnd is called during the traversal of the AST, whenever a node of type *cc.AsmIndex is exited.
	OnAsmIndexEnd(n *cc.AsmIndex)

	// OnAsmQualifierStart is called during the traversal of the AST, whenever a node of type *cc.AsmQualifier is entered.
	OnAsmQualifierStart(n *cc.AsmQualifier)
	// OnAsmQualifierEnd is called during the traversal of the AST, whenever a node of type *cc.AsmQualifier is exited.
	OnAsmQualifierEnd(n *cc.AsmQualifier)

	// OnAsmQualifierListStart is called during the traversal of the AST, whenever a node of type *cc.AsmQualifierList is entered.
	OnAsmQualifierListStart(n *cc.AsmQualifierList)
	// OnAsmQualifierListEnd is called during the traversal of the AST, whenever a node of type *cc.AsmQualifierList is exited.
	OnAsmQualifierListEnd(n *cc.AsmQualifierList)

	// OnAsmStatementStart is called during the traversal of the AST, whenever a node of type *cc.AsmStatement is entered.
	OnAsmStatementStart(n *cc.AsmStatement)
	// OnAsmStatementEnd is called during the traversal of the AST, whenever a node of type *cc.AsmStatement is exited.
	OnAsmStatementEnd(n *cc.AsmStatement)

	// OnAssignmentExpressionStart is called during the traversal of the AST, whenever a node of type *cc.AssignmentExpression is entered.
	OnAssignmentExpressionStart(n *cc.AssignmentExpression)
	// OnAssignmentExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.AssignmentExpression is exited.
	OnAssignmentExpressionEnd(n *cc.AssignmentExpression)

	// OnAtomicTypeSpecifierStart is called during the traversal of the AST, whenever a node of type *cc.AtomicTypeSpecifier is entered.
	OnAtomicTypeSpecifierStart(n *cc.AtomicTypeSpecifier)
	// OnAtomicTypeSpecifierEnd is called during the traversal of the AST, whenever a node of type *cc.AtomicTypeSpecifier is exited.
	OnAtomicTypeSpecifierEnd(n *cc.AtomicTypeSpecifier)

	// OnAttributeSpecifierStart is called during the traversal of the AST, whenever a node of type *cc.AttributeSpecifier is entered.
	OnAttributeSpecifierStart(n *cc.AttributeSpecifier)
	// OnAttributeSpecifierEnd is called during the traversal of the AST, whenever a node of type *cc.AttributeSpecifier is exited.
	OnAttributeSpecifierEnd(n *cc.AttributeSpecifier)

	// OnAttributeSpecifierListStart is called during the traversal of the AST, whenever a node of type *cc.AttributeSpecifierList is entered.
	OnAttributeSpecifierListStart(n *cc.AttributeSpecifierList)
	// OnAttributeSpecifierListEnd is called during the traversal of the AST, whenever a node of type *cc.AttributeSpecifierList is exited.
	OnAttributeSpecifierListEnd(n *cc.AttributeSpecifierList)

	// OnAttributeValueStart is called during the traversal of the AST, whenever a node of type *cc.AttributeValue is entered.
	OnAttributeValueStart(n *cc.AttributeValue)
	// OnAttributeValueEnd is called during the traversal of the AST, whenever a node of type *cc.AttributeValue is exited.
	OnAttributeValueEnd(n *cc.AttributeValue)

	// OnAttributeValueListStart is called during the traversal of the AST, whenever a node of type *cc.AttributeValueList is entered.
	OnAttributeValueListStart(n *cc.AttributeValueList)
	// OnAttributeValueListEnd is called during the traversal of the AST, whenever a node of type *cc.AttributeValueList is exited.
	OnAttributeValueListEnd(n *cc.AttributeValueList)

	// OnBlockItemStart is called during the traversal of the AST, whenever a node of type *cc.BlockItem is entered.
	OnBlockItemStart(n *cc.BlockItem)
	// OnBlockItemEnd is called during the traversal of the AST, whenever a node of type *cc.BlockItem is exited.
	OnBlockItemEnd(n *cc.BlockItem)

	// OnBlockItemListStart is called during the traversal of the AST, whenever a node of type *cc.BlockItemList is entered.
	OnBlockItemListStart(n *cc.BlockItemList)
	// OnBlockItemListEnd is called during the traversal of the AST, whenever a node of type *cc.BlockItemList is exited.
	OnBlockItemListEnd(n *cc.BlockItemList)

	// OnCastExpressionStart is called during the traversal of the AST, whenever a node of type *cc.CastExpression is entered.
	OnCastExpressionStart(n *cc.CastExpression)
	// OnCastExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.CastExpression is exited.
	OnCastExpressionEnd(n *cc.CastExpression)

	// OnCompoundStatementStart is called during the traversal of the AST, whenever a node of type *cc.CompoundStatement is entered.
	OnCompoundStatementStart(n *cc.CompoundStatement)
	// OnCompoundStatementEnd is called during the traversal of the AST, whenever a node of type *cc.CompoundStatement is exited.
	OnCompoundStatementEnd(n *cc.CompoundStatement)

	// OnConditionalExpressionStart is called during the traversal of the AST, whenever a node of type *cc.ConditionalExpression is entered.
	OnConditionalExpressionStart(n *cc.ConditionalExpression)
	// OnConditionalExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.ConditionalExpression is exited.
	OnConditionalExpressionEnd(n *cc.ConditionalExpression)

	// OnConstantExpressionStart is called during the traversal of the AST, whenever a node of type *cc.ConstantExpression is entered.
	OnConstantExpressionStart(n *cc.ConstantExpression)
	// OnConstantExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.ConstantExpression is exited.
	OnConstantExpressionEnd(n *cc.ConstantExpression)

	// OnDeclarationStart is called during the traversal of the AST, whenever a node of type *cc.Declaration is entered.
	OnDeclarationStart(n *cc.Declaration)
	// OnDeclarationEnd is called during the traversal of the AST, whenever a node of type *cc.Declaration is exited.
	OnDeclarationEnd(n *cc.Declaration)

	// OnDeclarationListStart is called during the traversal of the AST, whenever a node of type *cc.DeclarationList is entered.
	OnDeclarationListStart(n *cc.DeclarationList)
	// OnDeclarationListEnd is called during the traversal of the AST, whenever a node of type *cc.DeclarationList is exited.
	OnDeclarationListEnd(n *cc.DeclarationList)

	// OnDeclarationSpecifiersStart is called during the traversal of the AST, whenever a node of type *cc.DeclarationSpecifiers is entered.
	OnDeclarationSpecifiersStart(n *cc.DeclarationSpecifiers)
	// OnDeclarationSpecifiersEnd is called during the traversal of the AST, whenever a node of type *cc.DeclarationSpecifiers is exited.
	OnDeclarationSpecifiersEnd(n *cc.DeclarationSpecifiers)

	// OnDeclaratorStart is called during the traversal of the AST, whenever a node of type *cc.Declarator is entered.
	OnDeclaratorStart(n *cc.Declarator)
	// OnDeclaratorEnd is called during the traversal of the AST, whenever a node of type *cc.Declarator is exited.
	OnDeclaratorEnd(n *cc.Declarator)

	// OnDesignationStart is called during the traversal of the AST, whenever a node of type *cc.Designation is entered.
	OnDesignationStart(n *cc.Designation)
	// OnDesignationEnd is called during the traversal of the AST, whenever a node of type *cc.Designation is exited.
	OnDesignationEnd(n *cc.Designation)

	// OnDesignatorStart is called during the traversal of the AST, whenever a node of type *cc.Designator is entered.
	OnDesignatorStart(n *cc.Designator)
	// OnDesignatorEnd is called during the traversal of the AST, whenever a node of type *cc.Designator is exited.
	OnDesignatorEnd(n *cc.Designator)

	// OnDesignatorListStart is called during the traversal of the AST, whenever a node of type *cc.DesignatorList is entered.
	OnDesignatorListStart(n *cc.DesignatorList)
	// OnDesignatorListEnd is called during the traversal of the AST, whenever a node of type *cc.DesignatorList is exited.
	OnDesignatorListEnd(n *cc.DesignatorList)

	// OnDirectAbstractDeclaratorStart is called during the traversal of the AST, whenever a node of type *cc.DirectAbstractDeclarator is entered.
	OnDirectAbstractDeclaratorStart(n *cc.DirectAbstractDeclarator)
	// OnDirectAbstractDeclaratorEnd is called during the traversal of the AST, whenever a node of type *cc.DirectAbstractDeclarator is exited.
	OnDirectAbstractDeclaratorEnd(n *cc.DirectAbstractDeclarator)

	// OnDirectDeclaratorStart is called during the traversal of the AST, whenever a node of type *cc.DirectDeclarator is entered.
	OnDirectDeclaratorStart(n *cc.DirectDeclarator)
	// OnDirectDeclaratorEnd is called during the traversal of the AST, whenever a node of type *cc.DirectDeclarator is exited.
	OnDirectDeclaratorEnd(n *cc.DirectDeclarator)

	// OnEnumSpecifierStart is called during the traversal of the AST, whenever a node of type *cc.EnumSpecifier is entered.
	OnEnumSpecifierStart(n *cc.EnumSpecifier)
	// OnEnumSpecifierEnd is called during the traversal of the AST, whenever a node of type *cc.EnumSpecifier is exited.
	OnEnumSpecifierEnd(n *cc.EnumSpecifier)

	// OnEnumeratorStart is called during the traversal of the AST, whenever a node of type *cc.Enumerator is entered.
	OnEnumeratorStart(n *cc.Enumerator)
	// OnEnumeratorEnd is called during the traversal of the AST, whenever a node of type *cc.Enumerator is exited.
	OnEnumeratorEnd(n *cc.Enumerator)

	// OnEnumeratorListStart is called during the traversal of the AST, whenever a node of type *cc.EnumeratorList is entered.
	OnEnumeratorListStart(n *cc.EnumeratorList)
	// OnEnumeratorListEnd is called during the traversal of the AST, whenever a node of type *cc.EnumeratorList is exited.
	OnEnumeratorListEnd(n *cc.EnumeratorList)

	// OnEqualityExpressionStart is called during the traversal of the AST, whenever a node of type *cc.EqualityExpression is entered.
	OnEqualityExpressionStart(n *cc.EqualityExpression)
	// OnEqualityExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.EqualityExpression is exited.
	OnEqualityExpressionEnd(n *cc.EqualityExpression)

	// OnExclusiveOrExpressionStart is called during the traversal of the AST, whenever a node of type *cc.ExclusiveOrExpression is entered.
	OnExclusiveOrExpressionStart(n *cc.ExclusiveOrExpression)
	// OnExclusiveOrExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.ExclusiveOrExpression is exited.
	OnExclusiveOrExpressionEnd(n *cc.ExclusiveOrExpression)

	// OnExpressionStart is called during the traversal of the AST, whenever a node of type *cc.Expression is entered.
	OnExpressionStart(n *cc.Expression)
	// OnExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.Expression is exited.
	OnExpressionEnd(n *cc.Expression)

	// OnExpressionListStart is called during the traversal of the AST, whenever a node of type *cc.ExpressionList is entered.
	OnExpressionListStart(n *cc.ExpressionList)
	// OnExpressionListEnd is called during the traversal of the AST, whenever a node of type *cc.ExpressionList is exited.
	OnExpressionListEnd(n *cc.ExpressionList)

	// OnExpressionStatementStart is called during the traversal of the AST, whenever a node of type *cc.ExpressionStatement is entered.
	OnExpressionStatementStart(n *cc.ExpressionStatement)
	// OnExpressionStatementEnd is called during the traversal of the AST, whenever a node of type *cc.ExpressionStatement is exited.
	OnExpressionStatementEnd(n *cc.ExpressionStatement)

	// OnExternalDeclarationStart is called during the traversal of the AST, whenever a node of type *cc.ExternalDeclaration is entered.
	OnExternalDeclarationStart(n *cc.ExternalDeclaration)
	// OnExternalDeclarationEnd is called during the traversal of the AST, whenever a node of type *cc.ExternalDeclaration is exited.
	OnExternalDeclarationEnd(n *cc.ExternalDeclaration)

	// OnFunctionDefinitionStart is called during the traversal of the AST, whenever a node of type *cc.FunctionDefinition is entered.
	OnFunctionDefinitionStart(n *cc.FunctionDefinition)
	// OnFunctionDefinitionEnd is called during the traversal of the AST, whenever a node of type *cc.FunctionDefinition is exited.
	OnFunctionDefinitionEnd(n *cc.FunctionDefinition)

	// OnFunctionSpecifierStart is called during the traversal of the AST, whenever a node of type *cc.FunctionSpecifier is entered.
	OnFunctionSpecifierStart(n *cc.FunctionSpecifier)
	// OnFunctionSpecifierEnd is called during the traversal of the AST, whenever a node of type *cc.FunctionSpecifier is exited.
	OnFunctionSpecifierEnd(n *cc.FunctionSpecifier)

	// OnIdentifierListStart is called during the traversal of the AST, whenever a node of type *cc.IdentifierList is entered.
	OnIdentifierListStart(n *cc.IdentifierList)
	// OnIdentifierListEnd is called during the traversal of the AST, whenever a node of type *cc.IdentifierList is exited.
	OnIdentifierListEnd(n *cc.IdentifierList)

	// OnInclusiveOrExpressionStart is called during the traversal of the AST, whenever a node of type *cc.InclusiveOrExpression is entered.
	OnInclusiveOrExpressionStart(n *cc.InclusiveOrExpression)
	// OnInclusiveOrExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.InclusiveOrExpression is exited.
	OnInclusiveOrExpressionEnd(n *cc.InclusiveOrExpression)

	// OnInitDeclaratorStart is called during the traversal of the AST, whenever a node of type *cc.InitDeclarator is entered.
	OnInitDeclaratorStart(n *cc.InitDeclarator)
	// OnInitDeclaratorEnd is called during the traversal of the AST, whenever a node of type *cc.InitDeclarator is exited.
	OnInitDeclaratorEnd(n *cc.InitDeclarator)

	// OnInitDeclaratorListStart is called during the traversal of the AST, whenever a node of type *cc.InitDeclaratorList is entered.
	OnInitDeclaratorListStart(n *cc.InitDeclaratorList)
	// OnInitDeclaratorListEnd is called during the traversal of the AST, whenever a node of type *cc.InitDeclaratorList is exited.
	OnInitDeclaratorListEnd(n *cc.InitDeclaratorList)

	// OnInitializerStart is called during the traversal of the AST, whenever a node of type *cc.Initializer is entered.
	OnInitializerStart(n *cc.Initializer)
	// OnInitializerEnd is called during the traversal of the AST, whenever a node of type *cc.Initializer is exited.
	OnInitializerEnd(n *cc.Initializer)

	// OnInitializerListStart is called during the traversal of the AST, whenever a node of type *cc.InitializerList is entered.
	OnInitializerListStart(n *cc.InitializerList)
	// OnInitializerListEnd is called during the traversal of the AST, whenever a node of type *cc.InitializerList is exited.
	OnInitializerListEnd(n *cc.InitializerList)

	// OnIterationStatementStart is called during the traversal of the AST, whenever a node of type *cc.IterationStatement is entered.
	OnIterationStatementStart(n *cc.IterationStatement)
	// OnIterationStatementEnd is called during the traversal of the AST, whenever a node of type *cc.IterationStatement is exited.
	OnIterationStatementEnd(n *cc.IterationStatement)

	// OnJumpStatementStart is called during the traversal of the AST, whenever a node of type *cc.JumpStatement is entered.
	OnJumpStatementStart(n *cc.JumpStatement)
	// OnJumpStatementEnd is called during the traversal of the AST, whenever a node of type *cc.JumpStatement is exited.
	OnJumpStatementEnd(n *cc.JumpStatement)

	// OnLabelDeclarationStart is called during the traversal of the AST, whenever a node of type *cc.LabelDeclaration is entered.
	OnLabelDeclarationStart(n *cc.LabelDeclaration)
	// OnLabelDeclarationEnd is called during the traversal of the AST, whenever a node of type *cc.LabelDeclaration is exited.
	OnLabelDeclarationEnd(n *cc.LabelDeclaration)

	// OnLabeledStatementStart is called during the traversal of the AST, whenever a node of type *cc.LabeledStatement is entered.
	OnLabeledStatementStart(n *cc.LabeledStatement)
	// OnLabeledStatementEnd is called during the traversal of the AST, whenever a node of type *cc.LabeledStatement is exited.
	OnLabeledStatementEnd(n *cc.LabeledStatement)

	// OnLogicalAndExpressionStart is called during the traversal of the AST, whenever a node of type *cc.LogicalAndExpression is entered.
	OnLogicalAndExpressionStart(n *cc.LogicalAndExpression)
	// OnLogicalAndExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.LogicalAndExpression is exited.
	OnLogicalAndExpressionEnd(n *cc.LogicalAndExpression)

	// OnLogicalOrExpressionStart is called during the traversal of the AST, whenever a node of type *cc.LogicalOrExpression is entered.
	OnLogicalOrExpressionStart(n *cc.LogicalOrExpression)
	// OnLogicalOrExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.LogicalOrExpression is exited.
	OnLogicalOrExpressionEnd(n *cc.LogicalOrExpression)

	// OnMultiplicativeExpressionStart is called during the traversal of the AST, whenever a node of type *cc.MultiplicativeExpression is entered.
	OnMultiplicativeExpressionStart(n *cc.MultiplicativeExpression)
	// OnMultiplicativeExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.MultiplicativeExpression is exited.
	OnMultiplicativeExpressionEnd(n *cc.MultiplicativeExpression)

	// OnParameterDeclarationStart is called during the traversal of the AST, whenever a node of type *cc.ParameterDeclaration is entered.
	OnParameterDeclarationStart(n *cc.ParameterDeclaration)
	// OnParameterDeclarationEnd is called during the traversal of the AST, whenever a node of type *cc.ParameterDeclaration is exited.
	OnParameterDeclarationEnd(n *cc.ParameterDeclaration)

	// OnParameterListStart is called during the traversal of the AST, whenever a node of type *cc.ParameterList is entered.
	OnParameterListStart(n *cc.ParameterList)
	// OnParameterListEnd is called during the traversal of the AST, whenever a node of type *cc.ParameterList is exited.
	OnParameterListEnd(n *cc.ParameterList)

	// OnParameterTypeListStart is called during the traversal of the AST, whenever a node of type *cc.ParameterTypeList is entered.
	OnParameterTypeListStart(n *cc.ParameterTypeList)
	// OnParameterTypeListEnd is called during the traversal of the AST, whenever a node of type *cc.ParameterTypeList is exited.
	OnParameterTypeListEnd(n *cc.ParameterTypeList)

	// OnPointerStart is called during the traversal of the AST, whenever a node of type *cc.Pointer is entered.
	OnPointerStart(n *cc.Pointer)
	// OnPointerEnd is called during the traversal of the AST, whenever a node of type *cc.Pointer is exited.
	OnPointerEnd(n *cc.Pointer)

	// OnPostfixExpressionStart is called during the traversal of the AST, whenever a node of type *cc.PostfixExpression is entered.
	OnPostfixExpressionStart(n *cc.PostfixExpression)
	// OnPostfixExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.PostfixExpression is exited.
	OnPostfixExpressionEnd(n *cc.PostfixExpression)

	// OnPragmaSTDCStart is called during the traversal of the AST, whenever a node of type *cc.PragmaSTDC is entered.
	OnPragmaSTDCStart(n *cc.PragmaSTDC)
	// OnPragmaSTDCEnd is called during the traversal of the AST, whenever a node of type *cc.PragmaSTDC is exited.
	OnPragmaSTDCEnd(n *cc.PragmaSTDC)

	// OnPrimaryExpressionStart is called during the traversal of the AST, whenever a node of type *cc.PrimaryExpression is entered.
	OnPrimaryExpressionStart(n *cc.PrimaryExpression)
	// OnPrimaryExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.PrimaryExpression is exited.
	OnPrimaryExpressionEnd(n *cc.PrimaryExpression)

	// OnRelationalExpressionStart is called during the traversal of the AST, whenever a node of type *cc.RelationalExpression is entered.
	OnRelationalExpressionStart(n *cc.RelationalExpression)
	// OnRelationalExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.RelationalExpression is exited.
	OnRelationalExpressionEnd(n *cc.RelationalExpression)

	// OnSelectionStatementStart is called during the traversal of the AST, whenever a node of type *cc.SelectionStatement is entered.
	OnSelectionStatementStart(n *cc.SelectionStatement)
	// OnSelectionStatementEnd is called during the traversal of the AST, whenever a node of type *cc.SelectionStatement is exited.
	OnSelectionStatementEnd(n *cc.SelectionStatement)

	// OnShiftExpressionStart is called during the traversal of the AST, whenever a node of type *cc.ShiftExpression is entered.
	OnShiftExpressionStart(n *cc.ShiftExpression)
	// OnShiftExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.ShiftExpression is exited.
	OnShiftExpressionEnd(n *cc.ShiftExpression)

	// OnSpecifierQualifierListStart is called during the traversal of the AST, whenever a node of type *cc.SpecifierQualifierList is entered.
	OnSpecifierQualifierListStart(n *cc.SpecifierQualifierList)
	// OnSpecifierQualifierListEnd is called during the traversal of the AST, whenever a node of type *cc.SpecifierQualifierList is exited.
	OnSpecifierQualifierListEnd(n *cc.SpecifierQualifierList)

	// OnStatementStart is called during the traversal of the AST, whenever a node of type *cc.Statement is entered.
	OnStatementStart(n *cc.Statement)
	// OnStatementEnd is called during the traversal of the AST, whenever a node of type *cc.Statement is exited.
	OnStatementEnd(n *cc.Statement)

	// OnStorageClassSpecifierStart is called during the traversal of the AST, whenever a node of type *cc.StorageClassSpecifier is entered.
	OnStorageClassSpecifierStart(n *cc.StorageClassSpecifier)
	// OnStorageClassSpecifierEnd is called during the traversal of the AST, whenever a node of type *cc.StorageClassSpecifier is exited.
	OnStorageClassSpecifierEnd(n *cc.StorageClassSpecifier)

	// OnStructDeclarationStart is called during the traversal of the AST, whenever a node of type *cc.StructDeclaration is entered.
	OnStructDeclarationStart(n *cc.StructDeclaration)
	// OnStructDeclarationEnd is called during the traversal of the AST, whenever a node of type *cc.StructDeclaration is exited.
	OnStructDeclarationEnd(n *cc.StructDeclaration)

	// OnStructDeclarationListStart is called during the traversal of the AST, whenever a node of type *cc.StructDeclarationList is entered.
	OnStructDeclarationListStart(n *cc.StructDeclarationList)
	// OnStructDeclarationListEnd is called during the traversal of the AST, whenever a node of type *cc.StructDeclarationList is exited.
	OnStructDeclarationListEnd(n *cc.StructDeclarationList)

	// OnStructDeclaratorStart is called during the traversal of the AST, whenever a node of type *cc.StructDeclarator is entered.
	OnStructDeclaratorStart(n *cc.StructDeclarator)
	// OnStructDeclaratorEnd is called during the traversal of the AST, whenever a node of type *cc.StructDeclarator is exited.
	OnStructDeclaratorEnd(n *cc.StructDeclarator)

	// OnStructDeclaratorListStart is called during the traversal of the AST, whenever a node of type *cc.StructDeclaratorList is entered.
	OnStructDeclaratorListStart(n *cc.StructDeclaratorList)
	// OnStructDeclaratorListEnd is called during the traversal of the AST, whenever a node of type *cc.StructDeclaratorList is exited.
	OnStructDeclaratorListEnd(n *cc.StructDeclaratorList)

	// OnStructOrUnionStart is called during the traversal of the AST, whenever a node of type *cc.StructOrUnion is entered.
	OnStructOrUnionStart(n *cc.StructOrUnion)
	// OnStructOrUnionEnd is called during the traversal of the AST, whenever a node of type *cc.StructOrUnion is exited.
	OnStructOrUnionEnd(n *cc.StructOrUnion)

	// OnStructOrUnionSpecifierStart is called during the traversal of the AST, whenever a node of type *cc.StructOrUnionSpecifier is entered.
	OnStructOrUnionSpecifierStart(n *cc.StructOrUnionSpecifier)
	// OnStructOrUnionSpecifierEnd is called during the traversal of the AST, whenever a node of type *cc.StructOrUnionSpecifier is exited.
	OnStructOrUnionSpecifierEnd(n *cc.StructOrUnionSpecifier)

	// OnTranslationUnitStart is called during the traversal of the AST, whenever a node of type *cc.TranslationUnit is entered.
	OnTranslationUnitStart(n *cc.TranslationUnit)
	// OnTranslationUnitEnd is called during the traversal of the AST, whenever a node of type *cc.TranslationUnit is exited.
	OnTranslationUnitEnd(n *cc.TranslationUnit)

	// OnTypeNameStart is called during the traversal of the AST, whenever a node of type *cc.TypeName is entered.
	OnTypeNameStart(n *cc.TypeName)
	// OnTypeNameEnd is called during the traversal of the AST, whenever a node of type *cc.TypeName is exited.
	OnTypeNameEnd(n *cc.TypeName)

	// OnTypeQualifierStart is called during the traversal of the AST, whenever a node of type *cc.TypeQualifier is entered.
	OnTypeQualifierStart(n *cc.TypeQualifier)
	// OnTypeQualifierEnd is called during the traversal of the AST, whenever a node of type *cc.TypeQualifier is exited.
	OnTypeQualifierEnd(n *cc.TypeQualifier)

	// OnTypeQualifiersStart is called during the traversal of the AST, whenever a node of type *cc.TypeQualifiers is entered.
	OnTypeQualifiersStart(n *cc.TypeQualifiers)
	// OnTypeQualifiersEnd is called during the traversal of the AST, whenever a node of type *cc.TypeQualifiers is exited.
	OnTypeQualifiersEnd(n *cc.TypeQualifiers)

	// OnTypeSpecifierStart is called during the traversal of the AST, whenever a node of type *cc.TypeSpecifier is entered.
	OnTypeSpecifierStart(n *cc.TypeSpecifier)
	// OnTypeSpecifierEnd is called during the traversal of the AST, whenever a node of type *cc.TypeSpecifier is exited.
	OnTypeSpecifierEnd(n *cc.TypeSpecifier)

	// OnUnaryExpressionStart is called during the traversal of the AST, whenever a node of type *cc.UnaryExpression is entered.
	OnUnaryExpressionStart(n *cc.UnaryExpression)
	// OnUnaryExpressionEnd is called during the traversal of the AST, whenever a node of type *cc.UnaryExpression is exited.
	OnUnaryExpressionEnd(n *cc.UnaryExpression)
}

func BaseHandler() Handler {
	return &baseHandler{}
}

type baseHandler struct{}

func (h *baseHandler) Start() {}
func (h *baseHandler) End()   {}

func (h *baseHandler) OnAbstractDeclaratorStart(*cc.AbstractDeclarator) {}
func (h *baseHandler) OnAbstractDeclaratorEnd(*cc.AbstractDeclarator)   {}

func (h *baseHandler) OnAdditiveExpressionStart(*cc.AdditiveExpression) {}
func (h *baseHandler) OnAdditiveExpressionEnd(*cc.AdditiveExpression)   {}

func (h *baseHandler) OnAlignmentSpecifierStart(*cc.AlignmentSpecifier) {}
func (h *baseHandler) OnAlignmentSpecifierEnd(*cc.AlignmentSpecifier)   {}

func (h *baseHandler) OnAndExpressionStart(*cc.AndExpression) {}
func (h *baseHandler) OnAndExpressionEnd(*cc.AndExpression)   {}

func (h *baseHandler) OnArgumentExpressionListStart(*cc.ArgumentExpressionList) {}
func (h *baseHandler) OnArgumentExpressionListEnd(*cc.ArgumentExpressionList)   {}

func (h *baseHandler) OnAsmStart(*cc.Asm) {}
func (h *baseHandler) OnAsmEnd(*cc.Asm)   {}

func (h *baseHandler) OnAsmArgListStart(*cc.AsmArgList) {}
func (h *baseHandler) OnAsmArgListEnd(*cc.AsmArgList)   {}

func (h *baseHandler) OnAsmExpressionListStart(*cc.AsmExpressionList) {}
func (h *baseHandler) OnAsmExpressionListEnd(*cc.AsmExpressionList)   {}

func (h *baseHandler) OnAsmFunctionDefinitionStart(*cc.AsmFunctionDefinition) {}
func (h *baseHandler) OnAsmFunctionDefinitionEnd(*cc.AsmFunctionDefinition)   {}

func (h *baseHandler) OnAsmIndexStart(*cc.AsmIndex) {}
func (h *baseHandler) OnAsmIndexEnd(*cc.AsmIndex)   {}

func (h *baseHandler) OnAsmQualifierStart(*cc.AsmQualifier) {}
func (h *baseHandler) OnAsmQualifierEnd(*cc.AsmQualifier)   {}

func (h *baseHandler) OnAsmQualifierListStart(*cc.AsmQualifierList) {}
func (h *baseHandler) OnAsmQualifierListEnd(*cc.AsmQualifierList)   {}

func (h *baseHandler) OnAsmStatementStart(*cc.AsmStatement) {}
func (h *baseHandler) OnAsmStatementEnd(*cc.AsmStatement)   {}

func (h *baseHandler) OnAssignmentExpressionStart(*cc.AssignmentExpression) {}
func (h *baseHandler) OnAssignmentExpressionEnd(*cc.AssignmentExpression)   {}

func (h *baseHandler) OnAtomicTypeSpecifierStart(*cc.AtomicTypeSpecifier) {}
func (h *baseHandler) OnAtomicTypeSpecifierEnd(*cc.AtomicTypeSpecifier)   {}

func (h *baseHandler) OnAttributeSpecifierStart(*cc.AttributeSpecifier) {}
func (h *baseHandler) OnAttributeSpecifierEnd(*cc.AttributeSpecifier)   {}

func (h *baseHandler) OnAttributeSpecifierListStart(*cc.AttributeSpecifierList) {}
func (h *baseHandler) OnAttributeSpecifierListEnd(*cc.AttributeSpecifierList)   {}

func (h *baseHandler) OnAttributeValueStart(*cc.AttributeValue) {}
func (h *baseHandler) OnAttributeValueEnd(*cc.AttributeValue)   {}

func (h *baseHandler) OnAttributeValueListStart(*cc.AttributeValueList) {}
func (h *baseHandler) OnAttributeValueListEnd(*cc.AttributeValueList)   {}

func (h *baseHandler) OnBlockItemStart(*cc.BlockItem) {}
func (h *baseHandler) OnBlockItemEnd(*cc.BlockItem)   {}

func (h *baseHandler) OnBlockItemListStart(*cc.BlockItemList) {}
func (h *baseHandler) OnBlockItemListEnd(*cc.BlockItemList)   {}

func (h *baseHandler) OnCastExpressionStart(*cc.CastExpression) {}
func (h *baseHandler) OnCastExpressionEnd(*cc.CastExpression)   {}

func (h *baseHandler) OnCompoundStatementStart(*cc.CompoundStatement) {}
func (h *baseHandler) OnCompoundStatementEnd(*cc.CompoundStatement)   {}

func (h *baseHandler) OnConditionalExpressionStart(*cc.ConditionalExpression) {}
func (h *baseHandler) OnConditionalExpressionEnd(*cc.ConditionalExpression)   {}

func (h *baseHandler) OnConstantExpressionStart(*cc.ConstantExpression) {}
func (h *baseHandler) OnConstantExpressionEnd(*cc.ConstantExpression)   {}

func (h *baseHandler) OnDeclarationStart(*cc.Declaration) {}
func (h *baseHandler) OnDeclarationEnd(*cc.Declaration)   {}

func (h *baseHandler) OnDeclarationListStart(*cc.DeclarationList) {}
func (h *baseHandler) OnDeclarationListEnd(*cc.DeclarationList)   {}

func (h *baseHandler) OnDeclarationSpecifiersStart(*cc.DeclarationSpecifiers) {}
func (h *baseHandler) OnDeclarationSpecifiersEnd(*cc.DeclarationSpecifiers)   {}

func (h *baseHandler) OnDeclaratorStart(*cc.Declarator) {}
func (h *baseHandler) OnDeclaratorEnd(*cc.Declarator)   {}

func (h *baseHandler) OnDesignationStart(*cc.Designation) {}
func (h *baseHandler) OnDesignationEnd(*cc.Designation)   {}

func (h *baseHandler) OnDesignatorStart(*cc.Designator) {}
func (h *baseHandler) OnDesignatorEnd(*cc.Designator)   {}

func (h *baseHandler) OnDesignatorListStart(*cc.DesignatorList) {}
func (h *baseHandler) OnDesignatorListEnd(*cc.DesignatorList)   {}

func (h *baseHandler) OnDirectAbstractDeclaratorStart(*cc.DirectAbstractDeclarator) {}
func (h *baseHandler) OnDirectAbstractDeclaratorEnd(*cc.DirectAbstractDeclarator)   {}

func (h *baseHandler) OnDirectDeclaratorStart(*cc.DirectDeclarator) {}
func (h *baseHandler) OnDirectDeclaratorEnd(*cc.DirectDeclarator)   {}

func (h *baseHandler) OnEnumSpecifierStart(*cc.EnumSpecifier) {}
func (h *baseHandler) OnEnumSpecifierEnd(*cc.EnumSpecifier)   {}

func (h *baseHandler) OnEnumeratorStart(*cc.Enumerator) {}
func (h *baseHandler) OnEnumeratorEnd(*cc.Enumerator)   {}

func (h *baseHandler) OnEnumeratorListStart(*cc.EnumeratorList) {}
func (h *baseHandler) OnEnumeratorListEnd(*cc.EnumeratorList)   {}

func (h *baseHandler) OnEqualityExpressionStart(*cc.EqualityExpression) {}
func (h *baseHandler) OnEqualityExpressionEnd(*cc.EqualityExpression)   {}

func (h *baseHandler) OnExclusiveOrExpressionStart(*cc.ExclusiveOrExpression) {}
func (h *baseHandler) OnExclusiveOrExpressionEnd(*cc.ExclusiveOrExpression)   {}

func (h *baseHandler) OnExpressionStart(*cc.Expression) {}
func (h *baseHandler) OnExpressionEnd(*cc.Expression)   {}

func (h *baseHandler) OnExpressionListStart(*cc.ExpressionList) {}
func (h *baseHandler) OnExpressionListEnd(*cc.ExpressionList)   {}

func (h *baseHandler) OnExpressionStatementStart(*cc.ExpressionStatement) {}
func (h *baseHandler) OnExpressionStatementEnd(*cc.ExpressionStatement)   {}

func (h *baseHandler) OnExternalDeclarationStart(*cc.ExternalDeclaration) {}
func (h *baseHandler) OnExternalDeclarationEnd(*cc.ExternalDeclaration)   {}

func (h *baseHandler) OnFunctionDefinitionStart(*cc.FunctionDefinition) {}
func (h *baseHandler) OnFunctionDefinitionEnd(*cc.FunctionDefinition)   {}

func (h *baseHandler) OnFunctionSpecifierStart(*cc.FunctionSpecifier) {}
func (h *baseHandler) OnFunctionSpecifierEnd(*cc.FunctionSpecifier)   {}

func (h *baseHandler) OnIdentifierListStart(*cc.IdentifierList) {}
func (h *baseHandler) OnIdentifierListEnd(*cc.IdentifierList)   {}

func (h *baseHandler) OnInclusiveOrExpressionStart(*cc.InclusiveOrExpression) {}
func (h *baseHandler) OnInclusiveOrExpressionEnd(*cc.InclusiveOrExpression)   {}

func (h *baseHandler) OnInitDeclaratorStart(*cc.InitDeclarator) {}
func (h *baseHandler) OnInitDeclaratorEnd(*cc.InitDeclarator)   {}

func (h *baseHandler) OnInitDeclaratorListStart(*cc.InitDeclaratorList) {}
func (h *baseHandler) OnInitDeclaratorListEnd(*cc.InitDeclaratorList)   {}

func (h *baseHandler) OnInitializerStart(*cc.Initializer) {}
func (h *baseHandler) OnInitializerEnd(*cc.Initializer)   {}

func (h *baseHandler) OnInitializerListStart(*cc.InitializerList) {}
func (h *baseHandler) OnInitializerListEnd(*cc.InitializerList)   {}

func (h *baseHandler) OnIterationStatementStart(*cc.IterationStatement) {}
func (h *baseHandler) OnIterationStatementEnd(*cc.IterationStatement)   {}

func (h *baseHandler) OnJumpStatementStart(*cc.JumpStatement) {}
func (h *baseHandler) OnJumpStatementEnd(*cc.JumpStatement)   {}

func (h *baseHandler) OnLabelDeclarationStart(*cc.LabelDeclaration) {}
func (h *baseHandler) OnLabelDeclarationEnd(*cc.LabelDeclaration)   {}

func (h *baseHandler) OnLabeledStatementStart(*cc.LabeledStatement) {}
func (h *baseHandler) OnLabeledStatementEnd(*cc.LabeledStatement)   {}

func (h *baseHandler) OnLogicalAndExpressionStart(*cc.LogicalAndExpression) {}
func (h *baseHandler) OnLogicalAndExpressionEnd(*cc.LogicalAndExpression)   {}

func (h *baseHandler) OnLogicalOrExpressionStart(*cc.LogicalOrExpression) {}
func (h *baseHandler) OnLogicalOrExpressionEnd(*cc.LogicalOrExpression)   {}

func (h *baseHandler) OnMultiplicativeExpressionStart(*cc.MultiplicativeExpression) {}
func (h *baseHandler) OnMultiplicativeExpressionEnd(*cc.MultiplicativeExpression)   {}

func (h *baseHandler) OnParameterDeclarationStart(*cc.ParameterDeclaration) {}
func (h *baseHandler) OnParameterDeclarationEnd(*cc.ParameterDeclaration)   {}

func (h *baseHandler) OnParameterListStart(*cc.ParameterList) {}
func (h *baseHandler) OnParameterListEnd(*cc.ParameterList)   {}

func (h *baseHandler) OnParameterTypeListStart(*cc.ParameterTypeList) {}
func (h *baseHandler) OnParameterTypeListEnd(*cc.ParameterTypeList)   {}

func (h *baseHandler) OnPointerStart(*cc.Pointer) {}
func (h *baseHandler) OnPointerEnd(*cc.Pointer)   {}

func (h *baseHandler) OnPostfixExpressionStart(*cc.PostfixExpression) {}
func (h *baseHandler) OnPostfixExpressionEnd(*cc.PostfixExpression)   {}

func (h *baseHandler) OnPragmaSTDCStart(*cc.PragmaSTDC) {}
func (h *baseHandler) OnPragmaSTDCEnd(*cc.PragmaSTDC)   {}

func (h *baseHandler) OnPrimaryExpressionStart(*cc.PrimaryExpression) {}
func (h *baseHandler) OnPrimaryExpressionEnd(*cc.PrimaryExpression)   {}

func (h *baseHandler) OnRelationalExpressionStart(*cc.RelationalExpression) {}
func (h *baseHandler) OnRelationalExpressionEnd(*cc.RelationalExpression)   {}

func (h *baseHandler) OnSelectionStatementStart(*cc.SelectionStatement) {}
func (h *baseHandler) OnSelectionStatementEnd(*cc.SelectionStatement)   {}

func (h *baseHandler) OnShiftExpressionStart(*cc.ShiftExpression) {}
func (h *baseHandler) OnShiftExpressionEnd(*cc.ShiftExpression)   {}

func (h *baseHandler) OnSpecifierQualifierListStart(*cc.SpecifierQualifierList) {}
func (h *baseHandler) OnSpecifierQualifierListEnd(*cc.SpecifierQualifierList)   {}

func (h *baseHandler) OnStatementStart(*cc.Statement) {}
func (h *baseHandler) OnStatementEnd(*cc.Statement)   {}

func (h *baseHandler) OnStorageClassSpecifierStart(*cc.StorageClassSpecifier) {}
func (h *baseHandler) OnStorageClassSpecifierEnd(*cc.StorageClassSpecifier)   {}

func (h *baseHandler) OnStructDeclarationStart(*cc.StructDeclaration) {}
func (h *baseHandler) OnStructDeclarationEnd(*cc.StructDeclaration)   {}

func (h *baseHandler) OnStructDeclarationListStart(*cc.StructDeclarationList) {}
func (h *baseHandler) OnStructDeclarationListEnd(*cc.StructDeclarationList)   {}

func (h *baseHandler) OnStructDeclaratorStart(*cc.StructDeclarator) {}
func (h *baseHandler) OnStructDeclaratorEnd(*cc.StructDeclarator)   {}

func (h *baseHandler) OnStructDeclaratorListStart(*cc.StructDeclaratorList) {}
func (h *baseHandler) OnStructDeclaratorListEnd(*cc.StructDeclaratorList)   {}

func (h *baseHandler) OnStructOrUnionStart(*cc.StructOrUnion) {}
func (h *baseHandler) OnStructOrUnionEnd(*cc.StructOrUnion)   {}

func (h *baseHandler) OnStructOrUnionSpecifierStart(*cc.StructOrUnionSpecifier) {}
func (h *baseHandler) OnStructOrUnionSpecifierEnd(*cc.StructOrUnionSpecifier)   {}

func (h *baseHandler) OnTranslationUnitStart(*cc.TranslationUnit) {}
func (h *baseHandler) OnTranslationUnitEnd(*cc.TranslationUnit)   {}

func (h *baseHandler) OnTypeNameStart(*cc.TypeName) {}
func (h *baseHandler) OnTypeNameEnd(*cc.TypeName)   {}

func (h *baseHandler) OnTypeQualifierStart(*cc.TypeQualifier) {}
func (h *baseHandler) OnTypeQualifierEnd(*cc.TypeQualifier)   {}

func (h *baseHandler) OnTypeQualifiersStart(*cc.TypeQualifiers) {}
func (h *baseHandler) OnTypeQualifiersEnd(*cc.TypeQualifiers)   {}

func (h *baseHandler) OnTypeSpecifierStart(*cc.TypeSpecifier) {}
func (h *baseHandler) OnTypeSpecifierEnd(*cc.TypeSpecifier)   {}

func (h *baseHandler) OnUnaryExpressionStart(*cc.UnaryExpression) {}
func (h *baseHandler) OnUnaryExpressionEnd(*cc.UnaryExpression)   {}
