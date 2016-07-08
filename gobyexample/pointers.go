package main

import "fmt"

func main() {
	// & gets the memory address of something
	// * dereferences it

	// Create a variable
	zip := "zip"

	// Reference the variable's memory address
	ptr := &zip

	fmt.Println(ptr) // -> 0xc82000a3b0

	// Dereferencing
	fmt.Println(*ptr) // -> "zip"

	// Change the value
	changeZip(&zip)

	// Check the value again
	fmt.Println(zip)
}

// Exploring pass by reference vs pass by value
// tell Go that we're passing in a pointer to a string
func changeZip(course *string) string {
	// Dereference the value and assign to it
	*course = "baz!"
	fmt.Println("trying to change zip to ", *course)
	// return the de-reffed value
	return *course
}
