package gocliselect

// Terminal control codes
const (
	ShowCursor = "\033[?25h"
	HideCursor = "\033[?25l"
	// CursorUpFormat Requires formatting with number of lines
	CursorUpFormat = "\033[%dA"
	ClearLine      = "\r\033[K"
	KeyUp          = byte(65)
	KeyDown        = byte(66)
	KeyEscape      = byte(27)
	KeyEnter       = byte(13)
)

// NavigationKeys defines a map of specific byte keycodes related to
// navigation functionality, such as up or down actions.
var NavigationKeys = map[byte]bool{
	KeyUp:   true,
	KeyDown: true,
}
