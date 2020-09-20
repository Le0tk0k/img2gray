package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

var rm = flag.Bool("r", false, "Remove sorce file")

func main() {
	flag.Parse()
	src, err := os.Open(flag.Arg(0))
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

	if *rm {
		err = os.Remove(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}
}
