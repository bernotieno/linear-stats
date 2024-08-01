package main

import (
	"fmt"
	"os"

	"linear-stats/reader"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . data.txt")
	}
	filename := os.Args[1]

	reader.InputProcessor(filename)
}
