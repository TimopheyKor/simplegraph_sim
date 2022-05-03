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
	ClickRadius  = 50.0
	EdgeSnapping = 20.0
)
