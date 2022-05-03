package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("Launching SimpleGraph Simulator...")
	darkMode := 0
	//lightMode := 1

	// Create a new Simulator object, passing in the required parameters.
	simulator := NewSim(darkMode)

	// Set the simulator debug mode.
	simulator.debug = true

	// Set the window variables to launch the executable.
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("SimpleGraph Simulator")
	ebiten.SetWindowResizable(true)

	// Run the simulator if there are no issues.
	if err := ebiten.RunGame(simulator); err != nil {
		log.Fatal(err)
	}
}
