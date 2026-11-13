// Package input tracks keyboard state from tick to tick.
package input

import "fmt"

// Key identifies a keyboard key.
type Key int

const (
	KeyUnknown Key = iota

	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ

	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9

	KeyUp
	KeyDown
	KeyLeft
	KeyRight

	KeySpace
	KeyEnter
	KeyEscape
	KeyShift

	keyCount
)

func letter(k Key) string {
	return string(rune('A' + k - KeyA))
}

func number(k Key) string {
	return string(rune('0' + k - Key0))
}

// String returns the key's name.
func (k Key) String() string {
	switch {
	case k == KeyUnknown:
		return "Unknown"
	case k >= KeyA && k <= KeyZ:
		return letter(k)
	case k >= Key0 && k <= Key9:
		return number(k)
	case k == KeySpace:
		return "Space"
	case k == KeyEnter:
		return "Enter"
	case k == KeyEscape:
		return "Escape"
	case k == KeyShift:
		return "Shift"
	case k == KeyUp:
		return "Up"
	case k == KeyDown:
		return "Down"
	case k == KeyLeft:
		return "Left"
	case k == KeyRight:
		return "Right"
	default:
		return fmt.Sprintf("Key(%d)", k)
	}
}
