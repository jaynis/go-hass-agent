// Code generated by "stringer -type=batteryType -output batteryTypeStrings.go -linecomment"; DO NOT EDIT.

package linux

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[batteryTypeLinePower-1]
	_ = x[batteryTypeBattery-2]
	_ = x[batteryTypeUps-3]
	_ = x[batteryTypeMonitor-4]
	_ = x[batteryTypeMouse-5]
	_ = x[batteryTypeKeyboard-6]
	_ = x[batteryTypePda-7]
	_ = x[batteryTypePhone-8]
}

const _batteryType_name = "Line PowerBatteryUPSMonitorMouseKeyboardPdaPhone"

var _batteryType_index = [...]uint8{0, 10, 17, 20, 27, 32, 40, 43, 48}

func (i batteryType) String() string {
	i -= 1
	if i >= batteryType(len(_batteryType_index)-1) {
		return "batteryType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _batteryType_name[_batteryType_index[i]:_batteryType_index[i+1]]
}
