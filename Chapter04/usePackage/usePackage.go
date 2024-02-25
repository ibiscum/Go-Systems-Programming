package main

import (
	"fmt"

	"github.com/ibiscum/Go-Systems-Programming/Chapter04/internal/aSimplePackage"
)

func main() {
	temp := aSimplePackage.Add(5, 10)
	fmt.Println(temp)
	fmt.Println(aSimplePackage.Pi)
}
