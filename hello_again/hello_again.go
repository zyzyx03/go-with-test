package main

import "fmt"

const englishHelloPrefix = "Hello, "

func HelloAgain(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(HelloAgain("World"))
}
