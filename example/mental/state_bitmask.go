// Code generated by "bitmasker -type=State"; DO NOT EDIT.

package mental

import "strconv"

// Has returns TRUE if the given flag is present in the bitmask
func (i State) Has(m State) bool {
	if i == m {
		return true
	}
	return i&m != 0
}

// Set returns the bitmask with the given flag set
func (i State) Set(m State) State {
	return i | m
}

// Clear returns the bitmask with the given flag removed
func (i State) Clear(m State) State {
	return i &^ m
}

// Toggle returns the bitmask with the given flag toggled
func (i State) Toggle(m State) State {
	return i ^ m
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the bitmasker command to generate them again.
	var x [1]struct{}
	_ = x[Unconscious-0]
	_ = x[Conscious-2]
	_ = x[Meditative-4]
	_ = x[Distracted-8]
}

const (
	_State_name_0 = "unconscious"
	_State_name_1 = "conscious"
	_State_name_2 = "meditative"
	_State_name_3 = "distracted"
)

func (i State) String() (value string) {
	update := func(t State, n string) {
		if i.Has(t) {
			if len(value) > 0 {
				value += " | "
			}
			value += n
		}
	}
	update(0, _State_name_0)
	update(2, _State_name_1)
	update(4, _State_name_2)
	update(8, _State_name_3)
	if value == "" {
		return "State(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return
}