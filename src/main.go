package main

func main() {
	profiler := create_profiler()

  camera := create_camera(Point3{X:0, Y:0, Z:1.5}, Point3{X:0, Y:0, Z:0}, 30.0, 1280, 720) // look from, look at, FOV, width and height
	image := create_image(camera.Width, camera.Height)

  red_diff := Diffuse{Albedo:Vector3{1.0, 0.1, 0.2}}
  ground := Diffuse{Albedo:Vector3{0.1, 0.1, 0.1}}
  fuzzy_metal := Metal{Albedo:Vector3{0.2, 0.3, 0.8}, Fuzz:0.4}
  metal_green := Metal{Albedo:Vector3{0.1, 0.8, 0.3}, Fuzz: 0.2}
  sunshine := Diffuse{Albedo:Vector3{0.7, 0.8, 0.1}}

  world := []Sphere {
    Sphere{Center:Point3{X:-1.5, Y:0.0, Z:-1.0}, Radius:0.5, Mat:&fuzzy_metal},
    Sphere{Center:Point3{X:-0.35, Y:0.0, Z:-1.0}, Radius:0.5, Mat:&red_diff},
    Sphere{Center:Point3{X:0.6, Y:-0.1, Z:-1.0}, Radius:0.4, Mat:&metal_green},
    Sphere{Center:Point3{X:1.4, Y:-0.2, Z:-1.0}, Radius:0.3, Mat:&sunshine},
    Sphere{Center:Point3{X:0, Y:-100.5, Z:-1.0}, Radius:100, Mat:&ground},
  };

  renderer := Renderer{Cam:camera, Img:image, ThreadsCount:40, SamplesPerPixel: 100}

	profiler.Start("Rendering")
	renderer.Prepare()
  renderer.Render(world)
	profiler.End("Rendering")

	renderer.Shutdown()
}
