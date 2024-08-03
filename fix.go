package main

import (
	"bytes"
	"compress/gzip"
	"github.com/Tnze/go-mc/nbt"
	"io"
	"log"
	"os"
)

func fixCmd(inputs []string) {
	for _, path := range inputs {
		err := fix(path)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func fix(path string) error {

	m, err := LoadMap(path)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = nbt.NewEncoder(&buf).Encode(m, "")
	if err != nil {
		return err
	}

	// output file
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// wrap in gzip writer
	gzw := gzip.NewWriter(f)
	defer gzw.Close()

	// write data
	_, err = io.Copy(gzw, &buf)
	if err != nil {
		return err
	}

	return nil
}

func dimById(id int) string {
	switch id {
	case 0:
		return "minecraft:overworld"
	case -1:
		return "minecraft:the_nether"
	case 1:
		return "minecraft:the_end"
	default:
		return "minecraft:overworld"
	}
}
