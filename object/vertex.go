package object

import (
	_ "image/png"
	"simplegraph_sim/static"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vertex struct {
	position Vector
	name     string
	imgW     *ebiten.Image
}

// NewVertex takes a position Vector and an ebiten Image, then returns a new
// nameless Vertex with that position and image.
func NewVertex(pos Vector, img *ebiten.Image) Vertex {
	return Vertex{position: pos, imgW: img}
}

// NewNamedVertex takes a position Vector, a newName string, and an ebiten
// Image, and returns a new named Vertex with that position and image.
func NewNamedVertex(pos Vector, newName string, img *ebiten.Image) Vertex {
	return Vertex{position: pos, name: newName, imgW: img}
}

// GetPosition returns the position Vector of the vertex.
func (v Vertex) GetPosition() Vector {
	return v.position
}

// Draw draws a Vertex based on its position.
func (v *Vertex) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(static.VertScale, static.VertScale)
	op.GeoM.Translate(v.position.x-static.VertShift, v.position.y-static.VertShift)
	screen.DrawImage(v.imgW, op)
}
