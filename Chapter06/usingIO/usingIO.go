package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening %s: %s", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	buf := make([]byte, 8)
	if _, err := io.ReadFull(f, buf); err != nil {
		if err == io.EOF {
			// err = io.ErrUnexpectedEOF
			log.Println("error EOF")
		}
	}

	_, err = io.Writer.Write(os.Stdout, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
