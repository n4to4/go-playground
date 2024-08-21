package main

import (
	"fmt"

	"github.com/n4to4/go-playground/range-over-function-types/set"
)

func main() {
	set := set.New[int]()
	fmt.Println(set)
}
