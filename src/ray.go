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

  if !did_it_hit {
    return Vector3{0, 0, 0} // background
  }
    
  var scattered Ray
  var attenuation Vector3
  emitted := hit_record.Mat.emit() 
   
  if hit_record.Mat.scatter(ray, hit_record, &attenuation, &scattered) {
    return vectors_add(emitted, vectors_multiply(get_ray_color(&scattered, world, depth-1), attenuation))
  }

  return emitted
}
