package main

import "math"

type Camera struct {
	Origin Point3

	Width float64
	Height float64
	
	AspectRatio float64
	ViewportHeight float64
	ViewportWidth float64
	FocalLength float64
  Fov float64

	Horizontal Vector3
	Vertical Vector3
	LowerLeftCorner Vector3
}

func create_camera(origin Point3, lookat Point3, fov float64, width float64, height float64) Camera {
	var cam Camera

	cam.Width = width
	cam.Height = height

  radians := fov * math.Pi / 90
  h := math.Tan(radians / 2)

	cam.AspectRatio = float64(width)/float64(height)
	cam.ViewportHeight = 2.0 * h
	cam.ViewportWidth = cam.AspectRatio * cam.ViewportHeight
	cam.FocalLength = 1.0

  w := vectors_substract(origin, lookat).Unit()
  u := cross(Vector3{0, 1, 0}, w).Unit()
  v := cross(w, u)

	cam.Origin = origin
	cam.Horizontal = vectors_multiply(Vector3{X:cam.ViewportWidth, Y:0, Z:0}, u)
	cam.Vertical = vectors_multiply(Vector3{X:0, Y:cam.ViewportHeight, Z:0}, v)

	cam.LowerLeftCorner = vectors_substract(vectors_substract(origin, vectors_add(cam.Horizontal.Divide(2), cam.Vertical.Divide(2))), w)

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
