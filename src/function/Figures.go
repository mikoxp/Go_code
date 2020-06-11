package function

import (
	"fmt"
	"reflect"
)

type Figure interface {
	field() float64
}

type Rectangle struct {
	A, B float64
}

type Square struct {
	A float64
}

type Triangle struct {
	A, H float64
}

//------------------------------------
func (el Rectangle) field() float64 {
	return el.A * el.B
}

func (el Square) field() float64 {
	return el.A * el.A
}
func (el Triangle) field() float64 {
	return 0.5 * el.A * el.H
}

func WriteFigureInfo(figure Figure) {
	fmt.Println(figure)
	fmt.Println(reflect.TypeOf(figure))
	fmt.Printf("Figure field= %.2f", figure.field())
}
