package rl_ci

import rl "github.com/gen2brain/raylib-go/raylib"

// CubeImage is a cube with an image
type CubeImage struct {
	Position      rl.Vector3
	RotationAxis  rl.Vector3
	RotationAngle float32
	Scale         rl.Vector3
	Color         rl.Color
	Model         rl.Model
}

// NewCubeImage creates a new cube with an image
func NewCubeImage(image *rl.Image, position rl.Vector3, scale rl.Vector3, color rl.Color) CubeImage {
	// Create a cube mesh
	cube_mesh := rl.GenMeshCube(1., 1., 1.)
	// Load the cube mesh into a model
	model := rl.LoadModelFromMesh(cube_mesh)
	rl.ImageFlipVertical(image)
	// Load the image into a texture
	texture := rl.LoadTextureFromImage(image)
	// Set the texture to the model
	model.GetMaterials()[0].Maps.Texture = texture

	// Create the cube
	cube := CubeImage{
		Position:      position,
		RotationAxis:  rl.NewVector3(0.0, 0.0, 0.0),
		RotationAngle: 0.,
		Scale:         scale,
		Color:         color,
		Model:         model,
	}

	// Flip the image back
	rl.ImageFlipVertical(image)

	// Return the cube
	return cube
}

// NewCubeImageEx creates a new cube with an image with extra parameters for rotation
func NewCubeImageEx(image *rl.Image, position rl.Vector3, scale rl.Vector3, rotaton_axis rl.Vector3, rotation_angle float32, color rl.Color) CubeImage {
	// Create the cube
	cube := NewCubeImage(image, position, scale, color)
	// Set the extra parameters
	cube.RotationAxis = rotaton_axis
	cube.RotationAngle = rotation_angle

	// Return the cube
	return cube
}

// DrawCubeImage draws the cube with an image
func (cube CubeImage) DrawCubeImage() {
	rl.DrawModelEx(cube.Model, cube.Position, cube.RotationAxis, cube.RotationAngle, cube.Scale, cube.Color)
}

// Rotate rotates the cube
func (cube *CubeImage) Rotate(axis rl.Vector3, angle float32) {
	cube.RotationAxis = axis
	cube.RotationAngle = angle
}
