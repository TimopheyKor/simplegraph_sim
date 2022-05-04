package object

import (
	"simplegraph_sim/static"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Edge struct {
	start, finish *Vertex
}

// NewEdge takes two pointers to Vertex objects and returns a new Edge between
// them.
func NewEdge(one, two *Vertex) Edge {
	return Edge{start: one, finish: two}
}

func (e *Edge) GetVerts() (*Vertex, *Vertex) {
	return e.start, e.finish
}

// Draw draws a line between the two Vertex objects of an Edge.
func (e *Edge) Draw(screen *ebiten.Image) {
	sx, sy := e.start.GetPosition().GetXY()
	fx, fy := e.finish.GetPosition().GetXY()
	ebitenutil.DrawLine(screen, sx, sy, fx, fy, static.ColorWhite)
}
