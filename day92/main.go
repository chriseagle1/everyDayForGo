package main

import "fmt"

var x int

func init() {
	x++
}

func main() {

	fmt.Println(x)
	min(1225, 256)
}

func min(a int, b uint) {
	var min = 0
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("The min of %d and %d id %d\n", a, b, min)
}
