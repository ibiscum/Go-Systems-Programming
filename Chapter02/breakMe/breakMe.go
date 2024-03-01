package main

import "fmt"

func main() {
	myArray := [4]int{1, 2, 4, -4}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	fmt.Println("myArray[-1]:", myArray[-1])            //nolint
	fmt.Println("myArray[10]", myArray[10])             //nolint
	fmt.Println("threeD[-1][20][0]", threeD[-1][20][0]) //nolint
}
