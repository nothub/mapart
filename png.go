package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
