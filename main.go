package main

import (
	"fmt"
	"os"
)

func server() {
	var pid any
	pid = os.Getppid()

	fmt.Println(pid)
}

func main() {
	fmt.Println("Hello world")
	server()
}
