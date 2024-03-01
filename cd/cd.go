package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		filename := args[1]
		file, err := os.OpenFile(filename, os.O_RDONLY, fs.FileMode.Perm(0777))
		if err != nil {
			log.Fatalf("Error opening file: %s", err)
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.Fatalf("Could not close file: %s", err)
			}
		}()

		fileContent := []byte{}
		buf := make([]byte, 4096)

		for {
			n, err := file.Read(buf)

			if n > 0 {
				fileContent = append(fileContent, buf...)
			}

			if err == io.EOF {
				break
			}
		}

		log.Printf("Contents of the file: %s\n%s", filename, fileContent)
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			var stdin []byte
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				stdin = append(stdin, scanner.Bytes()...)
				stdin = append(stdin, byte('\n'))
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			str := string(stdin)
			log.Printf("stdin = %s\n", str)
		} else {
			read := bufio.NewReader(os.Stdin)
			str, err := read.ReadString(byte(24))
			if err != nil {
				log.Fatalf("Could not read from command line: %s", err)
			}

			fmt.Printf("%s", str)
            if str[len(str) - 1] != '\n'{
                fmt.Print("\n")
            }
		}
	}
}
