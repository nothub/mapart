package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
)

//go:embed USAGE.txt
var usage string

func main() {

	log.SetFlags(0)

	// no args
	if len(os.Args) < 2 {
		log.Print(usage)
		os.Exit(2)
	}

	for _, s := range []string{"help", "-h", "-help", "--help"} {
		if slices.Contains(os.Args, s) {
			log.Print(usage)
			os.Exit(0)
		}
	}

	err := loadColors()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if os.Args[1] == "colors" {
		fmt.Print(string(colorData))
		os.Exit(0)
	}

	inputs, err := resolveArgs()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if len(inputs) == 0 {
		log.Print(usage)
		os.Exit(2)
	}

	switch os.Args[1] {
	case "png":
		pngCmd(inputs)
	case "fix":
		fixCmd(inputs)
	default:
		log.Print(usage)
		os.Exit(2)
	}
}

func resolveArgs() ([]string, error) {

	var files []string

	for _, arg := range os.Args[2:] {

		path, err := filepath.Abs(arg)
		if err != nil {
			return nil, err
		}

		fi, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		if fi.IsDir() {
			// resolve all files in this directory
			err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
				if d.IsDir() {
					// skip all inner directories
					return nil
				}
				files = append(files, path)
				return nil
			})
			if err != nil {
				return nil, err
			}
		} else {
			files = append(files, path)
		}

	}

	return files, nil
}
