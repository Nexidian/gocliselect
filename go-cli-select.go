package go_cli_select

import (
	"bufio"
	"fmt"
	"github.com/buger/goterm"
	"github.com/pkg/term"
	"log"
	"os"
)

type Menu struct {
	Heading   	string
	Prompt  	string
	CursorPos 	int
	MenuItems 	[]*MenuItem
}

type MenuItem struct {
	Text     string
	ID       string
	SubMenu  *Menu
}

func NewMenu(heading string, subheading string) *Menu {
	return &Menu{
		Heading: heading,
		Prompt: subheading,
		MenuItems: make([]*MenuItem, 0),
	}
}

func (m *Menu) AddItem(choice string, id string) *Menu {
	menuItem := &MenuItem{
		Text: choice,
		ID: id,
	}

	m.MenuItems = append(m.MenuItems, menuItem)

	return m
}

func (m *Menu) renderMenuItems() {
	for _, menuItem := range m.MenuItems {
		fmt.Println(menuItem.Text)
	}
}

func (m *Menu) Display() {
	if m.Heading != "" {
		fmt.Println(m.Heading)
	}

	fmt.Printf("%s\n", goterm.Color(goterm.Bold(m.Prompt) + ":", goterm.GREEN))

	m.renderMenuItems()
	m.handleInput()
}

func (m *Menu) handleInput() {
	for {
		getInput()
	}
}

func getInput() {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)
}