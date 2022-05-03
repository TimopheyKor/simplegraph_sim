package object

import (
	"math"
)

type Vector struct {
	x, y float64
}

// NewVector takes an x, y float64 values and returns a new Vector.
func NewVector(xPos, yPos float64) Vector {
	return Vector{x: xPos, y: yPos}
}

// NewVector takes an x, y int values and returns a new Vector.
func NewVectorFromInt(xPos, yPos int) Vector {
	return NewVector(float64(xPos), float64(yPos))
}

// isInRange takes a Vector object b and a float64 proximity value p and
// checks if this Vector is within the proximity / range of the Vector b.
func (a Vector) IsInRange(b Vector, p float64) bool {
	distance := a.Distance(b)
	if distance.x < p && distance.y < p {
		return true // In proximity
	}
	return false // Out of proximity
}

// GetXY returns the x and y fields of the Vector.
func (v Vector) GetXY() (float64, float64) {
	return v.x, v.y
}

// Distance compares this Vector and an external Vector, then returns the
// distance between them as a Vector value.
func (a Vector) Distance(b Vector) Vector {
	xDist := math.Abs(a.x - b.x)
	yDist := math.Abs(a.y - b.y)
	return NewVector(xDist, yDist)
}
