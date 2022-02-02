# Golang CLI Select
Lightweight interactive CLI selection library 

![](https://media.giphy.com/media/Nmc3muJhaCfPe2LWd9/giphy.gif)


## Import the package
```go
import "github.com/nexidian/gocliselect"
```

## Usage
Create a new menu, supplying the question as a parameter

```go
menu := gocliselect.NewMenu("Chose a colour")
```

Add any number of options by calling `AddItem()` supplying the display text of the option
as well as the id
```go
menu.AddItem("Red", "red")
menu.AddItem("Blue", "blue")
menu.AddItem("Green", "green")
menu.AddItem("Yellow", "yellow")
menu.AddItem("Cyan", "cyan")
```

To display the menu and away the user choice call `Display()`

```go
choice := menu.Display()
```

## Example
```go
package main

import (
    "fmt"
    "github.com/nexidian/gocliselect"
)

func main() {
    menu := gocliselect.NewMenu("Chose a colour")

    menu.AddItem("Red", "red")
    menu.AddItem("Blue", "blue")
    menu.AddItem("Green", "green")
    menu.AddItem("Yellow", "yellow")
    menu.AddItem("Cyan", "cyan")

    choice := menu.Display()

    fmt.Printf("Choice: %s\n", choice)
}
```
