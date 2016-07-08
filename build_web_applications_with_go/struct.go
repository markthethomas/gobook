package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	var Mark Person
	Mark.name = "Mark"
	Mark.age = 24
	fmt.Printf("Hi! I'm %v years old and my name is %v", Mark.age, Mark.name)
}
