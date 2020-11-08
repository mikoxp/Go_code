package main

import (
	"image"
	progressing "Go_code/src/image_progressing"
)

func main() {
	path := "image/lena.jpg"
	pathOUT:="image/image.jpg"
	rgba := progressing.OpenRGBAImage(path)

	width := rgba.Rect.Dx()
	height := rgba.Rect.Dy()
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	result := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	//cyan := color.RGBA{100, 200, 200, 0xff}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := rgba.RGBAAt(x, y)
			result.Set(x, y, p)
		}
	}
	progressing.SaveImage(pathOUT, result,"JPG")

}
