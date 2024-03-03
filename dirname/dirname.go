package main

import (
	"fmt"
	"os"
	"strings"
)

func Dirname(path string) string {
	if len(path) == 1 && path[0] == '/' {
		return path
	}

	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	idx := strings.LastIndexFunc(path, func(c rune) bool {
		return c == '/'
	})

	first_idx := strings.IndexFunc(path, func(c rune) bool {
		return c == '/'
	})

	if first_idx == idx {
		return path
	}

	if idx == -1 {
		return path
	}

	return path[:idx]
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Not enough arguments for dirname")
		os.Exit(1)
	}

	path := args[1]

	fmt.Println(Dirname(path))
}
