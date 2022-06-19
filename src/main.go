package main

func main() {
	profiler := create_profiler()

	camera := create_camera(Point3{X:0, Y:0, Z:0}, 1280, 720)
	image := create_image(camera.Width, camera.Height)

  world := []Sphere {
    Sphere{Center:Point3{X:-0.8, Y:-0.2, Z:-1}, Radius:0.3, Albedo:Vector3{0.2, 0.3, 0.8}},
    Sphere{Center:Point3{X:0, Y:0, Z:-1}, Radius:0.5, Albedo:Vector3{0.9, 0.0, 0.1}},
    Sphere{Center:Point3{X:0, Y:-100.5, Z:-1}, Radius:100, Albedo:Vector3{0.5, 0.5, 0.5}},
  };

  renderer := Renderer{Cam:camera, Img:image, ThreadsCount:40, SamplesPerPixel: 100}

	profiler.Start("Rendering")
	renderer.Prepare()
  renderer.Render(world)
	profiler.End("Rendering")

	renderer.Shutdown()
}
