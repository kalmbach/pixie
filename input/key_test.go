package input

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		key  Key
		want string
	}{
		{name: "KeyUnknown", key: KeyUnknown, want: "Unknown"},
		{name: "KeyA", key: KeyA, want: "A"},
		{name: "KeyZ", key: KeyZ, want: "Z"},
		{name: "Key0", key: Key0, want: "0"},
		{name: "Key9", key: Key9, want: "9"},
		{name: "KeySpace", key: KeySpace, want: "Space"},
		{name: "KeyEnter", key: KeyEnter, want: "Enter"},
		{name: "KeyEscape", key: KeyEscape, want: "Escape"},
		{name: "KeyShift", key: KeyShift, want: "Shift"},
		{name: "KeyUp", key: KeyUp, want: "Up"},
		{name: "KeyDown", key: KeyDown, want: "Down"},
		{name: "KeyLeft", key: KeyLeft, want: "Left"},
		{name: "KeyRight", key: KeyRight, want: "Right"},
		{name: "999", key: 999, want: "Key(999)"},
		{name: "-1", key: -1, want: "Key(-1)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.key.String(); got != tt.want {
				t.Errorf("%v.String() = %s, want %s", tt.key, got, tt.want)
			}
		})
	}
}
