package util

import rl "github.com/gen2brain/raylib-go/raylib"

func ResizeWindow() {

}

type Cam struct {
	Camera2D rl.Camera2D
	zoom     Zoom
}

func InitCamera() Cam {
	z := InitZoom()
	return Cam{
		Camera2D: rl.Camera2D{
			Target:   rl.NewVector2(0, 0),
			Offset:   rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
			Rotation: 0.0,
			Zoom:     1.0, // Масштаб 1:1
		},
		zoom: z,
	}
}

func (cam *Cam) Navigate2D() {
	camera := &cam.Camera2D
	dt := rl.GetFrameTime()
	const speed = 300.0
	moveStep := speed * dt

	if rl.IsKeyDown(rl.KeyRight) {
		camera.Target.X += moveStep
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		camera.Target.X -= moveStep
	}
	if rl.IsKeyDown(rl.KeyDown) {
		camera.Target.Y += moveStep
	}
	if rl.IsKeyDown(rl.KeyUp) {
		camera.Target.Y -= moveStep
	}
}

func (cam *Cam) Begin() {
	rl.BeginMode2D(cam.Camera2D)
}

func (cam *Cam) Zooming(zd ZoomDirection) {
	cam.zoom.Scale(zd)
	s := cam.zoom.getScale()
	cam.Camera2D.Zoom = float32(s)
}
