package iimg

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
)
import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// ToRGB  convert image color to rgb
func ToRGB(img image.Image) image.Image {
	bounds := img.Bounds()
	rgbImage := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			r, g, b, _ := originalPixel.RGBA()
			rgbImage.Set(x, y, color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: 255,
			})
		}
	}
	return rgbImage
}

// ToJPG
func ToJPG(img image.Image, quality int) ([]byte, error) {
	b := new(bytes.Buffer)
	err := jpeg.Encode(b, img, &jpeg.Options{Quality: quality})
	return b.Bytes(), err
}
