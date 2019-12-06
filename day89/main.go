package main

import "fmt"

func main() {
	question1()
	question2()
}

func question1() {
	v := []int{1,2,3}
	for i, n := 0, len(v); i < n; i++ {
		v = append(v, i)
	}

	fmt.Println(v) //1,2,3, 0, 1, 2
}

type P *int
type Q *int

func question2()  {
	var p P = new(int)
	*p += 8
	var x *int = p
	var q Q = x
	*q++

	fmt.Println(*p, *q) //9, 9
}
