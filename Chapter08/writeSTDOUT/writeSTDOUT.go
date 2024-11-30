package main

import (
	"io"
	"log"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "You do not give an argument!"
	} else {
		myString = arguments[1]
	}

	buf := []byte(myString)
	_, err := io.Writer.Write(os.Stdout, buf)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(os.Stdout, "\n")
	if err != nil {
		log.Fatal(err)
	}
}
