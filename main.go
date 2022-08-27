package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return "Hello, " + name
}
