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

	result, err := menu.Display()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if _, ok := result.(int); ok {
		fmt.Printf("Selected option: %d\n", result)
	} else if _, ok := result.(string); ok {
		fmt.Printf("Selected option: %s\n", result)
	} else {
		fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result, result)
	}
}
