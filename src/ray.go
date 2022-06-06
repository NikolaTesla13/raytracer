package main

func point_at(ray Ray, scalar float64) Point3 { 
    // origin + direction * scalar
    ray.Scalar = scalar
    ds := ray.Direction.Multiply(ray.Scalar)
    return ray.Origin.Add(&ds);
}