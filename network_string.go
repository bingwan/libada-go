// Code generated by "stringer -type=Network"; DO NOT EDIT.

package libada

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Testnet-0]
	_ = x[Mainnet-1]
}

const _Network_name = "TestnetMainnet"

var _Network_index = [...]uint8{0, 7, 14}

func (i Network) String() string {
	if i >= Network(len(_Network_index)-1) {
		return "Network(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Network_name[_Network_index[i]:_Network_index[i+1]]
}
