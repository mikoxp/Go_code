package main

import (
	"fmt"
	fun "Go_code/src/function"
	"time"
	factory "Go_code/src/write_factory"
)

func main() {
	fmt.Println(fun.Add(42, 13))
	var n int = 5
	fmt.Printf("Factor %d = %d\n", n, fun.Factor(n))
	var tab = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("---------SORT---------------")
	fmt.Println(tab)
	var tabSorted = fun.BubbleSort(tab[:])
	fmt.Println(tabSorted)
	fmt.Println("------------------------")
	var fileName="dates"
	var line=time.Now().String()
	fun.AddLineToFile(fileName,line)
	var text=fun.ReadTextFile(fileName)
	fmt.Println(text)
	fmt.Println("------------------------")

	var rect=fun.Rectangle{A:5,B:6}
	var square=fun.Square{A:5}
	var triangle=fun.Triangle{A:6,H:7}
	fun.WriteFigureInfo(rect)
	fun.WriteFigureInfo(square)
	fun.WriteFigureInfo(triangle)

	fmt.Println("------------------------")
	var w1=factory.WriteObject{Place:"FILE",Data:"DATA",FileName:fileName}
	var w2=factory.WriteObject{Place:"OUT",Data:"DATA"}
	factory.WriteTextToPlace(w1)
	factory.WriteTextToPlace(w2)

}
