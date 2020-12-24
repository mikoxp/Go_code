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
	pathPrewitt:="image/tmp_prewitt.jpg"
	pathPrewittOut:="image/tmp_prewitt_out.jpg"
	pathLenaOut:="image/lena_prewitt_out.jpg"
	rgba := progressing.OpenRGBAImage(path)

	result:=progressing.TransformationBlackAndWhite(rgba)
	progressing.SaveImage(pathBW, result,"JPG")

	result=progressing.TransformationSepia(rgba,20)
	progressing.SaveImage(pathSepia20, result,"JPG")

	result=progressing.TransformationSepia(rgba,30)
	progressing.SaveImage(pathSepia30, result,"JPG")

	result=progressing.TransformationSepia(rgba,40)
	progressing.SaveImage(pathSepia40, result,"JPG")

	mask := []int{1,1, 1,0,0,0,-1,-1,-1}

	result=progressing.Prewitt(rgba,mask)
	progressing.SaveImage(pathLenaOut, result,"JPG")

	rgbaBlack := progressing.OpenRGBAImage(pathPrewitt)


	result=progressing.Prewitt(rgbaBlack,mask)
	progressing.SaveImage(pathPrewittOut, result,"JPG")

}
