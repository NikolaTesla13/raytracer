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

func create_image(width int, height int) Image {
	var img Image

	img.Width = width
	img.Height = height
	img.Target = image.NewRGBA(image.Rect(0, 0, width, height))

	return img
}

func set_pixel(img Image, x int, y int, color color.RGBA) {
	img.Target.SetRGBA(x, y, color)
}

func write_image(img Image) {
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