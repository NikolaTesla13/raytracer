package main

import (
    "image/color"
)

func point_at(ray *Ray, scalar float64) Point3 { 
    // origin + direction * scalar
    ray.Scalar = scalar
    return vectors_add(ray.Origin, ray.Direction.Multiply(ray.Scalar))
}

func get_ray_color(ray *Ray) color.RGBA {
    sphere := Sphere{Center:Point3{X:0, Y:0, Z:-1}, Radius:0.5}
    var hit_record HitRecord

    if sphere.hit(ray, 0.0, 1.0, &hit_record) {
        c := Vector3{X:hit_record.Normal.X+1, Y:hit_record.Normal.Y+1, Z:hit_record.Normal.Z+1}.Multiply(0.5)
    	return color.RGBA{R:uint8(255*c.X), G:uint8(255*c.Y), B:uint8(255*c.Z), A:255}
    }

	unit_dir := ray.Direction.Unit()
	t := 0.5 * (unit_dir.Y + 1.0)

	white := Vector3{X:1.0, Y:1.0, Z:1.0}
	blue := Vector3{X:0.5, Y:0.7, Z:1.0}
	c := vectors_add(white.Multiply(t), blue.Multiply(1.0-t))

	return color.RGBA{R:uint8(255*c.X), G:uint8(255*c.Y), B:uint8(255*c.Z), A:255}
}