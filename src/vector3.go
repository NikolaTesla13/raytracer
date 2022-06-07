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

func (v1 Vector3) Add(v2 *Vector3) Vector3 {
    result := Vector3{X:0, Y:0, Z:0}

    result.X = v1.X + v2.X
    result.Y = v1.Y + v2.Y
    result.Z = v1.Z + v2.Z

    return result
}

func vectors_add(v1 Vector3, v2 Vector3) Vector3 {
    result := Vector3{X:0, Y:0, Z:0}

    result.X = v1.X + v2.X
    result.Y = v1.Y + v2.Y
    result.Z = v1.Z + v2.Z

    return result
}

func (v Vector3) Multiply(t float64) Vector3 {
    result := Vector3{X:v.X, Y:v.Y, Z:v.Z}

    result.X *= t
    result.Y *= t
    result.Z *= t

    return result
}

func (v Vector3) Divide(t float64) Vector3 {
    result := Vector3{X:v.X, Y:v.Y, Z:v.Z}

    result.X /= t
    result.Y /= t
    result.Z /= t

    return result
}

func (v1 Vector3) Substract(v2 *Vector3) Vector3 {
    result := Vector3{X:v1.X, Y:v1.Y, Z:v1.Z}

    result.X -= v2.X
    result.Y -= v2.Y
    result.Z -= v2.Z

    return result
}

func vectors_substract(v1 Vector3, v2 Vector3) Vector3 {
    result := Vector3{X:v1.X, Y:v1.Y, Z:v1.Z}

    result.X -= v2.X
    result.Y -= v2.Y
    result.Z -= v2.Z

    return result
}

func (v Vector3) Unit() Vector3 {
    result := Vector3{X:v.X, Y:v.Y, Z:v.Z}
    length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

    result.X /= length
    result.Y /= length
    result.Z /= length

    return result
}

func (v Vector3) LengthSquared() float64 {
    return v.X * v.X + v.Y * v.Y + v.Z * v.Z;
}

func dot(v1 Vector3, v2 Vector3) float64 {
    return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z; 
}

func cross(v1 Vector3, v2 Vector3) Vector3 {
    return Vector3{X:v1.Y*v2.Z-v1.Z*v2.Y, Y:v1.Z*v2.X-v1.X*v2.Z, Z:v1.X*v2.Y-v1.Y*v2.X};
}