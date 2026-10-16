package main

import (
	"fmt"

	"github.com/Zenith-Kim-Light/WorldHexMap/internal/watabou_map"
	"github.com/Zenith-Kim-Light/WorldHexMap/reylib/util"
	"github.com/Zenith-Kim-Light/hexgrid"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	//runtime.LockOSThread()
	screenWidth := int32(800)
	screenHeight := int32(600)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagWindowMinimized)

	rl.InitWindow(screenWidth, screenHeight, "WatabouMap editor")
	rl.SetTargetFPS(60)

	//target := rl.LoadRenderTexture
	fmt.Println(rl.IsWindowResized())
	defer rl.CloseWindow()

	var WatabouMap *watabou_map.WatabouMap

	cam := util.InitCamera()

	rl.SetTargetFPS(60)

	var i float32 = 0.0
	for !rl.WindowShouldClose() {
		cam.Navigate2D()
		if rl.IsWindowResized() {
			util.ResizeWindow()
		}

		if rl.IsFileDropped() {
			paths := rl.LoadDroppedFiles()
			if len(paths) > 0 {
				wm, appErr := util.ProcessDroppedFile(paths[0])
				if appErr != nil {
					rl.TraceLog(rl.LogError, "Error occurred during process file: %v", appErr)
				} else {
					WatabouMap = wm
				}
			}
		}

		zoomDirection := util.GetZoomDirectionByWheel()
		if zoomDirection != util.ZoomNone {
			cam.Zooming(zoomDirection)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		cam.Begin()

		if nil != WatabouMap {
			fillWorld(WatabouMap)
		}
		rl.EndMode2D()
		rl.EndDrawing()
		i += 1.0
	}
}

func fillWorld(wm *watabou_map.WatabouMap) {
	origin := hexgrid.Point{X: 0, Y: 0}
	size := hexgrid.Point{X: 20, Y: 20}

	var layout hexgrid.Layout
	var r float32
	var f func(hex hexgrid.Hex) hexgrid.Hex
	if watabou_map.LayoutEvenQ == wm.Layout {
		layout = hexgrid.Layout{Orientation: hexgrid.OrientationFlat, Size: size, Origin: origin}
		r = 0.0
		f = evenq_to_axial
	} else {
		layout = hexgrid.Layout{Orientation: hexgrid.OrientationPointy, Size: size, Origin: origin}
		r = 90.0
		f = oddr_to_axial
	}

	for _, wmHex := range wm.Hexes {
		hex1 := hexgrid.NewHex(wmHex.Q, wmHex.R)
		drawHex(layout, hex1, r, f)
	}
}

func drawHex(layout hexgrid.Layout, hex hexgrid.Hex, r float32, f func(hex hexgrid.Hex) hexgrid.Hex) {
	aHex := f(hex)
	point := hexgrid.HexToPixel(layout, aHex)
	vV := rl.NewVector2(float32(point.X), float32(point.Y))

	colorRGBA_V := rl.DarkPurple
	rl.DrawPolyLines(vV, 6, 20, r, colorRGBA_V)
	s := fmt.Sprintf("Q:%d R:%d", hex.Q, hex.R)
	rl.DrawText(s, int32(point.X-10), int32(point.Y-10), 10, rl.Black)
}

func evenq_to_axial(hex hexgrid.Hex) hexgrid.Hex {
	parity := hex.Q & 1
	var q = hex.Q
	var r = hex.R - (hex.Q+parity)/2
	return hexgrid.NewHex(q, r)
}

func oddr_to_axial(hex hexgrid.Hex) hexgrid.Hex {
	parity := hex.R & 1
	var q = hex.Q - (hex.R-parity)/2
	var r = hex.R
	return hexgrid.NewHex(q, r)
}
