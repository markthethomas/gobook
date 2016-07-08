package main

import (
	"fmt"
	"reflect"
)

func main() {
	// slices are only typed by containing elements, not Length
	// len() -> <int> length of slice
	// make() -> <slice> creates a slice
	// append() -> <slice> mutated slice
	// copy() -> <slice> mutated slice
	s := make([]string, 3)
	fmt.Println("empty slice: ", s)
	fmt.Println("Length: ", len(s))

	s[0] = "a"
	s[1] = "b"

	fmt.Println("set:", s)
	fmt.Println("get s[1]:", s[1])
	fmt.Println("Still-unassigned s[2] is:", reflect.TypeOf(s[2]))

	// func make([]T, len, cap) -> []T
	// Slices just hold refs to underlying arrays
	newSlice := make([]string, 10)
	fmt.Println(newSlice)

	anotherSlice := make([]string, 10)
	append(anotherSlice, newSlice...)
	fmt.Println(anotherSlice)

}
