package main

import "fmt"

func New(str string) *Demo {
	return &Demo{Str: str}
}

type Demo struct {
	Str string
}

func main() {
	fmt.Printf("%v", *New("Hello, World!"))
}
