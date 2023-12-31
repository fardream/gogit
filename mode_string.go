// Code generated by "stringer -type=Mode -linecomment"; DO NOT EDIT.

package gogit

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Mode_RegularFile-100644]
	_ = x[Mode_Tree-40000]
}

const (
	_Mode_name_0 = "40000"
	_Mode_name_1 = "100644"
)

func (i Mode) String() string {
	switch {
	case i == 40000:
		return _Mode_name_0
	case i == 100644:
		return _Mode_name_1
	default:
		return "Mode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
