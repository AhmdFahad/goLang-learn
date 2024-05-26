package main

import "fmt"

func main() {

	var name = 231321

	fmt.Println("Hello World!")
	fmt.Println(r(2))
	fmt.Printf("type of name is %T", name)
	forLoop(10)
}

func r(h int) int {
	x := 20
	return x + h
}

func forLoop(count int) {
	for i := 0; i < count; i++ {
		fmt.Println(i)

	}
}

type Person struct {
	name   string
	age    int
	job    string
	salary int
}
