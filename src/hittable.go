package main

type HitRecord struct {
	Intersec Point3
	Normal Vector3
	T float64
	IsFrontFace bool
  Albedo Vector3
}

func (rec *HitRecord) SetFaceNormal(r *Ray, outward_normal Vector3) {
	rec.IsFrontFace = dot(r.Direction, outward_normal) < 0
	if rec.IsFrontFace {
		rec.Normal = outward_normal
	} else {
		rec.Normal = outward_normal.Multiply(-1)
	}
}

type Hittable interface {
	hit(ray *Ray, t_min float64, t_max float64, rec *HitRecord) bool
}
