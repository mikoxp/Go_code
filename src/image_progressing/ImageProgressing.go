package image_progressing

import (
	"image"
	"os"
	"image/jpeg"
	"image/draw"
	"image/png"
	"image/color"
)
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

func TransformationBlackAndWhite(img *image.RGBA) *image.RGBA{
	width := img.Rect.Dx()
	height := img.Rect.Dy()
	result:=createEmptyImage(width,height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := img.RGBAAt(x, y)
			val:=(int32(p.R)+int32(p.G)+int32(p.B))/3
			c:=uint8(val)
			result.Set(x, y, color.RGBA{R:c,G:c,B:c,A:p.A})
		}
	}
	return result
}

func TransformationSepia(img *image.RGBA,w int32) *image.RGBA{
	width := img.Rect.Dx()
	height := img.Rect.Dy()
	result:=createEmptyImage(width,height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := img.RGBAAt(x, y)
			r:=int32(p.R)*2+w
			if(r>255){
				r=255
			}
			g:=int32(p.G)+w
			if(g>255){
				g=255
			}
			result.Set(x, y, color.RGBA{R:uint8(r),G:uint8(g),B:p.B,A:p.A})
		}
	}
	return result
}

func createEmptyImage(width int,height int) *image.RGBA{
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	result := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	return result
}
