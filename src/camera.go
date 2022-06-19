package main

type Camera struct {
	Origin Point3

	Width float64
	Height float64
	
	AspectRatio float64
	ViewportHeight float64
	ViewportWidth float64
	FocalLength float64

	Horizontal Vector3
	Vertical Vector3
	LowerLeftCorner Vector3
}

func create_camera(origin Point3, width float64, height float64) Camera {
	var cam Camera

	cam.Width = width
	cam.Height = height

	cam.AspectRatio = float64(width)/float64(height)
	cam.ViewportHeight = 2.0
	cam.ViewportWidth = cam.AspectRatio * cam.ViewportHeight
	cam.FocalLength = 1.0

	cam.Origin = origin
	cam.Horizontal = Vector3{X:cam.ViewportWidth, Y:0, Z:0}
	cam.Vertical = Vector3{X:0, Y:cam.ViewportHeight, Z:0}

	cam.LowerLeftCorner = vectors_substract(vectors_substract(origin, vectors_add(cam.Horizontal.Divide(2), cam.Vertical.Divide(2))), Vector3{X:0, Y:0, Z:cam.FocalLength})

	return cam
}

func (cam *Camera) GetCoords(i int, j int) (float64, float64) {
	u := (float64(i)+rand_float(0.0, 1.0))/(cam.Width-1)
	v := (float64(j)+rand_float(0.0, 1.0))/(cam.Height-1)

	return u, v
}

func (cam *Camera) GetDirection(u float64, v float64) Vector3 {
	return vectors_substract(vectors_add(cam.LowerLeftCorner, vectors_add(cam.Horizontal.Multiply(u), cam.Vertical.Multiply(v))), cam.Origin)
}

func (cam *Camera) GetWidth() int {
	return int(cam.Width)
}

func (cam *Camera) GetHeight() int {
	return int(cam.Height)
}
