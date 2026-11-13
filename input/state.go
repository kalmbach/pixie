package input

// State tracks keyboard input from one tick to the next.
type State struct {
	held     [keyCount]bool
	pressed  [keyCount]bool
	released [keyCount]bool
}

func valid(k Key) bool {
	return k > KeyUnknown && k < keyCount
}

// Press records that k went down.
func (s *State) Press(k Key) {
	if !valid(k) {
		return
	}

	if !s.held[k] {
		s.held[k] = true
		s.pressed[k] = true
	}
}

// Release records that k came up.
func (s *State) Release(k Key) {
	if !valid(k) {
		return
	}

	if s.held[k] {
		s.held[k] = false
		s.released[k] = true
	}
}

// Advance ends the current tick.
func (s *State) Advance() {
	s.pressed = [keyCount]bool{}
	s.released = [keyCount]bool{}
}

// ClearAll releases every held key, and each one reports
// JustReleased until the next tick.
func (s *State) ClearAll() {
	for k := range keyCount {
		s.Release(k)
	}
}

// Held reports whether k is down.
func (s *State) Held(k Key) bool {
	if !valid(k) {
		return false
	}

	return s.held[k]
}

// JustPressed reports whether k went down since the last tick.
func (s *State) JustPressed(k Key) bool {
	if !valid(k) {
		return false
	}

	return s.pressed[k]
}

// JustReleased reports whether k came up since the last tick.
func (s *State) JustReleased(k Key) bool {
	if !valid(k) {
		return false
	}

	return s.released[k]
}
