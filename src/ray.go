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

func get_ray_color(ray *Ray, world []Sphere) color.RGBA {
  var hit_record HitRecord
  did_it_hit := false
  closest_so_far := math.Inf(1)
  
  for i := 0; i < len(world); i++ {
    if world[i].hit(ray, 0, closest_so_far, &hit_record) {
      did_it_hit = true
      closest_so_far = hit_record.T;
    }
  }

  if did_it_hit {
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
