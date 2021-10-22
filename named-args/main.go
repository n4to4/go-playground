package main

import "fmt"

// https://blog.uptrace.dev/posts/go-functional-options-named-args.html

type myFuncConfig struct {
	namedArg1 string
	namedArg2 bool
}

type MyFuncOption func(c *myFuncConfig)

func WithNamedArg1(val string) MyFuncOption {
	return func(c *myFuncConfig) {
		c.namedArg1 = val
	}
}

func WithNamedArg2(val bool) MyFuncOption {
	return func(c *myFuncConfig) {
		c.namedArg2 = val
	}
}

func MyFunc(arg1, arg2 string, opts ...MyFuncOption) {
	c := &myFuncConfig{
		namedArg1: "default value",
		namedArg2: true,
	}
	for _, opt := range opts {
		opt(c)
	}

	fmt.Printf("%q, %v\n", c.namedArg1, c.namedArg2)
}

func main() {
	MyFunc("dummy", "dummy",
		WithNamedArg1("this is named arg 1"),
		WithNamedArg2(false),
	)
}
