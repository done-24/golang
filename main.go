package main

import (
	"fmt"
	"jj/v2/user"
)

func main() {
	s := user.Hello()
	fmt.Printf("s: %v\n", s)

	// var name string
	// var age int
	// var bol bool

	// var (
	// 	name string
	// 	age  int
	// 	bol  bool
	// )

	// var (
	// 	name string = "tom"
	// 	age  int    = 20
	// 	bol  bool   = true
	// )

	// var name, age, bol = "tom", 20, true

	name := "tom"
	age := 20
	bol := true

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("bol: %v\n", bol)
}
