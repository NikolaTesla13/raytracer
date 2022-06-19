package main

import "math"

type Sphere struct {
	Center Point3
	Radius float64
  Mat Material
}

func (sphere *Sphere) hit(ray *Ray, t_min float64, t_max float64, rec *HitRecord) bool {
  oc := vectors_substract(ray.Origin, sphere.Center)

  a := ray.Direction.LengthSquared()
  half_b := dot(oc, ray.Direction)
  c := oc.LengthSquared() - sphere.Radius * sphere.Radius

  discriminant := half_b*half_b - a*c

	if discriminant < 0 {
    return false
  }

	sqrtd := math.Sqrt(discriminant)
	root := (-half_b-sqrtd)/a
	if root < t_min || t_max < root {
		root = (-half_b+sqrtd)/a
		if root < t_min || t_max < root {
			return false
		}
	}

	rec.T = root
	rec.Intersec = point_at(ray, root)
	outward_normal := vectors_substract(rec.Intersec, sphere.Center).Divide(sphere.Radius)
	rec.SetFaceNormal(ray, outward_normal)
  rec.Mat = sphere.Mat

	return true
}
