package main

import (
	_ "embed"
	"fmt"
	"github.com/fogleman/gg"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
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

func pngCmd(inputs []string) {
	for _, path := range inputs {
		err := toPng(path)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func toPng(path string) error {

	m, err := LoadMap(path)
	if err != nil {
		return err
	}

	dc := gg.NewContext(m.Data.Width, m.Data.Height)
	for y := 0; y < m.Data.Height; y++ {
		for x := 0; x < m.Data.Width; x++ {
			colorID := m.Data.Colors[y*m.Data.Width+x]
			dc.SetColor(colors[colorID])
			dc.SetPixel(x, y)
		}
	}

	f, err := os.Create(strings.TrimSuffix(filepath.Base(path), ".dat") + ".png")
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	err = png.Encode(f, dc.Image())
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}

func fixCmd(inputs []string) {

	// TODO

	/*
		if m.Data.Dimension == "" {
			m.Data.Dimension = dimById(m.Data.DimensionOld)
		}
	*/

}
