package main

import (
    "image"
    "image/draw"
    "image/png"
    "image/color"
    "os"
    "bufio"
    "bytes"
    "log"
	"fmt"
	"time"
)

func main() {
	width := 1280
	height := 720

    target := image.NewRGBA(image.Rect(0, 0, width, height))
    draw.Draw(target, target.Bounds(), image.White, image.ZP, draw.Src)

	radius := 200
	x := width/2
	y := height/2

	// placeholder code, not actual ray tracing yet.
	for i := 0; i<=width; i++ {
		for j := 0; j<=height; j++ {
			if (i-x)*(i-x) + (j-y)*(j-y) <= radius*radius {
				target.Set(i, j, color.RGBA{240, 70, 123, 255})
			} else {
				target.SetRGBA(i, j, color.RGBA{105, 207, 255, 255})
			}
		}
	}

	enc := png.Encoder{
		CompressionLevel: png.NoCompression,
	}

    var imageBuf bytes.Buffer
    err := enc.Encode(&imageBuf, target)
    if err != nil {
        log.Panic(err)
    }

	filename := fmt.Sprintf("%v.png", time.Now().Format("02-01-2006"))
    fo, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    fw := bufio.NewWriter(fo)

    fw.Write(imageBuf.Bytes())
}