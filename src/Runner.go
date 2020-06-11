package main

import (
	"fmt"
	fun "Go_code/src/function"
	"time"
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

	//var rectangle=fun.Rectangle{a:5,b:6}
	var x=fun.Rectangle{A:5,B:6}
	fun.WriteFigureInfo(x)
}
