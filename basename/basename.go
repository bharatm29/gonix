package main

import (
	"fmt"
	"os"
	"strings"
)

func Basename(path string) string {
    idx := strings.LastIndexFunc(path, func(c rune) bool {
        return c == '/'
    })

    if idx == -1 {
        return path
    }

    return path[idx + 1:]
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Not enough arguments for basename")
		os.Exit(1)
	}

	path := args[1]

	fmt.Println(Basename(path))
}
