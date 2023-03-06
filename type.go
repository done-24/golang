package main

import "fmt"

func main() {
	var name string = "tom"
	age := 100
	bol := true

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", bol)

	arr := [2]int{1, 2}
	fmt.Printf("%T\n", arr)
}
