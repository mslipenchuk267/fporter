package main

import (
	"fmt"
	"os"
)

func main() {
	// Parse Arguments
	args1 := os.Args
	args2 := os.Args[1:]

	fmt.Println("Arguments with Exe Call:")
	fmt.Println(args1)
	fmt.Println("Arguments without Exe Call:")
	fmt.Println(args2)
}
