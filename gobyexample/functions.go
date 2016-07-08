package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello, world!"
	fmt.Println(titleCase(text))
}

func titleCase(text string) string {
	return strings.Title(text)
}
