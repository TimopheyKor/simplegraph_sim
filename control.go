package main

import (
	"simplegraph_sim/object"
	"simplegraph_sim/static"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// RunCursorCheck gets the current position of the mouse cursor, then
// checks to see what state the mouse is in. The internal sub-checks for this
// provide their own responses based on the user's actions.
func (sim *Simulator) RunCursorChecks() {
	x, y := ebiten.CursorPosition()
	cursorPos := object.NewVectorFromInt(x, y)
	sim.CheckClick(cursorPos)
	sim.CheckDrag(cursorPos)
	sim.CheckRelease(cursorPos)
}

// CheckClick checks if the mouse is pressed, runs checks on what type of click
// it was and if any existing objects were selected, then calls the appropriate
// functions for the type of click.
func (sim *Simulator) CheckClick(cursorPos object.Vector) {
	// Get a Vertex if it's within SelectRad of the cursor; get nil otherwise.
	hoveredVert := sim.getNearbyVertex(cursorPos, static.SelectRad)
	// Select or create a new Vertex on left click.
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		sim.selected = hoveredVert
		// Check if there are any other Vertices within VertPad of the click
		occupied := sim.getNearbyVertex(cursorPos, static.VertPad)
		if sim.selected == nil && occupied == nil {
			sim.addVertex(cursorPos)
		}
	}
	// Delete the vertex that's being hovered over.
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		if hoveredVert != nil {
			sim.graph.RemoveVertex(*hoveredVert)
		}
	}
}

// CheckDrag checks to see if the mouse is being dragged from a selected
// Vertex, and creates a temporary edge from the selected vertex to the cursor.
func (sim *Simulator) CheckDrag(cursorPos object.Vector) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if sim.selected != nil {
			vert := object.NewVertex(cursorPos, nil)
			edge := object.NewEdge(sim.selected, &vert)
			sim.tempEdge = &edge
		}
	}
}

// Check release checks to see if the mouse was released on a frame. If there
// is a new Edge being drawn and it's within range of another Vertex, adds the
// Edge to the Graph, and removes the tempEdge. Otherwise, it just removes
// the tempEdge.
func (sim *Simulator) CheckRelease(cursorPos object.Vector) {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if sim.tempEdge != nil && sim.selected != nil {
			vert := sim.getNearbyVertex(cursorPos, static.SnapRad)
			if vert != nil {
				start, _ := sim.tempEdge.GetVerts()
				sim.graph.AddEdge(*start, *vert)
			}
			sim.tempEdge = nil
		}
	}
}

// getNearbyVertex takes a position Vector c and a proximity float64 value p,
// then returns a Vertex if there is one within the proximity of the position.
func (sim *Simulator) getNearbyVertex(c object.Vector, p float64) *object.Vertex {
	for vert := range sim.graph.Verts {
		vPos := vert.GetPosition()
		if c.IsInRange(vPos, p) {
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
