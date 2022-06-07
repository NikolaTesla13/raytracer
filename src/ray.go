package main

import (
    "image/color"
    "math"
)

func point_at(ray *Ray, scalar float64) Point3 { 
    // origin + direction * scalar
    ray.Scalar = scalar
    return vectors_add(ray.Origin, ray.Direction.Multiply(ray.Scalar))
}

func hit_sphere(center Point3, radius float64, ray *Ray) float64 {
    oc := vectors_substract(ray.Origin, center)

    a := ray.Direction.LengthSquared()
    half_b := dot(oc, ray.Direction)
    c := oc.LengthSquared() - radius * radius

    discriminant := half_b*half_b - a*c

    if discriminant < 0 {
        return -1.0;
    } else {
        return (-half_b - math.Sqrt(discriminant)) / a;
    }
}

func get_ray_color(ray *Ray) color.RGBA {
    t := hit_sphere(Point3{X:0, Y:0, Z:-1}, 0.5, ray)
    if t > 0.0 {
        normal := vectors_substract(point_at(ray, t), Vector3{X:0.0, Y:0.0, Z:-1.0}).Unit()
        c := Vector3{X:normal.X+1, Y:normal.Y+1, Z:normal.Z+1}.Multiply(0.5)
    	return color.RGBA{R:uint8(255*c.X), G:uint8(255*c.Y), B:uint8(255*c.Z), A:255}
    }

	unit_dir := ray.Direction.Unit()
	t = 0.5 * (unit_dir.Y + 1.0)

	white := Vector3{X:1.0, Y:1.0, Z:1.0}
	blue := Vector3{X:0.5, Y:0.7, Z:1.0}
	c := vectors_add(white.Multiply(t), blue.Multiply(1.0-t))

	return color.RGBA{R:uint8(255*c.X), G:uint8(255*c.Y), B:uint8(255*c.Z), A:255}
}