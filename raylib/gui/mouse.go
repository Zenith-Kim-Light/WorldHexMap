package gui

import (
	"fmt"
	"strconv"

	"github.com/Zenith-Kim-Light/hexgrid"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetZoomDirectionByWheel() ZoomDirection {
	return ZoomDirection(rl.GetMouseWheelMove())
}

func GetMouseLeftClickPosition() *Click {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		vector := rl.GetMousePosition()
		return &Click{button: int32(rl.MouseButtonLeft), pos: vector}
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		vector := rl.GetMousePosition()
		return &Click{button: int32(rl.MouseButtonRight), pos: vector}
	}
	return nil
}

func ProcessMouseClickOnHex(cam Cam, layout hexgrid.Layout) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		clickPos := rl.GetMousePosition()
		worldMousePos := rl.GetScreenToWorld2D(clickPos, cam.Camera2D)
		p := hexgrid.Point{float64(worldMousePos.X), float64(worldMousePos.Y)}
		h := hexgrid.PixelToHex(layout, p)

		fmt.Println(strconv.Itoa(int(h.Q)), int(h.R))
	}
}
