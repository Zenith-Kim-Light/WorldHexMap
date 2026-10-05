package util

import (
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func ProcessMouseClick(cam Cam, tiles map[string]*TileType) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		clickPos := rl.GetMousePosition()
		worldMousePos := rl.GetScreenToWorld2D(clickPos, cam.Camera2D)
		tileClickX := int32(worldMousePos.X) / TileSize
		tileClickY := int32(worldMousePos.Y) / TileSize

		tileClickedKey := fmt.Sprintf("X:%d,Y:%d", tileClickX, tileClickY)
		if tileClicked, ok := tiles[tileClickedKey]; ok {
			tileClicked.ShiftColor()
		}
		fmt.Println(strconv.Itoa(int(tileClickX)), int(tileClickY))
	}
}

//func ProcessMouseClickOnHex(cam Cam, layout hexgrid.Layout) {
//	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
//		clickPos := rl.GetMousePosition()
//		worldMousePos := rl.GetScreenToWorld2D(clickPos, cam.Camera2D)
//		p := hexgrid.Point{float64(worldMousePos.X), float64(worldMousePos.Y)}
//		h := hexgrid.PixelToHex(layout, p)
//
//		fmt.Println(strconv.Itoa(int(h.Q)), int(h.R))
//	}
//}
