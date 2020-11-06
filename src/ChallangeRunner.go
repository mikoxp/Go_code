package main

import (
	"fmt"
	"strconv"
)

func main() {
	var index int =1
	var x=5
	var y=10
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print("___")
		}
		fmt.Println()
		fmt.Print("|")
		for j := 0; j < 10; j++ {
			strconv.Itoa(index)
			fmt.Printf("%2s|",strconv.Itoa(index))
			index++
		}
		fmt.Println()
	}
	for j := 0; j < 10; j++ {
		fmt.Print("___")
	}
}
