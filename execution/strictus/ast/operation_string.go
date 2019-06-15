// Code generated by "stringer -type=Operation"; DO NOT EDIT.

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OperationOr-0]
	_ = x[OperationAnd-1]
	_ = x[OperationEqual-2]
	_ = x[OperationUnequal-3]
	_ = x[OperationLess-4]
	_ = x[OperationGreater-5]
	_ = x[OperationLessEqual-6]
	_ = x[OperationGreaterEqual-7]
	_ = x[OperationPlus-8]
	_ = x[OperationMinus-9]
	_ = x[OperationMul-10]
	_ = x[OperationDiv-11]
	_ = x[OperationMod-12]
}

const _Operation_name = "OperationOrOperationAndOperationEqualOperationUnequalOperationLessOperationGreaterOperationLessEqualOperationGreaterEqualOperationPlusOperationMinusOperationMulOperationDivOperationMod"

var _Operation_index = [...]uint8{0, 11, 23, 37, 53, 66, 82, 100, 121, 134, 148, 160, 172, 184}

func (i Operation) String() string {
	if i < 0 || i >= Operation(len(_Operation_index)-1) {
		return "Operation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Operation_name[_Operation_index[i]:_Operation_index[i+1]]
}
