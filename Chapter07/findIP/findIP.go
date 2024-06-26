package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s\n", err)
		os.Exit(-1)
	}
	defer f.Close()

	myIPs := make(map[string]int)
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}

		ip := findIP(line)
		trial := net.ParseIP(ip)
		if trial.To4() == nil {
			continue
		} else {
			_, ok := myIPs[ip]
			if ok {
				myIPs[ip] = myIPs[ip] + 1
			} else {
				myIPs[ip] = 1
			}
		}
	}
	for key := range myIPs {
		fmt.Printf("%s %d\n", key, myIPs[key])
	}
}
