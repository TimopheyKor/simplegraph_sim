package static

import (
	"image/color"
)

// Global static color variables
var (
	ColorGrey  = color.RGBA{33, 33, 33, 0xff}
	ColorWhite = color.White
	ColorBlack = color.Black
)

// Global static proximity variables
var (
	VertPad   = 20.0
	SnapRad   = 20.0
	SelectRad = 5.0
)

// Image transformation variables
var (
	VertScale = 0.1
	VertShift = 5.0
)
