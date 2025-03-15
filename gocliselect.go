package gocliselect

import (
	"errors"
	"fmt"
	"github.com/buger/goterm"
	"github.com/pkg/term"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	ErrNoMenuItems = errors.New("menu has no items to display")
)

type Menu struct {
	Prompt    string
	CursorPos int
	MenuItems []*MenuItem
}

type MenuItem struct {
	Text    string
	ID      interface{}
	SubMenu *Menu
}

func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt:    prompt,
		MenuItems: make([]*MenuItem, 0),
	}
}

// AddItem will add a new menu option to the menu list
func (m *Menu) AddItem(option string, id interface{}) *Menu {
	menuItem := &MenuItem{
		Text: option,
		ID:   id,
	}

	m.MenuItems = append(m.MenuItems, menuItem)
	return m
}

// renderMenuItems prints the menu item list.
// Setting redraw to true will re-render the options list with updated current selection.
func (m *Menu) renderMenuItems(redraw bool) {
	if redraw {
		// Move the cursor up n lines where n is the number of options, setting the new
		// location to start printing from, effectively redrawing the option list
		//
		// This is done by sending a VT100 escape code to the terminal
		// @see http://www.climagic.org/mirrors/VT100_Escape_Codes.html
		fmt.Printf(CursorUpFormat, len(m.MenuItems)-1)
	}

	for index, menuItem := range m.MenuItems {
		var newline = "\n"
		if index == len(m.MenuItems)-1 {
			// Adding a new line on the last option will move the cursor position out of range
			// For out redrawing
			newline = ""
		}

		menuItemText := menuItem.Text
		cursor := "  "
		if index == m.CursorPos {
			cursor = goterm.Color("> ", goterm.YELLOW)
			menuItemText = goterm.Color(menuItemText, goterm.YELLOW)
		}

		fmt.Printf("\r%s %s%s", cursor, menuItemText, newline)
	}
}

// Display will display the current menu options and awaits user selection
// It returns the users selected choice
func (m *Menu) Display() (interface{}, error) {
	defer func() {
		// Show cursor again.
		fmt.Printf(ShowCursor)
	}()

	if len(m.MenuItems) == 0 {
		return nil, ErrNoMenuItems
	}

	fmt.Printf("%s\n", goterm.Color(goterm.Bold(m.Prompt)+":", goterm.CYAN))

	m.renderMenuItems(false)

	// Turn the terminal cursor off
	fmt.Printf(HideCursor)

	// Channel to signal interrupt
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	for {
		keyCode := getInput()
		switch keyCode {
		case KeyEscape:
			return "", nil
		case KeyEnter:
			menuItem := m.MenuItems[m.CursorPos]
			fmt.Println("\r")
			return menuItem.ID, nil
		case KeyUp:
			m.CursorPos = (m.CursorPos + len(m.MenuItems) - 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		case KeyDown:
			m.CursorPos = (m.CursorPos + 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		}
	}
}

// getInput will read raw input from the terminal
// It returns the raw ASCII value inputted
func getInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	defer t.Close() // Close t in defer to ensure it's always closed

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)
	if err != nil {
		// Handle read error, it might be due to signal interruption
		return 0 // Or some other value indicating error/interruption if needed
	}

	defer t.Restore() // Restore terminal mode in defer
	
	// Arrow keys are prefixed with the ANSI escape code which take up the first two bytes.
	// The third byte is the key specific value we are looking for.
	// For example the left arrow key is '<esc>[A' while the right is '<esc>[C'
	// See: https://en.wikipedia.org/wiki/ANSI_escape_code
	if read == 3 {
		if _, ok := NavigationKeys[readBytes[2]]; ok {
			return readBytes[2]
		}
	} else {
		return readBytes[0]
	}

	return 0
}
