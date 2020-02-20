package main

import (
	"fmt"
	fun "Go_code/src/function"
)

func main() {
	fmt.Println(fun.Add(42, 13))
	var n int =5
	fmt.Printf("Factor %d = %d",n, fun.Factor(n))
	var tab = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("")
	fmt.Println(tab)
	var tabSorted =fun.BubbleSort(tab[:])
	fmt.Println(tabSorted)

}
