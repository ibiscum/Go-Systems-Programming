package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two command line arguments!")
		os.Exit(1)
	}

	sourceFile := os.Args[1]
	destinationFile := os.Args[2]

	input, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating the new file", destinationFile)
		fmt.Println(err)
		os.Exit(1)
	}
}
