package main

import (
	"compress/gzip"
	"fmt"
	"github.com/Tnze/go-mc/nbt"
	"os"
)

// <= 1.12.2
// "classic" map format
type Map1343 struct{}

// 1.13
// add field: banners
// del field: width
// del field: height
type Map1519 struct{}

// 1.13.1
// change dimension type to int
// add field: frames
type Map1628 struct{}

// 1.14
// add field: locked (byte)
type Map1952 struct{}

// 1.16
// change dimension type to string (resource location)
type MapXXXX struct{}

// 1.16.5
// add field: UUIDMost (long)
// add field: UUIDLeast (long)
type Map2586 struct{}

type Map struct {
	Data struct {
		Width  int `nbt:"width"`
		Height int `nbt:"height"`

		// Width * Height array of color values (16384 entries for a default 128×128 map).
		// Color can be accessed via the following method:
		// colorID = Colors[widthOffset + heightOffset * width]
		// (where widthOffset==0, heightOffset==0 is left upper point)
		Colors []byte `nbt:"colors"`

		// Center of map according to real world by X.
		XCenter int `nbt:"xCenter"`

		// Center of map according to real world by Z.
		ZCenter int `nbt:"zCenter"`

		// TODO: somehow handle type conflict
		// Resource location of a dimension.
		Dimension string `nbt:"dimension"`
		// Pre-1.16 dimension id:
		// 0 = The Overworld
		// -1 = The Nether
		// 1 = The End
		// any other value = a static image with no player pin.
		DimensionOld byte `nbt:"dimension"`

		// How zoomed in the map is (it is in 2scale wide blocks square per pixel, even for 0, where the map is 1:1). Default 3, minimum 0 and maximum 4.
		Scale byte `nbt:"scale"`

		UuidLeast         int  `nbt:"UUIDLeast"`
		UuidMost          int  `nbt:"UUIDMost"`
		Locked            int  `nbt:"locked"`
		TrackingPosition  byte `nbt:"trackingPosition"`
		UnlimitedTracking byte `nbt:"unlimitedTracking"`

		// TODO: banners
		//  List of banner markers added to this map. May be empty.

		// TODO: frames
		//  List map markers added to this map. May be empty.

	} `nbt:"data"`
	DataVersion int `nbt:"DataVersion"`
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
