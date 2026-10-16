package internal

import (
	"github.com/Zenith-Kim-Light/WorldHexMap/internal/watabou_map"
	"github.com/Zenith-Kim-Light/hexgrid"
)

type RegionMap struct {
	// Карта в формате Watabou. Выполняет роль контейнера данных
	WatabouMap *watabou_map.WatabouMap

	// Костыль. Виртуальный слой гексов. Нужно для вычисления и преобразования координат
	Layout hexgrid.Layout
	// Костыль. Конвертер для получения осевых координат
	converterToAxial func(hex hexgrid.Hex) hexgrid.Hex
	// Костыль. Ориентация гекса
	HexOrientation string
}

func NewRegionMap(wm *watabou_map.WatabouMap) *RegionMap {
	origin := hexgrid.Point{X: 0, Y: 0}
	size := hexgrid.Point{X: 20, Y: 20}

	var layout hexgrid.Layout
	var f func(hex hexgrid.Hex) hexgrid.Hex
	var hexOrientation string
	if watabou_map.LayoutEvenQ == wm.Layout {
		layout = hexgrid.Layout{Orientation: hexgrid.OrientationFlat, Size: size, Origin: origin}
		hexOrientation = "OrientationFlat"
		f = evenQToAxial
	} else {
		layout = hexgrid.Layout{Orientation: hexgrid.OrientationPointy, Size: size, Origin: origin}
		hexOrientation = "OrientationPointy"
		f = oddRToAxial
	}

	return &RegionMap{
		WatabouMap:       wm,
		Layout:           layout,
		converterToAxial: f,
		HexOrientation:   hexOrientation,
	}
}

// ConvertHexToAxial Конвертирование гексагона в осевые координаты его центра
func (RegionMap RegionMap) ConvertHexToAxial(hex hexgrid.Hex) hexgrid.Point {
	axialHex := RegionMap.converterToAxial(hex)
	return hexgrid.HexToPixel(RegionMap.Layout, axialHex)
}

func evenQToAxial(hex hexgrid.Hex) hexgrid.Hex {
	parity := hex.Q & 1
	var q = hex.Q
	var r = hex.R - (hex.Q+parity)/2
	return hexgrid.NewHex(q, r)
}

func oddRToAxial(hex hexgrid.Hex) hexgrid.Hex {
	parity := hex.R & 1
	var q = hex.Q - (hex.R-parity)/2
	var r = hex.R
	return hexgrid.NewHex(q, r)
}
