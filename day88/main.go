package main

import (
	"fmt"
	"os"
)

func main() {
	question1()
	question2()
}

func question1()  {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}

func foo() error {
	var err *os.PathError = nil

	return err
}

func question2() {
	err := foo()
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	fmt.Println(err == nil)
}
