package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"simplegraph_sim/object"
	"simplegraph_sim/static"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Simulator implements ebiten's Game interface.
type Simulator struct {
	debug        bool
	graph        object.Graph
	main, accent color.Color
	selected     *object.Vertex
	tempEdge     *object.Edge
	pointImg     *ebiten.Image
}

// Update proceeds the game state, and is called every tick.
func (sim *Simulator) Update() error {
	sim.RunCursorChecks()
	return nil
}

// Draw draws the game screen, and is called every tick.
func (sim *Simulator) Draw(screen *ebiten.Image) {
	screen.Fill(sim.main)

	// Draw all the vertecies and edges.
	for _, edge := range sim.graph.Edges {
		edge.Draw(screen)
	}
	for vert := range sim.graph.Verts {
		vert.Draw(screen)
	}

	// Draw the user dragging an edge out from the selected Vertex.
	if sim.tempEdge != nil {
		sim.tempEdge.Draw(screen)
	}

	// Debug draws
	if sim.debug {
		ebitenutil.DebugPrint(screen, "Hello, World!\n")
		x, y := ebiten.CursorPosition()
		s := fmt.Sprintf("\nCursor Position: %v, %v\n", x, y)
		ebitenutil.DebugPrint(screen, s)
	}
}

// NewSim creates a new Simulator object and returns a pointer to it.
// Parameters: mode int.
// mode: 0 = "Dark Mode", 1 = "Light Mode".
func NewSim(mode int) *Simulator {
	sim := &Simulator{}
	sim.loadImages()
	sim.setColor(mode)
	sim.graph = object.NewGraph()
	return sim
}

// Layout takes the outside window size and returns the logical screen size.
func (sim *Simulator) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// SetColor takes an int value and sets the main and accent colors of the
// Simulator object.
//
// 0: "Dark Mode"
// 1: "Light Mode"
// Default: "Dark Mode"
func (sim *Simulator) setColor(mode int) {
	switch mode {
	case 0:
		sim.main = static.ColorGrey
		sim.accent = static.ColorWhite
	case 1:
		sim.main = static.ColorWhite
		sim.accent = static.ColorBlack
	default:
		sim.main = static.ColorGrey
		sim.accent = static.ColorWhite
	}
}

// loadImages loads all the necessary image files from the assets folder
// into the Simulator.
func (sim *Simulator) loadImages() {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile("assets/whitevert.png")
	if err != nil {
		log.Fatal(err)
	}
	sim.pointImg = img
}
