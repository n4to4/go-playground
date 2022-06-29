package main

import (
	"fmt"
	"os"
)

func main() {
	hello := "world"
	if len(os.Args) == 2 {
		hello = os.Args[1]
	}
	fmt.Printf("Hello, %s\n", hello)
}
