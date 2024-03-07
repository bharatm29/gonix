package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) >= 2 {
		fmt.Printf("%s\n", args[1])
	}
}
