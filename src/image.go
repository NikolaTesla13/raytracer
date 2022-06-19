package main

import (    
	"image"
	"image/png"
	"image/color"
	"os"
  "bufio"
  "bytes"
  "log"
  "fmt"
  "time"
  "math"
)

type Image struct {
	Target *image.RGBA
	Width int
	Height int
}

func create_image(width float64, height float64) Image {
	var img Image

	img.Width = int(width)
	img.Height = int(height)
	img.Target = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	return img
}

func (img *Image) SetPixel(x int, y int, c1 Vector3, samples int) {
  scale := 1.0 / float64(samples)
  c1.X = math.Sqrt(c1.X * scale)
  c1.Y = math.Sqrt(c1.Y * scale)
  c1.Z = math.Sqrt(c1.Z * scale)

  c1.X = clamp(c1.X, 0.0, 0.999)
  c1.Y = clamp(c1.Y, 0.0, 0.999)
  c1.Z = clamp(c1.Z, 0.0, 0.999)

  c := color.RGBA{R:uint8(255 * c1.X), G:uint8(255*c1.Y), B:uint8(255*c1.Z), A:255};
	img.Target.SetRGBA(x, img.Height - y, c)
}

func (img *Image) Write() {
	enc := png.Encoder{
		CompressionLevel: png.NoCompression,
	}

	var imageBuf bytes.Buffer
    err := enc.Encode(&imageBuf, img.Target)
    if err != nil {
		log.Panic(err)
    }

    filename := fmt.Sprintf("renders/%v.png", time.Now().Format("02-01-2006"))
    fo, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    fw := bufio.NewWriter(fo)

    fw.Write(imageBuf.Bytes())
}
