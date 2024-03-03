package main

import (
	"flag"
	"fmt"
	"gonix/color"
	"os"
	"path/filepath"
)

func main() {
	hiddenFlag := flag.Bool("a", false, "Show hidden directories and files")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Print("[ERROR] Not enough arguments")
		os.Exit(1)
	}

	filenames := []string{}

	filepath.WalkDir(flag.Arg(0), func(path string, items os.DirEntry, err error) error {
		/*TODO: Implement -l
		  info, err := items.Info()
		*/

		if err != nil {
			fmt.Printf("[ERROR] Could not recursive down the directory: %s", err)
			return filepath.SkipAll
		}
		if items.IsDir() {
			dirName := items.Name()

			if path == flag.Arg(0) {
				return nil
			}

			if len(dirName) > 0 && dirName[0] == '.' {
				if *hiddenFlag {
					fmt.Print(color.Colorize(color.GRAY, fmt.Sprintf("%s/ ", dirName)))
				} else {
					return filepath.SkipDir
				}
			} else {
				fmt.Print(color.Colorize(color.RED, fmt.Sprintf("%s/ ", dirName)))
			}

			return filepath.SkipDir
		} else {
			fileName := items.Name()

			var coloredName string

			if len(fileName) > 0 && fileName[0] == '.' {
				if *hiddenFlag {
					coloredName = color.Colorize(color.GRAY, fmt.Sprintf("%s ", fileName))
				} else {
					return nil
				}
			} else {
				coloredName = color.Colorize(color.BLUE, fmt.Sprintf("%s ", fileName))
			}

			filenames = append(filenames, coloredName)
			return nil
		}
	})

	for _, filename := range filenames {
		fmt.Print(filename + " ")
	}

	fmt.Printf("\n")
}
