package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	colors["white"] = "#ffffff"

	// fmt.Println(colors)
	// delete(colors, "white")

	// fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, code := range c {
		fmt.Printf("this is the code for the %v color: %v\n", color, code)
	}
}
