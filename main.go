package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	grayImg := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			grayImg.Set(x, y, gray)
		}
	}

	dst, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	if err := png.Encode(dst, grayImg); err != nil {
		log.Fatal(err)
	}
}
