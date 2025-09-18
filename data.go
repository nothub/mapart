package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/Tnze/go-mc/nbt"
	"io"
	"os"
)

const LatestDataVersion = 2586

// >= 1.0 && <= 1.8.9
// classic map format
type MapClassic struct {
	Data struct {
		Width     int    `nbt:"width"`
		Height    int    `nbt:"height"`
		XCenter   int32  `nbt:"xCenter"`
		ZCenter   int32  `nbt:"zCenter"`
		Scale     byte   `nbt:"scale"`
		Dimension byte   `nbt:"dimension"`
		Colors    []byte `nbt:"colors"`
	} `nbt:"data"`
}

// >= 1.9 && <= 1.10.2
// add field: trackingPosition (byte)
type Map169 struct {
	Data struct {
		Width            int    `nbt:"width"`
		Height           int    `nbt:"height"`
		XCenter          int32  `nbt:"xCenter"`
		ZCenter          int32  `nbt:"zCenter"`
		Scale            byte   `nbt:"scale"`
		Dimension        byte   `nbt:"dimension"`
		TrackingPosition byte   `nbt:"trackingPosition"`
		Colors           []byte `nbt:"colors"`
	} `nbt:"data"`
}

// >= 1.11 && <= 1.12.2
// add field: unlimitedTracking (byte)
type Map819 struct {
	Data struct {
		Width             int    `nbt:"width"`
		Height            int    `nbt:"height"`
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         byte   `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
}

type Versioned struct {
	DataVersion int32 `nbt:"DataVersion"`
}

// >= 1.13 && <= 1.13
// add field: banners (but we ignore this anyway)
// del field: width
// del field: height
type Map1519 struct {
	Data struct {
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         byte   `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

// >= 1.13.1 && <= 1.13.2
// change dimension type to int
// add field: frames (but we ignore this anyways)
type Map1628 struct {
	Data struct {
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         int    `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

// >= 1.14 && <= 1.14.4
// add field: locked (byte)
type Map1952 struct {
	Data struct {
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         int    `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

// >= 1.16 && <= 1.16.4
// change dimension type to string (resource location)
type Map2566 struct {
	Data struct {
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

// 1.16.5 - 1.19.4
// add field: UUIDMost (long)
// add field: UUIDLeast (long)
type Map2586 struct {
	Data struct {
		UuidMost          int64  `nbt:"UUIDMost"`
		UuidLeast         int64  `nbt:"UUIDLeast"`
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

// 1.20 - 1.20.5
// remove field: UUIDMost (long)
// remove field: UUIDLeast (long)
type Map3463 struct {
	Data struct {
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

// 1.20.6 - 1.21
// add field: UUIDMost (long)
// add field: UUIDLeast (long)
type Map3839 struct {
	Data struct {
		UuidMost          int64  `nbt:"UUIDMost"`
		UuidLeast         int64  `nbt:"UUIDLeast"`
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

// 1.21.1 - 1.21.4
// remove field: UUIDMost (long)
// remove field: UUIDLeast (long)
type Map3955 struct {
	Data struct {
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
		Scale             byte   `nbt:"scale"`
		Dimension         string `nbt:"dimension"`
		UnlimitedTracking byte   `nbt:"unlimitedTracking"`
		TrackingPosition  byte   `nbt:"trackingPosition"`
		Locked            byte   `nbt:"locked"`
		Colors            []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

// >= 1.21.5
// remove field: locked (byte)
// remove field: scale (byte)
// remove field: trackingPosition (byte)
// remove field: unlimitedTracking (byte)
// remove field: banners (array)
// remove field: frames (array)
type Map4325 struct {
	Data struct {
		XCenter   int32  `nbt:"xCenter"`
		ZCenter   int32  `nbt:"zCenter"`
		Dimension string `nbt:"dimension"`
		Colors    []byte `nbt:"colors"`
	} `nbt:"data"`
	Versioned
}

type Map struct {
	Data struct {
		Width             int    `nbt:"-"`
		Height            int    `nbt:"-"`
		UuidMost          int64  `nbt:"UUIDMost"`
		UuidLeast         int64  `nbt:"UUIDLeast"`
		XCenter           int32  `nbt:"xCenter"`
		ZCenter           int32  `nbt:"zCenter"`
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

func ReadNbt(b []byte) (m Map, err error) {

	var v Versioned

	_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&v)
	if err != nil {
		return m, fmt.Errorf("failed to decode nbt data: %w", err)
	}

	if v.DataVersion <= 1343 {
		var map1343 Map819
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map1343)
		m.Data.Width = map1343.Data.Width
		m.Data.Height = map1343.Data.Height
		m.Data.UuidMost = 0
		m.Data.UuidLeast = 0
		m.Data.XCenter = map1343.Data.XCenter
		m.Data.ZCenter = map1343.Data.ZCenter
		m.Data.Scale = map1343.Data.Scale
		m.Data.Dimension = dimById(int(map1343.Data.Dimension))
		m.Data.UnlimitedTracking = map1343.Data.UnlimitedTracking
		m.Data.TrackingPosition = map1343.Data.TrackingPosition
		m.Data.Locked = 0
		m.Data.Colors = map1343.Data.Colors
		m.DataVersion = LatestDataVersion

	} else if v.DataVersion <= 1519 {
		var map1519 Map1519
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map1519)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0
		m.Data.UuidLeast = 0
		m.Data.XCenter = map1519.Data.XCenter
		m.Data.ZCenter = map1519.Data.ZCenter
		m.Data.Scale = map1519.Data.Scale
		m.Data.Dimension = dimById(int(map1519.Data.Dimension))
		m.Data.UnlimitedTracking = map1519.Data.UnlimitedTracking
		m.Data.TrackingPosition = map1519.Data.TrackingPosition
		m.Data.Locked = 0
		m.Data.Colors = map1519.Data.Colors
		m.DataVersion = LatestDataVersion

	} else if v.DataVersion <= 1628 {
		var map1628 Map1628
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map1628)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0
		m.Data.UuidLeast = 0
		m.Data.XCenter = map1628.Data.XCenter
		m.Data.ZCenter = map1628.Data.ZCenter
		m.Data.Scale = map1628.Data.Scale
		m.Data.Dimension = dimById(map1628.Data.Dimension)
		m.Data.UnlimitedTracking = map1628.Data.UnlimitedTracking
		m.Data.TrackingPosition = map1628.Data.TrackingPosition
		m.Data.Locked = 0
		m.Data.Colors = map1628.Data.Colors
		m.DataVersion = LatestDataVersion

	} else if v.DataVersion <= 1952 {
		var map1952 Map1952
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map1952)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0
		m.Data.UuidLeast = 0
		m.Data.XCenter = map1952.Data.XCenter
		m.Data.ZCenter = map1952.Data.ZCenter
		m.Data.Scale = map1952.Data.Scale
		m.Data.Dimension = dimById(map1952.Data.Dimension)
		m.Data.UnlimitedTracking = map1952.Data.UnlimitedTracking
		m.Data.TrackingPosition = map1952.Data.TrackingPosition
		m.Data.Locked = map1952.Data.Locked
		m.Data.Colors = map1952.Data.Colors
		m.DataVersion = LatestDataVersion

	} else if v.DataVersion <= 2566 {
		var map2566 Map2566
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map2566)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0
		m.Data.UuidLeast = 0
		m.Data.XCenter = map2566.Data.XCenter
		m.Data.ZCenter = map2566.Data.ZCenter
		m.Data.Scale = map2566.Data.Scale
		m.Data.Dimension = map2566.Data.Dimension
		m.Data.UnlimitedTracking = map2566.Data.UnlimitedTracking
		m.Data.TrackingPosition = map2566.Data.TrackingPosition
		m.Data.Locked = map2566.Data.Locked
		m.Data.Colors = map2566.Data.Colors
		m.DataVersion = LatestDataVersion

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
		m.DataVersion = LatestDataVersion

	} else if v.DataVersion <= 3463 {
		var map3463 Map3463
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map3463)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0
		m.Data.UuidLeast = 0
		m.Data.XCenter = map3463.Data.XCenter
		m.Data.ZCenter = map3463.Data.ZCenter
		m.Data.Scale = map3463.Data.Scale
		m.Data.Dimension = map3463.Data.Dimension
		m.Data.UnlimitedTracking = map3463.Data.UnlimitedTracking
		m.Data.TrackingPosition = map3463.Data.TrackingPosition
		m.Data.Locked = map3463.Data.Locked
		m.Data.Colors = map3463.Data.Colors
		m.DataVersion = LatestDataVersion

	} else if v.DataVersion <= 3839 {
		var map3839 Map3839
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map3839)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = map3839.Data.UuidMost
		m.Data.UuidLeast = map3839.Data.UuidLeast
		m.Data.XCenter = map3839.Data.XCenter
		m.Data.ZCenter = map3839.Data.ZCenter
		m.Data.Scale = map3839.Data.Scale
		m.Data.Dimension = map3839.Data.Dimension
		m.Data.UnlimitedTracking = map3839.Data.UnlimitedTracking
		m.Data.TrackingPosition = map3839.Data.TrackingPosition
		m.Data.Locked = map3839.Data.Locked
		m.Data.Colors = map3839.Data.Colors
		m.DataVersion = LatestDataVersion

	} else if v.DataVersion <= 3955 {
		var map3955 Map3955
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map3955)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0
		m.Data.UuidLeast = 0
		m.Data.XCenter = map3955.Data.XCenter
		m.Data.ZCenter = map3955.Data.ZCenter
		m.Data.Scale = map3955.Data.Scale
		m.Data.Dimension = map3955.Data.Dimension
		m.Data.UnlimitedTracking = map3955.Data.UnlimitedTracking
		m.Data.TrackingPosition = map3955.Data.TrackingPosition
		m.Data.Locked = map3955.Data.Locked
		m.Data.Colors = map3955.Data.Colors
		m.DataVersion = LatestDataVersion

	} else if v.DataVersion <= 4440 /* 1.21.8 */ {
		var map4325 Map4325
		_, err = nbt.NewDecoder(bytes.NewReader(b)).Decode(&map4325)
		m.Data.Width = 128
		m.Data.Height = 128
		m.Data.UuidMost = 0
		m.Data.UuidLeast = 0
		m.Data.XCenter = map4325.Data.XCenter
		m.Data.ZCenter = map4325.Data.ZCenter
		m.Data.Scale = 0
		m.Data.Dimension = map4325.Data.Dimension
		m.Data.UnlimitedTracking = 0
		m.Data.TrackingPosition = 0
		m.Data.Locked = 0
		m.Data.Colors = map4325.Data.Colors
		m.DataVersion = LatestDataVersion

	} else {
		err = fmt.Errorf("unsupported data version %v", v.DataVersion)
	}

	if len(m.Data.Colors) == 0 {
		err = fmt.Errorf("no colors, probably not a map")
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
