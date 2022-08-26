// Code generated by "stringer -type Level -linecomment"; DO NOT EDIT.

package log

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LevelInfo-1]
	_ = x[LevelDebug-2]
	_ = x[LevelWarn-3]
	_ = x[LevelError-4]
	_ = x[LevelFatal-5]
}

const _Level_name = "INFODEBUGWARNERRORFATAL"

var _Level_index = [...]uint8{0, 4, 9, 13, 18, 23}

func (i Level) String() string {
	i -= 1
	if i < 0 || i >= Level(len(_Level_index)-1) {
		return "Level(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Level_name[_Level_index[i]:_Level_index[i+1]]
}