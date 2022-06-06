package main

import "math"

type Vector3 struct {
    X, Y, Z float64
}

type Point3 = Vector3

type Ray struct {
    Origin Point3
    Direction Vector3
    Scalar float64
}

func add_vectors(p1 Point3, p2 Point3) Point3 {
    result := Point3{X:0, Y:0, Z:0}

    result.X = p1.X + p2.X
    result.Y = p1.Y + p2.Y
    result.Z = p1.Z + p2.Z

    return result
}

func multiply_vector(v Vector3, t float64) Vector3 {
    result := v

    result.X *= t
    result.Y *= t
    result.Z *= t

    return result
}

func divide_vector(v Vector3, t float64) Vector3 {
    result := v

    result.X /= t
    result.Y /= t
    result.Z /= t

    return result
}

func substract_vectors(v1 Vector3, v2 Vector3) Vector3 {
    result := Vector3{X:0, Y:0, Z:0}

    result.X = v1.X - v2.X
    result.Y = v1.Y - v2.Y
    result.Z = v1.Z - v2.Z

    return result
}

func unit_vector(v Vector3) Vector3 {
    result := v
    length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

    result.X /= length
    result.Y /= length
    result.Z /= length

    return result
}

func point_at(ray Ray, scalar float64) Point3 {
    ray.Scalar = scalar
    return add_vectors(ray.Origin, multiply_vector(ray.Direction, ray.Scalar));
}