package main

import "fmt"

func main() {
	var g Greeter
	g = Helloer{}
	g.Greet()
}

type Greeter interface{ Greet() }

type (
	Helloer  struct{}
	Goodbyer struct{}
)

var (
	_ Greeter = Helloer{}  // Helloer implements Greeter
	_ Greeter = Goodbyer{} // Goodbyer implements Greeter
)

func (Helloer) Greet()  { hello() }
func (Goodbyer) Greet() { goodbye() }

func hello()   { fmt.Println("hello") }
func goodbye() { fmt.Println("goodbye") }
