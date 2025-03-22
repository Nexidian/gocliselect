package main

import (
	"fmt"
	"github.com/nexidian/gocliselect"
)

func main() {
	menu := gocliselect.NewMenu("Chose a colour")

	menu.AddItem("Red", 1)
	menu.AddItem("Blue", 2)
	menu.AddItem("Green", 3)
	menu.AddItem("Yellow", 4)
	menu.AddItem("Red", 5)
	menu.AddItem("Blue", 6)
	menu.AddItem("Green", 7)
	menu.AddItem("Yellow", 8)

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
