package main

import "fmt"

func main() {
	question1()
	question2()
}

func question1() {
	n := 43210

	fmt.Println(n/(60*60), " hours and ", n%(60*60), " seconds")
}

const (
	Century = 0b100
	Decade  = 0b010
	Year    = 0b001
)
func question2() {
	fmt.Println(Century, Decade, Year)
	fmt.Println(Century + 2 * Decade + 2 * Year)
}