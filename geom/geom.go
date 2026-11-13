// Package geom provides 2D vectors and rectangles in world coordinates.
package geom

// Vec is a 2D vector, used for positions, velocities and sizes.
// Y grows downward.
type Vec struct {
	X, Y float64
}

// Rect is an axis-aligned rectangle with its top-left corner at (X, Y).
// It is half-open: it covers [X, X+W) and [Y, Y+H).
// Y grows downward.
type Rect struct {
	X, Y, W, H float64
}

// Add returns the sum v + w.
func (v Vec) Add(w Vec) Vec {
	return Vec{X: v.X + w.X, Y: v.Y + w.Y}
}

// Sub returns the difference v - w.
func (v Vec) Sub(w Vec) Vec {
	return Vec{X: v.X - w.X, Y: v.Y - w.Y}
}

// Scale returns v with both components multiplied by n.
func (v Vec) Scale(n float64) Vec {
	return Vec{X: v.X * n, Y: v.Y * n}
}

// Pos returns the top-left corner of r.
func (r Rect) Pos() Vec {
	return Vec{X: r.X, Y: r.Y}
}

// Right returns X+W, the first x coordinate past the right edge of r.
func (r Rect) Right() float64 {
	return r.X + r.W
}

// Bottom returns Y+H, the first y coordinate past the bottom edge of r.
func (r Rect) Bottom() float64 {
	return r.Y + r.H
}
