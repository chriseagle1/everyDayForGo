package main

import "fmt"

type User struct{
	Name string
}

func (u *User)SetName(name string)  {
	u.Name = name
	fmt.Println(u.Name)
}

type Employee User

func main() {
	employee := new(Employee)

	employee.SetName("Jack") //employee.SetName undefined  type 定义新类型不会继承原有类型的方法集
}
