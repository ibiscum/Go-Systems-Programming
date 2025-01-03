package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

func createNumber(max int, randomNumberChannel chan<- int, finishedChannel chan bool) {
	for {
		select {
		case randomNumberChannel <- rand.Intn(max):
		case x := <-finishedChannel:
			if x {
				close(finishedChannel)
				close(randomNumberChannel)
				return
			}
		}
	}
}

func main() {
	// rand.Seed(time.Now().Unix())
	randomNumberChannel := make(chan int)
	finishedChannel := make(chan bool)

	if len(os.Args) != 3 {
		fmt.Printf("usage: %s count max\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	n1, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Printf("Invalid count value: %s\n", os.Args[1])
		os.Exit(1)
	}
	count := int(n1)
	n2, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		fmt.Printf("Invalid max value: %s\n", os.Args[2])
		os.Exit(1)
	}
	max := int(n2)

	fmt.Printf("Going to create %d random numbers.\n", count)
	go createNumber(max, randomNumberChannel, finishedChannel)
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", <-randomNumberChannel)
	}

	finishedChannel <- false
	fmt.Println()
	_, ok := <-randomNumberChannel
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	finishedChannel <- true
	_, ok = <-randomNumberChannel
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
