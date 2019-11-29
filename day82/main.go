package main

import "fmt"

func main() {
	question2()
}

func question1() {
	count := 0
	var arrSet []int8
	var byteSet []byte
	for i := range [256]struct{}{}{
		m, n := byte(i), int8(i)

		if n == -n {
			count++
			arrSet = append(arrSet, n)
		}

		if m == -m {
			count++
			byteSet = append(byteSet, m)
		}
	}

	fmt.Println(count)
	fmt.Println(arrSet)
	fmt.Println(byteSet)
}

const (
	azero = iota
	aone = iota
)

const (
	info = "msg"
	bone = "world"
	bzero = "hello"
	ctwo = iota
	dthree = iota
)
func question2() {
	fmt.Println(azero, aone)
	fmt.Println(bone, ctwo, dthree)
}