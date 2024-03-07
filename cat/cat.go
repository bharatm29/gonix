package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func setFlags() map[string]*bool {
	flags := []struct {
		alias string
		desc  string
		def   bool
	}{
		{
			alias: "n",
			def:   false,
			desc:  "Enable newline at the end",
		},
		{
			alias: "h",
			def:   false,
			desc:  "Show help",
		},
	}

	aliases := map[string]*bool{}

	for _, fg := range flags {
		aliases[fg.alias] = flag.Bool(fg.alias, fg.def, fg.desc)
	}

	return aliases
}

func main() {
	_ = setFlags()

	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		filename := args[0]
		file, err := os.OpenFile(filename, os.O_RDONLY, fs.FileMode.Perm(0777))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %s", err)
			os.Exit(1)
		}

		defer func() {
			if err := file.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "Could not close file: %s", err)
				os.Exit(1)
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

		fmt.Printf("Contents of the file: %s\n%s", filename, fileContent)
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
				fmt.Fprintf(os.Stderr, "Could not scan from piped stdin: %s", err)
				os.Exit(1)
			}
			str := string(stdin)
			fmt.Printf("%s", str)
		} else {
			read := bufio.NewReader(os.Stdin)
			str, err := read.ReadString(byte(24))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not read from command line: %s", err)
			}

			fmt.Printf("%s", str)
			if str[len(str)-1] != '\n' {
				fmt.Print("\n")
			}
		}
	}
}
