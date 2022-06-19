package main

import (
   "math"
)

func point_at(ray *Ray, scalar float64) Point3 { 
    // origin + direction * scalar
    ray.Scalar = scalar
    return vectors_add(ray.Origin, ray.Direction.Multiply(ray.Scalar))
}

func get_ray_color(ray *Ray, world []Sphere, depth int) Vector3 {
  if depth == 0 {
    return Vector3{0, 0, 0};
  }

  var hit_record HitRecord
  did_it_hit := false
  closest_so_far := math.Inf(1)
  
  for i := 0; i < len(world); i++ {
    if world[i].hit(ray, 0.001, closest_so_far, &hit_record) {
      did_it_hit = true
      closest_so_far = hit_record.T;
    }
  }

  if did_it_hit {
    var scattered Ray
    var attenuation Vector3
    
    if hit_record.Mat.scatter(ray, hit_record, &attenuation, &scattered) {
      return vectors_multiply(get_ray_color(&scattered, world, depth-1), attenuation)
    }
    return Vector3{0, 0, 0}
  }

  /*if did_it_hit {
    target := vectors_add(vectors_add(hit_record.Intersec, hit_record.Normal), rand_unit_vec())
    return vectors_multiply(get_ray_color(&Ray{hit_record.Intersec, vectors_substract(target, hit_record.Intersec), 0}, world, depth-1), Vector3{0.1, 0.1, 0.1});
  }*/

	unit_dir := ray.Direction.Unit()
	t := 0.5 * (unit_dir.Y + 1.0)

	white := Vector3{X:0.9, Y:0.9, Z:0.9}
	blue := Vector3{X:0.5, Y:0.7, Z:1.0}
	c := vectors_add(white.Multiply(1.0-t), blue.Multiply(t))

  return c
}
