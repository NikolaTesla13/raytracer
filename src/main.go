package main

import (
    "image/color"
)

func get_ray_color(ray *Ray) color.RGBA {
	unit_dir := ray.Direction.Unit()
	t := 0.5 * (unit_dir.Y + 1.0)

	white := Vector3{X:1.0, Y:1.0, Z:1.0}
	blue := Vector3{X:0.5, Y:0.7, Z:1.0}
	c := vectors_add(blue.Multiply(t), white.Multiply(1.0-t))

	return color.RGBA{R:uint8(255*c.X), G:uint8(255*c.Y), B:uint8(255*c.Z), A:255}
}

func main() {
	origin := Point3{X:0, Y:0, Z:0}
	camera := create_camera(origin, 1280, 720)
	img := create_image(camera.Width, camera.Height)
	profiler := create_profiler()

	profiler.Start("Rendering")
	for i := 0; i<=int(camera.Width); i++ {
		for j := 0; j<=int(camera.Height); j++ {
			u, v := camera.GetCoords(i, j)

			direction := camera.GetDirection(u, v)
 			ray := Ray{Origin:origin, Direction:direction, Scalar:0}

			img.SetPixel(i, j, get_ray_color(&ray))
		}
	}
	profiler.End("Rendering")

	img.Write()
}
