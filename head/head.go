package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func main() {
	lineNumbers := flag.Int("n", 10, "Total number of line numbers to display")

	flag.Parse()

	if len(flag.Args()) > 0 {
		filename := flag.Arg(0)
		file, err := os.OpenFile(filename, os.O_RDONLY, fs.FileMode.Perm(0777))
		if err != nil {
			fmt.Printf("Error opening file: %s", err)
			os.Exit(1)
		}

		defer func() {
			if err := file.Close(); err != nil {
				fmt.Printf("Could not close file: %s", err)
				os.Exit(1)
			}
		}()

		r := bufio.NewReader(file)
		for i := 0; i < *lineNumbers; i++ {
			line, _, err := r.ReadLine()

			if err == io.EOF {
				break
			}

			fmt.Printf("%s\n", line)
		}
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {

			reader := bufio.NewReader(os.Stdin)

			for i := 0; i < *lineNumbers; i++ {
				line, _, err := reader.ReadLine()

				if err == io.EOF {
					break
				}

				fmt.Printf("%s\n", line)
			}
		} else {
			read := bufio.NewReader(os.Stdin)
			for i := 0; i < *lineNumbers; i++ {
				str, err := read.ReadString('\n')
				if err != nil {
					fmt.Printf("Could not read from command line: %s", err)
				}

				fmt.Printf("%s", str)
			}
		}
	}
}
