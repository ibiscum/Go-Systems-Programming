//go:build ignore

package main

import (
	"fmt"
)

func x() int {
	return -1
	fmt.Println("Exiting x()") //nolint
}

func y() int {
	return -1
	fmt.Println("Exiting y()") //nolint
}

func main() {
	fmt.Println(x())
	fmt.Println("Exiting program...")
}
