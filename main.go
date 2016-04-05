package main

import "fmt"



func main() {
	defer fmt.Println("two")
	fmt.Println(is_prime(1000))
}

func is_prime(n int) int {
	for i := 2; i < n/2; i++ {
		if n % i == 0 {
			return i;
		}
	}
	return -1;
}
