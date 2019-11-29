package main

import "fmt"

func main() {
	a := [3]int{0,1,2}
	fmt.Println(a, len(a), cap(a))
	s := a[1:2]

	fmt.Println(s, len(s), cap(s))

	s[0] = 11
	fmt.Println(a)
	fmt.Println(s)
	s = append(s, 12)
	s = append(s, 13)
	s[0] = 21

	fmt.Println(a)
	fmt.Println(s)
}
