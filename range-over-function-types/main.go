package main

import (
	"fmt"

	"github.com/n4to4/go-playground/range-over-function-types/set"
)

func main() {
	set1 := set.New[int]()
	set1.Add(42)
	set1.Add(123)

	set2 := set.New[int]()
	set2.Add(42)
	set2.Add(99)

	union := set.Union(set1, set2)
	fmt.Println(union)
}
