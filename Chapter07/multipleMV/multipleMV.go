package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var RE string
var renameString string

func walk(path string, f os.FileInfo, errA error) error {
	regex, err := regexp.Compile(RE)
	if err != nil {
		fmt.Printf("Error in RE: %s\n", RE)
		return err
	}

	if path == "." {
		return nil
	}
	nameOfFile := filepath.Base(path)
	if regex.MatchString(nameOfFile) {
		newName := filepath.Dir(path) + "/" + renameString + "_" + nameOfFile
		err := os.Rename(path, newName)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 3 {
		fmt.Printf("Usage: %s REGEXP RENAME Path.", filepath.Base(os.Args[0]))
		os.Exit(-1)
	}

	RE = flag.Arg(0)
	renameString = flag.Arg(1)
	Path := flag.Arg(2)
	Path, _ = filepath.EvalSymlinks(Path)
	err := filepath.Walk(Path, walk)
	if err != nil {
		log.Fatal(err)
	}
}
