# Raylib Cube Image

Simple library that makes it possible to draw cubes with images.<hr>
Install with `go get github.com/antosmichael07/Raylib-Cube-Image`

## Example

```go
package main

import (
	rl_ci "github.com/antosmichael07/Raylib-Cube-Image"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(int32(rl.GetMonitorWidth(rl.GetCurrentMonitor())), int32(rl.GetMonitorHeight(rl.GetCurrentMonitor())), "WTF is this?")
	rl.ToggleFullscreen()
	rl.DisableCursor()
	rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))
	defer rl.CloseWindow()

	camera := rl.NewCamera3D(
		rl.NewVector3(10.0, 10.0, 10.0),
		rl.NewVector3(0.0, 0.0, 0.0),
		rl.NewVector3(0.0, 1.0, 0.0),
		70.0,
		rl.CameraPerspective,
	)

	// Here are the cubes initialized
	image := rl.LoadImage("image.png")
	cube_images := []rl_ci.CubeImage{
		rl_ci.NewCubeImage(image, rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector3(2.0, 5.0, 2.0), rl.Red),
		rl_ci.NewCubeImage(image, rl.NewVector3(3.0, 0.0, 0.0), rl.NewVector3(2.0, 3.0, 2.0), rl.White),
		rl_ci.NewCubeImage(image, rl.NewVector3(-3.0, 0.0, 0.0), rl.NewVector3(3.0, 1.0, 2.0), rl.White),
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.UpdateCamera(&camera, rl.CameraFree)
		rl.BeginMode3D(camera)

		// Here are the cubes drawn
		for i := range cube_images {
			cube_images[i].DrawCubeImage()
		}

		rl.EndMode3D()
		rl.EndDrawing()
	}
}
```
