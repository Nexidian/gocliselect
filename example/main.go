package main

import "github.com/nexidian/gocliselect"

func main() {
	menu := gocliselect.NewMenu("", "Select an option")

	menu.AddItem("Option 1", "option1")
	menu.AddItem("Option 2", "option2")
	menu.AddItem("Option 3", "option3")
	menu.AddItem("Option 4", "option4")
	menu.AddItem("Option 5", "option5")

	menu.Display()
}