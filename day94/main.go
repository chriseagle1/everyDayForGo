package main

import "fmt"

func main() {
	question1()
}

func question1()  {
	a := 2 ^ 15 // 0010 ^ 1111 = 1101 -> 13
	b := 4 ^ 15 // 0100 ^ 1111 = 1011 -> 11

	if a > b {
		println("a")
	} else {
		println("b")
	}

	fmt.Println(A("hello"))
	fmt.Println(B(4))
}

func A(string string) string{
	return string + string
}

func B(len int) int {
	return len + len
}

//func c(val, default string) string  {
//	if val == "" {
//		return default
//	}
//	return val
//}