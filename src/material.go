package main

type Material interface {
  scatter(r_in *Ray, hit_record HitRecord, attenuation *Vector3, scattered *Ray) bool
  emit() Vector3 
}

// diffuse
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

func (d *Diffuse) emit() Vector3 {
  return Vector3{0, 0, 0}
}

// metal
type Metal struct {
  Albedo Vector3
  Fuzz float64
}

func reflect(v Vector3, n Vector3) Vector3 {
  return vectors_substract(v, n.Multiply(dot(v, n)*2));
}

func (m *Metal) scatter(r_in *Ray, hit_record HitRecord, attenuation *Vector3, scattered *Ray) bool {
  reflected := reflect(r_in.Direction.Unit(), hit_record.Normal)
  *scattered = Ray{hit_record.Intersec, vectors_add(reflected, rand_in_unit_sphere().Multiply(m.Fuzz)), 0}
  *attenuation = m.Albedo
  return (dot(scattered.Direction, hit_record.Normal) > 0)
}

func (m *Metal) emit() Vector3 {
  return Vector3{0, 0, 0}
}

// light
type Light struct {
  Emit Vector3
}

func (l *Light) scatter(r_in *Ray, hit_record HitRecord, attenuation *Vector3, scattered *Ray) bool {
  return false
}

func (l *Light) emit() Vector3 {
  return l.Emit;
}
