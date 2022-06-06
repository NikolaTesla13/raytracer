package main

import (
    "image/color"
    "time"
	"log"
)

func get_ray_color(ray Ray) color.RGBA { // fix this / refactor
	unit_dir := unit_vector(ray.Direction)
	t := 0.5 * (unit_dir.Y + 1.0)
	c := add_vectors(multiply_vector(Vector3{X:1.0, Y:1.0, Z:1.0}, (1.0-t)), multiply_vector(Vector3{X:0.5, Y:0.7, Z:1.0}, t))

	return color.RGBA{R:uint8(255*c.X), G:uint8(255*c.Y), B:uint8(255*c.Z), A:255}
}

func main() {
	camera := create_camera(Vector3{X:0, Y:0, Z:0}, 1280, 720)
	img := create_image(1280, 720)
	start := time.Now()

	for i := 0; i<=int(camera.Width); i++ {
		for j := 0; j<=int(camera.Height); j++ {
			u := float64(i)/(camera.Width-1)
			v := float64(j)/(camera.Height-1)

			origin := Point3{X:0, Y:0, Z:0}
			direction := substract_vectors(add_vectors(camera.LowerLeftCorner, add_vectors(multiply_vector(camera.Horizontal, u), multiply_vector(camera.Vertical, v))), camera.Origin)
			ray := Ray{Origin:origin, Direction:direction, Scalar:0} // refactor these lines

			set_pixel(img, i, j, get_ray_color(ray))
		}
	}

	elapsed := time.Since(start)
    log.Printf("\033[32mRendering took %s\033[97m", elapsed)

	write_image(img)
}
