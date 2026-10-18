package main

import (
	"fmt"
	"runtime"

	"github.com/Zenith-Kim-Light/WorldHexMap/internal"
	"github.com/Zenith-Kim-Light/WorldHexMap/reylib/util"
	"github.com/Zenith-Kim-Light/hexgrid"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidthDefault = 800
const screenHeightDefault = 600

func init() {
	runtime.LockOSThread()
}

func main() {
	screenHeight := int32(screenHeightDefault)
	screenWidth := int32(screenWidthDefault)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagWindowMinimized)

	rl.InitWindow(screenWidth, screenHeight, "WatabouMap editor")
	rl.SetTargetFPS(60)

	rl.SetWindowFocused()

	defer rl.CloseWindow()

	var regionMap *internal.RegionMap

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
					regionMap = internal.NewRegionMap(wm)
				}
			}
		}

		zoomDirection := util.GetZoomDirectionByWheel()
		if zoomDirection != util.ZoomNone {
			cam.Zooming(zoomDirection)
		}

		clickPosition := util.GetMouseLeftClickPosition()
		if clickPosition != nil {

		}

		rl.DrawCircleLines(
			int32(cam.Camera2D.Offset.X),
			int32(cam.Camera2D.Offset.Y),
			5.,
			rl.Red,
		)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		cam.Begin()

		if nil != regionMap {
			drawMap(regionMap)
		}
		rl.EndMode2D()
		rl.EndDrawing()
		i += 1.0
	}
}

func drawMap(regionMap *internal.RegionMap) {
	// todo проверить как работает ориентация гекса в hexgrid
	var rotation float32
	if regionMap.HexOrientation == "OrientationFlat" {
		rotation = 0.0
	} else {
		rotation = 90.0
	}

	for _, wmHex := range regionMap.WatabouMap.Hexes {
		hex := hexgrid.NewHex(wmHex.Q, wmHex.R)
		point := regionMap.ConvertHexToAxial(hex)

		vV := rl.NewVector2(float32(point.X), float32(point.Y))
		colorRGBA_V := rl.DarkPurple
		rl.DrawPolyLines(vV, 6, 20, rotation, colorRGBA_V)
		s := fmt.Sprintf("Q:%d R:%d", hex.Q, hex.R)
		rl.DrawText(s, int32(point.X-10), int32(point.Y-10), 10, rl.Black)
	}
}
