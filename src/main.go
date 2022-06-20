package main

func main() {
	profiler := create_profiler()

  camera := create_camera(Point3{X:0, Y:0, Z:1.5}, Point3{X:0, Y:0, Z:0}, 30.0, 1280, 720) // look from, look at, FOV, width and height
	image := create_image(camera.Width, camera.Height)

  red_diff := Diffuse{Albedo:Vector3{1.0, 0.1, 0.2}}
  ground := Diffuse{Albedo:Vector3{0.1, 0.1, 0.1}}
  blue_light := Light{Emit:Vector3{0.2, 0.3, 0.8}}
  metal_green := Light{Emit:Vector3{0.1, 0.8, 0.3}}
  sunshine := Metal{Albedo:Vector3{0.7, 0.8, 0.1}, Fuzz:0.8}

  world := []Sphere {
    Sphere{Center:Point3{X:-1.5, Y:0.0, Z:-1.0}, Radius:0.5, Mat:&blue_light},
    Sphere{Center:Point3{X:-0.35, Y:0.0, Z:-1.0}, Radius:0.5, Mat:&red_diff},
    Sphere{Center:Point3{X:0.6, Y:-0.1, Z:-1.0}, Radius:0.4, Mat:&metal_green},
    Sphere{Center:Point3{X:1.4, Y:-0.2, Z:-1.0}, Radius:0.3, Mat:&sunshine},
    Sphere{Center:Point3{X:0, Y:-100.5, Z:-1.0}, Radius:100, Mat:&ground},
  };

  renderer := Renderer{Cam:camera, Img:image, ThreadsCount:40, SamplesPerPixel: 500, MaxDepth: 500}

	profiler.Start("Rendering")
	renderer.Prepare()
  renderer.Render(world)
	profiler.End("Rendering")

	renderer.Shutdown()
}
