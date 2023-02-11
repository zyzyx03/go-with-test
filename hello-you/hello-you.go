package main

import "fmt"

func HelloYou(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(HelloYou("World"))
}
