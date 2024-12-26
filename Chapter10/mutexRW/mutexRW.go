package main

import (
	"fmt"
	"sync"
	"time"
)

var Password = secret{counter: 1, password: "myPassword"}

type secret struct {
	sync.RWMutex
	counter  int
	password string
}

func Change(c *secret, pass string) {
	c.Lock()
	fmt.Println("LChange")
	time.Sleep(20 * time.Second)
	c.counter = c.counter + 1
	c.password = pass
	c.Unlock()
}

func Show(c *secret) string {
	fmt.Println("LShow: Accessing password")
	time.Sleep(time.Second)
	c.RLock()
	defer c.RUnlock()
	return "****" // Do not return the actual password
}

func Counts(c *secret) int {
	c.RLock()
	defer c.RUnlock()
	return c.counter
}

func main() {

	fmt.Println("Accessing password")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Accessing password in goroutine")
		}()
	}

	go func() {
		Change(&Password, "123456")
	}()

	fmt.Println("Accessing password")
	time.Sleep(time.Second)
	fmt.Println("Counter:", Counts(&Password))
}
