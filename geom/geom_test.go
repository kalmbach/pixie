package geom

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		v    Vec
		w    Vec
		want Vec
	}{
		{name: "positive", v: Vec{X: 1, Y: 2}, w: Vec{X: 3, Y: 4}, want: Vec{X: 4, Y: 6}},
		{name: "negative", v: Vec{X: 1, Y: 2}, w: Vec{X: -2, Y: -4}, want: Vec{X: -1, Y: -2}},
		{name: "zero", v: Vec{X: 1, Y: 2}, w: Vec{X: 0, Y: 0}, want: Vec{X: 1, Y: 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Add(tt.w); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name string
		v    Vec
		w    Vec
		want Vec
	}{
		{name: "positive", v: Vec{X: 4, Y: 3}, w: Vec{X: 1, Y: 2}, want: Vec{X: 3, Y: 1}},
		{name: "negative", v: Vec{X: 1, Y: 2}, w: Vec{X: 3, Y: 5}, want: Vec{X: -2, Y: -3}},
		{name: "zero", v: Vec{X: 1, Y: 2}, w: Vec{X: 1, Y: 2}, want: Vec{X: 0, Y: 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sub(tt.w); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScale(t *testing.T) {
	tests := []struct {
		name string
		v    Vec
		n    float64
		want Vec
	}{
		{name: "positive", v: Vec{X: 4, Y: 3}, n: 2, want: Vec{X: 8, Y: 6}},
		{name: "negative", v: Vec{X: 1, Y: -2}, n: -2, want: Vec{X: -2, Y: 4}},
		{name: "zero", v: Vec{X: 1, Y: 2}, n: 0, want: Vec{X: 0, Y: 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Scale(tt.n); got != tt.want {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		name string
		r    Rect
		want float64
	}{
		{name: "at negative pos", r: Rect{X: -5, Y: -3, W: 10, H: 4}, want: 5},
		{name: "staying negative", r: Rect{X: -5, Y: -3, W: 2, H: 1}, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Right(); got != tt.want {
				t.Errorf("Right() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBottom(t *testing.T) {
	tests := []struct {
		name string
		r    Rect
		want float64
	}{
		{name: "at negative pos", r: Rect{X: -5, Y: -3, W: 10, H: 4}, want: 1},
		{name: "staying negative", r: Rect{X: -5, Y: -3, W: 2, H: 1}, want: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Bottom(); got != tt.want {
				t.Errorf("Bottom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPos(t *testing.T) {
	tests := []struct {
		name string
		r    Rect
		want Vec
	}{
		{name: "positive", r: Rect{X: 5, Y: 3, W: 10, H: 4}, want: Vec{X: 5, Y: 3}},
		{name: "negative", r: Rect{X: -5, Y: -3, W: 10, H: 4}, want: Vec{X: -5, Y: -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Pos(); got != tt.want {
				t.Errorf("Pos() = %v, want %v", got, tt.want)
			}
		})
	}
}
