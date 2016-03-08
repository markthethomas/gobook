package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if 8%4 == 0 {
		fmt.Println("even!")
	}

	if true {
		fmt.Println("true!")
	} else if false {
		fmt.Println("false!")
	} else {
		fmt.Println("who knows?")
	}

	// no ternary operator :(

}
