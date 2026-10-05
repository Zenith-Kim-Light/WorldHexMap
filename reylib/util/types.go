package util

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const TileSize = 32 // Размер одного блока в пикселях мира

type TileType struct {
	Title  rune
	WorldX int32
	WorldY int32
	IndexX int32
	IndexY int32
	Color  color.RGBA
}

func (t *TileType) ShiftColor() {
	switch t.Color {
	case rl.Gray:
		t.Color = rl.Gold
	case rl.Gold:
		t.Color = rl.Gray
	case rl.DarkGreen:
		t.Color = rl.DarkPurple
	case rl.DarkPurple:
		t.Color = rl.DarkGreen
	}
}
