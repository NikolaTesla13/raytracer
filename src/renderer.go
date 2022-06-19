package main

import "sync"

type Renderer struct {
	Cam Camera
	Img Image
	Wg sync.WaitGroup
	ThreadsCount int
}

func (renderer *Renderer) Prepare() {
	renderer.Wg.Add(renderer.Cam.GetWidth()/renderer.ThreadsCount+1)
}

func (renderer *Renderer) Render(world []Sphere) {
	for i := 0; i<=renderer.Cam.GetWidth(); i+=renderer.ThreadsCount {
		go func(i int, camera *Camera, img *Image) {
			defer renderer.Wg.Done()

			for k := 0; k<=renderer.ThreadsCount; k++ {
				for j := 0; j<=camera.GetHeight(); j++ {
					u, v := camera.GetCoords(i+k, j)
		
					direction := camera.GetDirection(u, v)
					ray := Ray{Origin:camera.Origin, Direction:direction, Scalar:0}
				
					img.SetPixel(i+k, j, get_ray_color(&ray, world))
				}	
			}
		}(i, &renderer.Cam, &renderer.Img)
	}
	renderer.Wg.Wait()
}

func (renderer *Renderer) Shutdown() {
	renderer.Img.Write()
}
