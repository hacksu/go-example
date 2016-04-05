package main

import "fmt"

func main() {
	defer fmt.Println("two")
	fmt.Println("one")
}
