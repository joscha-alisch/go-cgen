package ast

import (
	"strconv"

	"modernc.org/cc/v3"
)

func FindValueConstantExpression(e *cc.ConstantExpression) Value {
	if e == nil {
		return Value{}
	}
	return findValueConditionalExpression(e.ConditionalExpression)
}

func findValueConditionalExpression(e *cc.ConditionalExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueLogicalOrExpression(e.LogicalOrExpression),
		findValueExpression(e.Expression),
		findValueConditionalExpression(e.ConditionalExpression),
	)
}

func findValueLogicalOrExpression(e *cc.LogicalOrExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueLogicalAndExpression(e.LogicalAndExpression),
		findValueLogicalOrExpression(e.LogicalOrExpression),
	)
}

func findValueLogicalAndExpression(e *cc.LogicalAndExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueInclusiveOrExpression(e.InclusiveOrExpression),
		findValueLogicalAndExpression(e.LogicalAndExpression),
	)
}

func findValueInclusiveOrExpression(e *cc.InclusiveOrExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueExclusiveOrExpression(e.ExclusiveOrExpression),
		findValueInclusiveOrExpression(e.InclusiveOrExpression),
	)
}

func findValueExclusiveOrExpression(e *cc.ExclusiveOrExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueAndExpression(e.AndExpression),
		findValueExclusiveOrExpression(e.ExclusiveOrExpression),
	)
}

func findValueAndExpression(e *cc.AndExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueEqualityExpression(e.EqualityExpression),
		findValueAndExpression(e.AndExpression),
	)
}

func findValueEqualityExpression(e *cc.EqualityExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueRelationalExpression(e.RelationalExpression),
		findValueEqualityExpression(e.EqualityExpression),
	)
}

func findValueRelationalExpression(e *cc.RelationalExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueShiftExpression(e.ShiftExpression),
		findValueRelationalExpression(e.RelationalExpression),
	)
}

func findValueShiftExpression(e *cc.ShiftExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueAdditiveExpression(e.AdditiveExpression),
		findValueShiftExpression(e.ShiftExpression),
	)
}

func findValueAdditiveExpression(e *cc.AdditiveExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueMultiplicativeExpression(e.MultiplicativeExpression),
		findValueAdditiveExpression(e.AdditiveExpression),
	)
}

func findValueMultiplicativeExpression(e *cc.MultiplicativeExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueCastExpression(e.CastExpression),
		findValueMultiplicativeExpression(e.MultiplicativeExpression),
	)
}

func findValueCastExpression(e *cc.CastExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueUnaryExpression(e.UnaryExpression),
		findValueCastExpression(e.CastExpression),
	)
}

func findValueUnaryExpression(e *cc.UnaryExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValuePostfixExpression(e.PostfixExpression),
		findValueCastExpression(e.CastExpression),
		findValueUnaryExpression(e.UnaryExpression),
	)
}

func findValuePostfixExpression(e *cc.PostfixExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValuePrimaryExpression(e.PrimaryExpression),
		findValueAssignmentExpression(e.AssignmentExpression),
		findValueExpression(e.Expression),
		findValuePostfixExpression(e.PostfixExpression),
	)
}

func findValueAssignmentExpression(e *cc.AssignmentExpression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueConditionalExpression(e.ConditionalExpression),
		findValueUnaryExpression(e.UnaryExpression),
		findValueAssignmentExpression(e.AssignmentExpression),
	)
}

func findValueExpression(e *cc.Expression) Value {
	if e == nil {
		return Value{}
	}
	return anyOf(
		findValueAssignmentExpression(e.AssignmentExpression),
		findValueExpression(e.Expression),
	)
}

func findValuePrimaryExpression(e *cc.PrimaryExpression) Value {
	if e == nil {
		return Value{}
	}

	switch e.Case {
	case cc.PrimaryExpressionEnum:
		return Value{
			Type:  ValueTypeIdentifier,
			Value: e.Token.String(),
		}
	case cc.PrimaryExpressionInt:
		n, err := readNumeric([]rune(e.Token.String()))
		if err != nil {
			println("error converting " + e.Token.String() + "to number")
			return Value{}
		}
		return Value{
			Type:  ValueTypeConstant,
			Value: n,
		}
	}

	return anyOf(
		findValueExpression(e.Expression),
	)
}

func anyOf(possibilities ...Value) Value {
	for _, possibility := range possibilities {
		if possibility.Type != ValueTypeUnsupported {
			return possibility
		}
	}
	return Value{}
}

func readNumeric(v []rune) (interface{}, error) {
	long := false
	unsigned := false
	result := make([]rune, 0, len(v))
	isHex := false

	wrap := func() (interface{}, error) {
		res := string(result)

		if unsigned {
			n, err := strconv.ParseUint(res, 0, 64)
			if long {
				return n, err
			} else {
				return uint(n), err
			}
		} else {
			n, err := strconv.ParseInt(res, 0, 64)
			if long {
				return n, err
			} else {
				return int(n), err
			}
		}
	}

	for _, r := range v {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			result = append(result, r)
		case 'x':
			isHex = true
			result = append(result, r)
		case '-', '+', '.', 'e', 'b':
			result = append(result, r)
		case 'A', 'B', 'C', 'D', 'E', 'F', 'a', 'c', 'd', 'f':
			if isHex {
				result = append(result, r)
			}
		case 'L', 'l':
			long = true
			continue
		case 'U', 'u':
			unsigned = true
			continue
		default:
			return wrap()
		}
	}
	return wrap()
}
