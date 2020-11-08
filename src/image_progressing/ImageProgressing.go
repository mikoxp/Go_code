package image_progressing

import (
	"image"
	"os"
	"image/jpeg"
	"image/draw"
	"image/png"
)
type Figure interface {
	field() float64
}
func OpenRGBAImage(path string) *image.RGBA {
	imgfile, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer imgfile.Close()

	img, err := jpeg.Decode(imgfile)
	if err != nil {
		panic(err.Error())
	}
	var rect=img.Bounds()
	result := image.NewRGBA(rect)
	draw.Draw(result, rect, img, rect.Min, draw.Src)
	return result
}

func SaveImage(path string, img *image.RGBA,fileType string)  {
	f, _ := os.Create(path)

	switch fileType {
	case "PNG":
		png.Encode(f, img)
	case "JPG":
		jpeg.Encode(f,img, &jpeg.Options{jpeg.DefaultQuality})
	}
}
