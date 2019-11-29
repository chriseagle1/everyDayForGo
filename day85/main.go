package main

import (
	"fmt"
	"strings"
)

func main() {
	question1()
	question2()
}

func question1() {
	str := "BACBBA"
	sli := strings.TrimRight(str, "BA")
	fmt.Println(sli)//BAC

	fmt.Println(strings.TrimLeft("llolohe Tlo", "loh"))
}

func question2()  {
	var src []int
	dist := make([]int, 4)
	src = []int{1,2,3}
	copy(dist, src)

	fmt.Println(dist)
}