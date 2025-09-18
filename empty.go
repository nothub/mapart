package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func emptyCmd(inputs []string) {
	for _, path := range inputs {

		m, err := LoadMap(path)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		isEmpty := true
		for _, color := range m.Data.Colors {
			if color != 0 {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			fmt.Println(filepath.Abs(path))
		}

	}
}
