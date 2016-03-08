package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("a constant string: ", s)

	const n = 500000

	fmt.Println(math.Sin(n))

}
