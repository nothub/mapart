package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
)

func emptyCmd(inputs []string) {

	r, _ := regexp.Compile("map_[0-9]+\\.dat")

	for _, path := range inputs {

		if !r.MatchString(filepath.Base(path)) {
			log.Println("skipping (not a map): " + path)
			continue
		}

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
			abs, _ := filepath.Abs(path)
			fmt.Println(abs)
		} else {
			log.Println("skipping (not empty): " + path)
		}

	}
}
