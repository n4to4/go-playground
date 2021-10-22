package main

import (
	"fmt"
	"io/fs"
	"os"
)

// https://blog.uptrace.dev/posts/go-functional-options-named-args.html

type myFuncConfig struct {
	namedArg1 string
	namedArg2 bool
	dir       fs.FS
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

func WithDir(dir string) MyFuncOption {
	return func(c *myFuncConfig) {
		c.dir = os.DirFS(dir)
	}
}

func WithDirFS(dir fs.FS) MyFuncOption {
	return func(c *myFuncConfig) {
		c.dir = dir
	}
}

func DirFromEnv(envName string) MyFuncOption {
	return func(c *myFuncConfig) {
		if dir, ok := os.LookupEnv(envName); ok {
			c.dir = os.DirFS(dir)
		}
	}
}

func MyFunc(arg1, arg2 string, opts ...MyFuncOption) {
	c := &myFuncConfig{
		namedArg1: "default value",
		namedArg2: true,
		dir:       os.DirFS("/"),
	}
	for _, opt := range opts {
		opt(c)
	}

	fmt.Printf("%+v\n", c)
}

func main() {
	MyFunc("dummy", "dummy",
		WithNamedArg1("this is named arg 1"),
		WithNamedArg2(false),
		WithDir("/home/me"),
	)
}
