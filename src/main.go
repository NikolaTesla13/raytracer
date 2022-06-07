package main

func main() {
	profiler := create_profiler()

	camera := create_camera(Point3{X:0, Y:0, Z:0}, 1280, 720)
	image := create_image(camera.Width, camera.Height)

	renderer := Renderer{Cam:camera, Img:image, ThreadsCount:40}
	renderer.Prepare()

	profiler.Start("Rendering")
	renderer.Render()
	profiler.End("Rendering")

	renderer.Shutdown()
}
