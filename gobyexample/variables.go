package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var uninitializedString string
	fmt.Println("uninitializedString string is: ", uninitializedString)

	var uninitializedInt int
	fmt.Println("uninitializedInt string is: ", uninitializedInt)

	var inferred = 1
	fmt.Println(inferred)

	var multiple, assignment int = 1, 2
	fmt.Println(multiple, assignment)

	g := "hi"
	// Same as:
	// var f string = "short"
	fmt.Println("g is: ", g)
}
