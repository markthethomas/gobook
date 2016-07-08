package main

import (
	"fmt"
	"reflect"
)

var (
	name   = "Mark"
	course = "Learn React"
	module = 4.2
)

func main() {
	fmt.Println("Name is", name, "and type is ", reflect.TypeOf(name))
	fmt.Println("Module is ", module, "and type is", reflect.TypeOf(module))
	// a := 10.000
	var b bool
	fmt.Println(b)
}
