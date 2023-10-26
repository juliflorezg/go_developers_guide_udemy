package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// fmt.Println(os.Args)

	// filename
	fn := os.Args[1]
	fmt.Println(fn)

	file, err := os.Open(fn)

	if err != nil {
		fmt.Println("Error while trying to open file::", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
