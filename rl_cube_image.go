package rl_ci

import rl "github.com/gen2brain/raylib-go/raylib"

type CubeImage struct {
	Texture       rl.Texture2D
	Position      rl.Vector3
	RotationAxis  rl.Vector3
	RotationAngle float32
	Scale         rl.Vector3
	Color         rl.Color
	Model         rl.Model
}

func NewCubeImage(image *rl.Image, position rl.Vector3, scale rl.Vector3, color rl.Color) CubeImage {
	cube_mesh := rl.GenMeshCube(1., 1., 1.)
	model := rl.LoadModelFromMesh(cube_mesh)
	rl.ImageFlipVertical(image)
	texture := rl.LoadTextureFromImage(image)
	model.GetMaterials()[0].Maps.Texture = texture

	cube := CubeImage{
		Position:      position,
		RotationAxis:  rl.NewVector3(0.0, 0.0, 0.0),
		RotationAngle: 0.,
		Scale:         scale,
		Color:         color,
		Model:         model,
	}

	rl.ImageFlipVertical(image)

	return cube
}

func (cube CubeImage) DrawCubeImage() {
	rl.DrawModelEx(cube.Model, cube.Position, cube.RotationAxis, cube.RotationAngle, cube.Scale, cube.Color)
}

func (cube *CubeImage) Rotate(axis rl.Vector3, angle float32) {
	cube.RotationAxis = axis
	cube.RotationAngle = angle
}
