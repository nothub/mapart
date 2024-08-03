package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/Tnze/go-mc/nbt"
	"io"
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
	Versioned
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
	Versioned
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
	Versioned
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
	Versioned
}

// 1.16.5
// add field: UUIDMost (long)
// add field: UUIDLeast (long)
type Map2586 struct {
	Data struct {
		UuidMost          int64  `nbt:"UUIDMost"`
		UuidLeast         int64  `nbt:"UUIDLeast"`
		XCenter           int    `nbt:"xCenter"`
		ZCenter           int    `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

type Map struct {
	Data struct {
		Width             int    `nbt:"-"`
		Height            int    `nbt:"-"`
		UuidMost          int64  `nbt:"UUIDMost"`
		UuidLeast         int64  `nbt:"UUIDLeast"`
		XCenter           int    `nbt:"xCenter"`
		ZCenter           int    `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
		Banners           []byte `nbt:"banners,list"`
		Frames            []byte `nbt:"frames,list"`
	} `nbt:"data"`
	Versioned
}

type Versioned struct {
	DataVersion int `nbt:"DataVersion"`
}

func ReadNbt(b []byte) (m Map, err error) {

	var v Versioned

	_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&v)
	if err != nil {
		return m, fmt.Errorf("failed to decode nbt data: %w", err)
	}

	if v.DataVersion <= 1343 {
		var map1343 Map1343
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map1343)
		m.Data.Width = map1343.Data.Width
		m.Data.Height = map1343.Data.Height
		m.Data.UuidMost = 0  // TODO
		m.Data.UuidLeast = 0 // TODO
		m.Data.XCenter = map1343.Data.XCenter
		m.Data.ZCenter = map1343.Data.ZCenter
		m.Data.Scale = map1343.Data.Scale
		m.Data.Dimension = dimById(int(map1343.Data.Dimension))
		m.Data.UnlimitedTracking = map1343.Data.UnlimitedTracking
		m.Data.TrackingPosition = map1343.Data.TrackingPosition
		m.Data.Locked = 0
		m.Data.Colors = map1343.Data.Colors
		m.DataVersion = 1343

	} else if v.DataVersion <= 1519 {
		var map1519 Map1519
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map1519)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0  // TODO
		m.Data.UuidLeast = 0 // TODO
		m.Data.XCenter = map1519.Data.XCenter
		m.Data.ZCenter = map1519.Data.ZCenter
		m.Data.Scale = map1519.Data.Scale
		m.Data.Dimension = dimById(int(map1519.Data.Dimension))
		m.Data.UnlimitedTracking = map1519.Data.UnlimitedTracking
		m.Data.TrackingPosition = map1519.Data.TrackingPosition
		m.Data.Locked = 0
		m.Data.Colors = map1519.Data.Colors
		m.DataVersion = map1519.DataVersion

	} else if v.DataVersion <= 1628 {
		var map1628 Map1628
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map1628)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0  // TODO
		m.Data.UuidLeast = 0 // TODO
		m.Data.XCenter = map1628.Data.XCenter
		m.Data.ZCenter = map1628.Data.ZCenter
		m.Data.Scale = map1628.Data.Scale
		m.Data.Dimension = dimById(map1628.Data.Dimension)
		m.Data.UnlimitedTracking = map1628.Data.UnlimitedTracking
		m.Data.TrackingPosition = map1628.Data.TrackingPosition
		m.Data.Locked = 0
		m.Data.Colors = map1628.Data.Colors
		m.DataVersion = map1628.DataVersion

	} else if v.DataVersion <= 1952 {
		var map1952 Map1952
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map1952)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0  // TODO
		m.Data.UuidLeast = 0 // TODO
		m.Data.XCenter = map1952.Data.XCenter
		m.Data.ZCenter = map1952.Data.ZCenter
		m.Data.Scale = map1952.Data.Scale
		m.Data.Dimension = dimById(map1952.Data.Dimension)
		m.Data.UnlimitedTracking = map1952.Data.UnlimitedTracking
		m.Data.TrackingPosition = map1952.Data.TrackingPosition
		m.Data.Locked = map1952.Data.Locked
		m.Data.Colors = map1952.Data.Colors
		m.DataVersion = map1952.DataVersion

	} else if v.DataVersion <= 2566 {
		var map2566 Map2566
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map2566)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0  // TODO
		m.Data.UuidLeast = 0 // TODO
		m.Data.XCenter = map2566.Data.XCenter
		m.Data.ZCenter = map2566.Data.ZCenter
		m.Data.Scale = map2566.Data.Scale
		m.Data.Dimension = map2566.Data.Dimension
		m.Data.UnlimitedTracking = map2566.Data.UnlimitedTracking
		m.Data.TrackingPosition = map2566.Data.TrackingPosition
		m.Data.Locked = map2566.Data.Locked
		m.Data.Colors = map2566.Data.Colors
		m.DataVersion = map2566.DataVersion

	} else if v.DataVersion <= 2586 {
		var map2586 Map2586
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map2586)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = map2586.Data.UuidMost
		m.Data.UuidLeast = map2586.Data.UuidLeast
		m.Data.XCenter = map2586.Data.XCenter
		m.Data.ZCenter = map2586.Data.ZCenter
		m.Data.Scale = map2586.Data.Scale
		m.Data.Dimension = map2586.Data.Dimension
		m.Data.UnlimitedTracking = map2586.Data.UnlimitedTracking
		m.Data.TrackingPosition = map2586.Data.TrackingPosition
		m.Data.Locked = map2586.Data.Locked
		m.Data.Colors = map2586.Data.Colors
		m.DataVersion = map2586.DataVersion

	} else {
		err = fmt.Errorf("unsupported data version %v", v.DataVersion)
	}

	if err != nil {
		return m, fmt.Errorf("failed to decode nbt data (data version %v): %w", v.DataVersion, err)
	}

	return m, nil
}

func LoadMap(path string) (m Map, err error) {

	fr, err := os.Open(path)
	if err != nil {
		return m, fmt.Errorf("failed to open file: %w", err)
	}
	defer fr.Close()

	// map files are always gzipped
	gzr, err := gzip.NewReader(fr)
	if err != nil {
		return m, fmt.Errorf("failed to decompress gzip: %w", err)
	}

	b, err := io.ReadAll(gzr)
	if err != nil {
		return m, fmt.Errorf("failed to read data: %w", err)
	}

	return ReadNbt(b)
}
