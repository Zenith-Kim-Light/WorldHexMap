package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Zenith-Kim-Light/WorldHexMap/internal"
	"github.com/Zenith-Kim-Light/WorldHexMap/reylib/gui"
	"github.com/Zenith-Kim-Light/WorldHexMap/reylib/util"
	"github.com/Zenith-Kim-Light/hexgrid"
	rlgui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const windowWidthDefault = 800
const windowHeightDefault = 600

func init() {
	runtime.LockOSThread()
}

func main() {
	windowHeight := int32(windowHeightDefault)
	windowWidth := int32(windowWidthDefault)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagWindowMinimized)

	buttons := initButtons()

	rl.InitWindow(windowWidth, windowHeight, "WatabouMap editor")
	rl.SetTargetFPS(60)

	rl.SetWindowFocused()

	defer rl.CloseWindow()

	var regionMap *internal.RegionMap

	cam := util.InitCamera()

	rl.SetTargetFPS(60)

	var i float32 = 0.0

	var rects = initMenuRects(buttons)
	var rt2d rl.RenderTexture2D
	var textureRectangle rl.Rectangle
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
					textureWidth, textureHeight := util.CalcNearestPowersOfTwo(regionMap.WatabouMap.Bp.Width, regionMap.WatabouMap.Bp.Height)
					rt2d = rl.LoadRenderTexture(textureWidth, textureHeight)
					rl.SetTextureFilter(rt2d.Texture, rl.FilterTrilinear)
					rl.BeginTextureMode(rt2d)
					rl.ClearBackground(rl.RayWhite)
					drawMap(regionMap)
					rl.EndTextureMode()
					rl.GenTextureMipmaps(&rt2d.Texture)
					textureRectangle = rl.NewRectangle(0, 0, float32(textureWidth), -float32(textureHeight))
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

		rl.ClearBackground(rl.RayWhite)
		cam.Begin()

		if rt2d.ID != 0 {
			rl.DrawTextureRec(rt2d.Texture, textureRectangle, rl.Vector2{}, rl.White)
		}
		rl.DrawCircleLines(
			int32(cam.Camera2D.Target.X),
			int32(cam.Camera2D.Target.Y),
			5.,
			rl.Red,
		)
		cam.End()

		for i, rect := range rects {
			buttonInfo := buttons[i]
			button := rlgui.Button(rect, buttonInfo.Label)
			if button {
				handleButton(gui.ButtonAction(i))
			}
		}

		rl.EndDrawing()
		i += 1.0
	}
}

func initButtons() []gui.Button {
	return []gui.Button{
		{ID: gui.ActionField, Label: "Add Field"},
		{ID: gui.ActionForest, Label: "Add Forest"},
		{ID: gui.ActionMount, Label: "Add Mount"},
		{ID: gui.ActionWater, Label: "Add Water"},
		{ID: gui.ActionExit, Label: "EXIT"},
	}
}

func handleButton(buttonId gui.ButtonAction) {
	switch buttonId {
	case gui.ActionExit:
		os.Exit(0)
	default:
		fmt.Print("Пу")
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

func initMenuRects(buttons []gui.Button) []rl.Rectangle {
	screenHeight := float32(rl.GetScreenHeight())
	sidebarWidth := float32(220)
	sidebarRect := rl.NewRectangle(0, 0, sidebarWidth, screenHeight)
	rlgui.Panel(sidebarRect, "Menu")

	var (
		btnPaddingTop  float32 = 1
		btnPaddingLeft float32 = 1
		btnHeight      float32 = 30
		btnSpacing     float32 = 1
		btnWidth       float32 = sidebarWidth - (btnPaddingLeft * 2)
	)

	var rectangles []rl.Rectangle

	for i := range buttons {
		btnY := btnPaddingTop + float32(i)*(btnHeight+btnSpacing)
		btnRect := rl.NewRectangle(btnPaddingLeft, btnY, btnWidth, btnHeight)
		rectangles = append(rectangles, btnRect)
	}

	return rectangles
}
