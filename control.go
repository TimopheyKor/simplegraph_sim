package main

import (
	"fmt"
	"simplegraph_sim/object"
	"simplegraph_sim/static"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// CheckClick checks if the mouse is pressed, runs checks on what type of click
// it was and if any existing objects were selected, then calls the appropriate
// functions for the type of click.
func (sim *Simulator) CheckClick() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		clickPos := object.NewVectorFromInt(x, y)
		sim.selected = sim.getClickedObject(clickPos)
		if sim.selected == nil {
			sim.addVertex(clickPos)
		}
	}
}

// TODO: Separate the logic of a new line into a field in the Simulator struct.
// TODO: Separate the drawing of a line into a function without update logic.
// DrawEdgeDrag checks to see if the mouse is being dragged from a selected
// Vertex, and draws a line from that vertex to the cursor position.
func (sim *Simulator) DrawEdgeDrag(screen *ebiten.Image) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if sim.selected != nil {
			x, y := ebiten.CursorPosition()
			cx, cy := float64(x), float64(y)
			dx, dy := sim.selected.GetPosition().GetXY()
			fmt.Println(cx, cy, dx, dy)
			ebitenutil.DrawLine(screen, cx, cy, dx, dy, static.ColorWhite)
		}
	}
}

// getClickedObject takes x, y positional float64 values, and checks if there is
// an existing vertex in the simulator within click range of that position.
func (sim *Simulator) getClickedObject(click object.Vector) *object.Vertex {
	for vert := range sim.graph.Data {
		vPos := vert.GetPosition()
		if click.IsInRange(vPos, static.ClickRadius) {
			return &vert
		}
	}
	return nil
}

// addVertex takes positional values, creates a new vertex with those values,
// and adds it to the simulator's graph.
func (sim *Simulator) addVertex(position object.Vector) {
	v := object.NewVertex(position, sim.pointImg)
	sim.graph.AddVertex(v)
}
