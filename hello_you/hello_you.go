package main

import "fmt"

const englishHelloPrefix = "Hello, "

func HelloYou(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(HelloYou("world"))
}
