package main

import (
	"flag"
	"fmt"
	"gonix/color"
	"os"
	"path/filepath"
)

type File struct {
	name  string
	perms string
	size  float64
	isDir bool
}

func main() {
	hiddenFlag := flag.Bool("a", false, "Show hidden directories and files")
	longFlag := flag.Bool("l", false, "Show verbose directory and file info")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Print("[ERROR] Not enough arguments")
		os.Exit(1)
	}

	dirs := []File{}
	files := []File{}

	filepath.WalkDir(flag.Arg(0), func(path string, items os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("[ERROR] Could not recursive down the directory: %s", err)
			return filepath.SkipAll
		}

		info, err := items.Info()
		if err != nil {
			fmt.Printf("[ERROR] Could not get directory or file into: %s", err)
			return filepath.SkipAll
		}

		if items.IsDir() {
			dirName := items.Name()

			var coloredName string

			if path == flag.Arg(0) {
				return nil
			}

			if len(dirName) > 0 && dirName[0] == '.' {
				if *hiddenFlag {
					coloredName = (color.Colorize(color.GRAY, fmt.Sprintf("%s/ ", dirName)))
				} else {
					return filepath.SkipDir
				}
			} else {
				coloredName = (color.Colorize(color.RED, fmt.Sprintf("%s/ ", dirName)))
			}

			dir := File{
				name:  coloredName,
				perms: info.Mode().String(),
				size:  float64(info.Size()) / 1000,
				isDir: true,
			}
			dirs = append(dirs, dir)

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

			file := File{
				name:  coloredName,
				perms: info.Mode().String(),
				size:  float64(info.Size()) / 1000,
				isDir: false,
			}
			files = append(files, file)
			return nil
		}
	})

	for _, dir := range dirs {
		if *longFlag {
			fmt.Printf("%s %.1fk %s\n", dir.perms, dir.size, dir.name)
		} else {
			fmt.Printf("%s ", dir.name)
		}
	}

	for _, file := range files {
		if *longFlag {
			fmt.Printf("%s %.1fk %s\n", file.perms, file.size, file.name)
		} else {
			fmt.Printf("%s ", file.name)
		}
	}

	if !(*longFlag) {
		fmt.Print("\n")
	}
}
