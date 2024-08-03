package main

import (
	"compress/gzip"
	"fmt"
	"github.com/Tnze/go-mc/nbt"
	"os"
)

// <= 1.12.2
// classic map format
type Map1343 struct {
	Data struct {
		Width             int    `nbt:"width"`
		Height            int    `nbt:"height"`
		XCenter           int    `nbt:"xCenter"`
		ZCenter           int    `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         byte   `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
}

// 1.13
// add field: banners (but we ignore this anyways)
// del field: width
// del field: height
type Map1519 struct {
	Data struct {
		XCenter           int    `nbt:"xCenter"`
		ZCenter           int    `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         byte   `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
}

// 1.13.1
// change dimension type to int
// add field: frames (but we ignore this anyways)
type Map1628 struct {
	Data struct {
		XCenter           int    `nbt:"xCenter"`
		ZCenter           int    `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         int    `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
}

// 1.14
// add field: locked (byte)
type Map1952 struct {
	Data struct {
		XCenter           int    `nbt:"xCenter"`
		ZCenter           int    `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         int    `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
}

// 1.16
// change dimension type to string (resource location)
type Map2566 struct {
	Data struct {
		XCenter           int    `nbt:"xCenter"`
		ZCenter           int    `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
}

// 1.16.5
// add field: UUIDMost (long)
// add field: UUIDLeast (long)
type Map2586 struct {
	Data struct {
		UUIDMost          int    `nbt:"UUIDMost"`
		UUIDLeast         int    `nbt:"UUIDLeast"`
		XCenter           int    `nbt:"xCenter"`
		ZCenter           int    `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
}

type Map Map2586

type Versioned struct {
	DataVersion int `nbt:"DataVersion"`
}

func readData(version int) (m Map, err error) {
	if version <= 1343 {
		var map1343 Map1343
		_ = map1343
		// TODO
	}
	if version <= 1519 {
		var map1519 Map1519
		_ = map1519
		// TODO
	}
	if version <= 1628 {
		var map1628 Map1628
		_ = map1628
		// TODO
	}
	if version <= 1952 {
		var map1952 Map1952
		_ = map1952
		// TODO
	}
	if version <= 2566 {
		var map2566 Map2566
		_ = map2566
		// TODO
	}
	if version <= 2586 {
		var map2586 Map2586
		_ = map2586
		// TODO
	}
	return m, nil
}

func LoadMap(path string) (*Map, error) {

	fr, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer fr.Close()

	// map files are always gzipped
	gzr, err := gzip.NewReader(fr)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress gzip: %w", err)
	}

	var m Map

	_, err = nbt.NewDecoder(gzr).Decode(&m)
	if err != nil {
		return nil, fmt.Errorf("failed to decode nbt data: %w", err)
	}

	return &m, nil
}
