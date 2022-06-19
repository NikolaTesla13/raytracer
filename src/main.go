package main

func main() {
	profiler := create_profiler()

	camera := create_camera(Point3{X:0, Y:0, Z:0}, 1280, 720)
	image := create_image(camera.Width, camera.Height)

  world := []Sphere { 
    Sphere{Center:Point3{X:0, Y:0, Z:-1}, Radius:0.5},
    Sphere{Center:Point3{X:0, Y:100.5, Z:-1}, Radius:100},
  };

	renderer := Renderer{Cam:camera, Img:image, ThreadsCount:40}
	renderer.Prepare()

	profiler.Start("Rendering")
	renderer.Render(world)
	profiler.End("Rendering")

	renderer.Shutdown()
}
