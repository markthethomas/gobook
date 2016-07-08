package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("an empty array: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("lenght: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("intialized: ", b)

	fmt.Println("We can even create 2D arrays!")

	const width = 2
	const height = 2
	const depth = 2
	var threeD [width][depth][depth]int
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			for z := 0; z < depth; z++ {
				threeD[i][j][z] = i + j + z
			}
		}
	}

	fmt.Println(threeD)
}
