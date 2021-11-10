package ast

import "modernc.org/cc/v3"

func EachInitDeclarator(n *cc.InitDeclaratorList, f func(n *cc.InitDeclarator)) {
	current := n
	for current != nil {
		if current.InitDeclarator != nil {
			f(current.InitDeclarator)
		}
		current = current.InitDeclaratorList
	}
}

func EachDeclarationSpecifier(n *cc.DeclarationSpecifiers, f func(n *cc.DeclarationSpecifiers)) {
	current := n
	for current != nil {
		f(current)
		current = current.DeclarationSpecifiers
	}
}

func EachEnumerator(n *cc.EnumeratorList, f func(n *cc.Enumerator)) {
	current := n
	for current != nil {
		if current.Enumerator != nil {
			f(current.Enumerator)
		}
		current = current.EnumeratorList
	}
}
