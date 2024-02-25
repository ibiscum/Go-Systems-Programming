package main

import (
	"fmt"

	"github.com/ibiscum/Go-Systems-Programming/Chapter04/internal/anotherPackage"
)

func main() {
	anotherPackage.Version()
	// fmt.Println(anotherPackage.version)
	fmt.Println(anotherPackage.Pi)
}
