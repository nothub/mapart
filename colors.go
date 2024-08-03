package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"image/color"
	"strconv"
)

//go:generate curl -fsSLO https://github.com/nothub/MapColors/releases/latest/download/colors.csv

//go:embed colors.csv
var colorData []byte

var colors = make(map[byte]color.Color)

func loadColors() error {

	rows, err := csv.NewReader(bytes.NewReader(colorData)).ReadAll()
	if err != nil {
		return err
	}

	for _, row := range rows {

		id, err := strconv.Atoi(row[0])
		if err != nil {
			return err
		}

		r, err := strconv.Atoi(row[1])
		if err != nil {
			return err
		}

		g, err := strconv.Atoi(row[2])
		if err != nil {
			return err
		}

		b, err := strconv.Atoi(row[3])
		if err != nil {
			return err
		}

		a, err := strconv.Atoi(row[4])
		if err != nil {
			return err
		}

		colors[byte(id)] = color.RGBA{R: byte(r), G: byte(g), B: byte(b), A: byte(a)}
	}

	return nil
}
