package main

import "fmt"

func hello() string {
	return "Hello"
}

type Struct struct {}

func (s *Struct) AnotherHello() string {
	return "Another Hello"
}

func main() {
	fmt.Println(hello())

	another := &Struct{}
	fmt.Println(another.AnotherHello())

	superStruct := &Struct{}
	fmt.Println(superStruct.AnotherHello())
}
