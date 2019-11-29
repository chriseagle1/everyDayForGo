package main

import "fmt"

func main() {
	p := NewTimeMatcher(3)

	fmt.Println(p)
}

type TimeMatcher struct {
	base int
}

func NewTimeMatcher(base int) *TimeMatcher {
	return &TimeMatcher{base:base}
}

func doubleAssignment() {
	var a []int = nil

	a, a[0] = []int{1,2}, 9

	fmt.Println(a)
}
