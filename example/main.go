package main

import (
	"fmt"
	"github.com/nexidian/gocliselect"
)

func main() {
	menu := gocliselect.NewMenu("Chose a colour")

	menu.AddItem("Red", "red")
	menu.AddItem("Blue", 1)
	menu.AddItem("Green", 123)
	menu.AddItem("Yellow", 1.0)
	menu.AddItem("Cyan", "cyan")

	result := menu.Display()

	if id, ok := result.(int); ok {
		fmt.Printf("Choice int: %d\n", id)
	} else if id, ok := result.(string); ok {
		fmt.Printf("Choice string: %s\n", id)
	} else {
		fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result, result)
	}
}
