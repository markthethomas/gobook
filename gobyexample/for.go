package main

import "fmt"

func main() {
	// single condition loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// class inital/condition/after loop
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	// while loop
	for {
		fmt.Println("looping")
		break
	}
}
