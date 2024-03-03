package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting pwd: %s", err)
        os.Exit(1)
	}

	fmt.Println(pwd)
}
