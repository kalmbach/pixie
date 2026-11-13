package input

import (
	"testing"
)

func assertState(t *testing.T, s *State, k Key, held, pressed, released bool) {
	t.Helper()

	if got := s.Held(k); got != held {
		t.Errorf("s.Held(%v) = %t, want %t", k, got, held)
	}

	if got := s.JustPressed(k); got != pressed {
		t.Errorf("s.JustPressed(%v) = %t, want %t", k, got, pressed)
	}

	if got := s.JustReleased(k); got != released {
		t.Errorf("s.JustReleased(%v) = %t, want %t", k, got, released)
	}
}

func TestPress(t *testing.T) {
	var s State

	s.Press(KeyA)

	assertState(t, &s, KeyA, true, true, false)
}

func TestAfterAdvance(t *testing.T) {
	var s State

	s.Press(KeyA)
	s.Advance()

	assertState(t, &s, KeyA, true, false, false)
}

func TestRelease(t *testing.T) {
	var s State

	s.Press(KeyA)
	s.Advance()
	s.Release(KeyA)

	assertState(t, &s, KeyA, false, false, true)

	s.Advance()

	assertState(t, &s, KeyA, false, false, false)
}

func TestFastTap(t *testing.T) {
	var s State

	s.Press(KeyA)
	s.Release(KeyA)

	assertState(t, &s, KeyA, false, true, true)
}

func TestPressAfterAdvance(t *testing.T) {
	var s State

	s.Press(KeyA)
	s.Advance()
	s.Press(KeyA)

	assertState(t, &s, KeyA, true, false, false)
}

func TestReleaseUnheldKey(t *testing.T) {
	var s State

	s.Release(KeyA)

	assertState(t, &s, KeyA, false, false, false)
}

func TestUnknown(t *testing.T) {
	var s State

	s.Press(KeyUnknown)

	assertState(t, &s, KeyUnknown, false, false, false)

	s.Release(KeyUnknown)

	assertState(t, &s, KeyUnknown, false, false, false)
}

func TestOutOfRange(t *testing.T) {
	var s State

	s.Press(-1)
	s.Release(-1)
	assertState(t, &s, keyCount+1, false, false, false)

	s.Press(keyCount)
	s.Release(keyCount)
	assertState(t, &s, keyCount+1, false, false, false)

	s.Press(keyCount + 1)
	s.Release(keyCount + 1)
	assertState(t, &s, keyCount+1, false, false, false)
}

func TestClearAll(t *testing.T) {
	var s State

	s.Press(KeyA)
	s.Press(KeyLeft)
	s.Advance()
	s.ClearAll()

	assertState(t, &s, KeyA, false, false, true)
	assertState(t, &s, KeyLeft, false, false, true)
	assertState(t, &s, KeyB, false, false, false)

	s.Advance()
	s.ClearAll()

	assertState(t, &s, KeyA, false, false, false)
	assertState(t, &s, KeyLeft, false, false, false)
	assertState(t, &s, KeyB, false, false, false)
}

func TestReleaseAfterClearAll(t *testing.T) {
	var s State

	s.Press(KeyA)
	s.Advance()
	s.ClearAll()
	s.Advance()
	s.Release(KeyA)

	assertState(t, &s, KeyA, false, false, false)
}
