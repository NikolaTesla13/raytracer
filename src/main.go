package main

func main() {
	profiler := create_profiler()

	camera := create_camera(Point3{X:0, Y:0, Z:0}, 1280, 720)
	image := create_image(camera.Width, camera.Height)

  center := Diffuse{Albedo:Vector3{0.7, 0.1, 0.2}}
  ground := Diffuse{Albedo:Vector3{0.1, 0.1, 0.1}}

  world := []Sphere {
    //Sphere{Center:Point3{X:-0.4, Y:-0.3, Z:-1}, Radius:0.4, Albedo:Vector3{0.2, 0.3, 0.8}},
    Sphere{Center:Point3{X:0, Y:0.0, Z:-1.0}, Radius:0.5, Mat:&center},
    Sphere{Center:Point3{X:0, Y:-100.5, Z:-1}, Radius:100, Mat:&ground},
  };

  renderer := Renderer{Cam:camera, Img:image, ThreadsCount:40, SamplesPerPixel: 10}

	profiler.Start("Rendering")
	renderer.Prepare()
  renderer.Render(world)
	profiler.End("Rendering")

	renderer.Shutdown()
}
