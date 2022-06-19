package main

type Material interface {
  scatter(r_in *Ray, hit_record HitRecord, attenuation *Vector3, scattered *Ray) bool
}

type Diffuse struct {
  Albedo Vector3
}

func (d *Diffuse) scatter(r_in *Ray, hit_record HitRecord, attenuation *Vector3, scattered *Ray) bool {
  scatter_direction := vectors_add(hit_record.Normal, rand_unit_vec())
  
  if is_near_zero(&scatter_direction) {
    scatter_direction = hit_record.Normal
  }

  *scattered = Ray{hit_record.Intersec, scatter_direction, 0}
  *attenuation = d.Albedo
  return true
}

