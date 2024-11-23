package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/ibiscum/Go-Systems-Programming/Chapter12/sharedRPC"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string!")
		os.Exit(100)
	}

	CONNECT := arguments[1]
	c, err := rpc.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	args := sharedRPC.MyInts{A1: 17, A2: 18, S1: true, S2: false}
	var reply int

	err = c.Call("MyInterface.Add", args, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	fmt.Printf("Reply (Add): %d\n", reply)

	err = c.Call("MyInterface.Subtract", args, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	fmt.Printf("Reply (Subtract): %d\n", reply)

}
