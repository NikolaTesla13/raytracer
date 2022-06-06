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

func create_camera(origin Vector3, width float64, height float64) Camera {
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
	cam.LowerLeftCorner = substract_vectors(substract_vectors(origin, add_vectors(divide_vector(cam.Horizontal, 2), divide_vector(cam.Vertical, 2))), Vector3{X:0, Y:0, Z:cam.FocalLength});

	return cam
}
