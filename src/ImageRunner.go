package main

import (
	progressing "Go_code/src/image_progressing"
)

func main() {
	path := "image/lena.jpg"
	pathBW:="image/lena_black_white.jpg"
	pathSepia20:="image/lena_sepia20.jpg"
	pathSepia30:="image/lena_sepia30.jpg"
	pathSepia40:="image/lena_sepia40.jpg"
	rgba := progressing.OpenRGBAImage(path)

	//width := rgba.Rect.Dx()
	//height := rgba.Rect.Dy()
	//upLeft := image.Point{0, 0}
	//lowRight := image.Point{width, height}
	//result := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	//cyan := color.RGBA{100, 200, 200, 0xff}

	result:=progressing.TransformationBlackAndWhite(rgba)
	progressing.SaveImage(pathBW, result,"JPG")

	result=progressing.TransformationSepia(rgba,20)
	progressing.SaveImage(pathSepia20, result,"JPG")

	result=progressing.TransformationSepia(rgba,30)
	progressing.SaveImage(pathSepia30, result,"JPG")

	result=progressing.TransformationSepia(rgba,40)
	progressing.SaveImage(pathSepia40, result,"JPG")

}
