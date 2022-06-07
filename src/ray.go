package main

import "image/color"

func point_at(ray Ray, scalar float64) Point3 { 
    // origin + direction * scalar
    ray.Scalar = scalar
    ds := ray.Direction.Multiply(ray.Scalar)
    return ray.Origin.Add(&ds);
}

func get_ray_color(ray *Ray) color.RGBA {
	unit_dir := ray.Direction.Unit()
	t := 0.5 * (unit_dir.Y + 1.0)

	white := Vector3{X:1.0, Y:1.0, Z:1.0}
	blue := Vector3{X:0.5, Y:0.7, Z:1.0}
	c := vectors_add(white.Multiply(t), blue.Multiply(1.0-t))

	return color.RGBA{R:uint8(255*c.X), G:uint8(255*c.Y), B:uint8(255*c.Z), A:255}
}