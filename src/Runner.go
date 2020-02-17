package main

import (
	"fmt"
	fun "Go_code/src/function"
)

func main() {
	fmt.Println(fun.Add(42, 13))
	var n int =5
	fmt.Printf("Factor %d = %d",n, fun.Factor(n))
}
