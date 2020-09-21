package img2gray

import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func removeSrc(src string) error {
	if err := os.Remove(src); err != nil {
		return err
	}
	return nil
}

func ToGray(src, dst string, rmsrc bool) error {
	flag.Parse()
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
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

	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer df.Close()

	switch filepath.Ext(src) {
	case ".png":
		err = png.Encode(df, grayImg)
	case ".jpeg", ".jpg":
		err = jpeg.Encode(df, grayImg, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}
	if err != nil {
		return nil
	}

	if rmsrc {
		if err = removeSrc(src); err != nil {
			return err
		}
	}
	return nil
}
