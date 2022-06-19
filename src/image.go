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

func (img *Image) SetPixel(x int, y int, color color.RGBA) {
	img.Target.SetRGBA(x, img.Height - y, color)
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
