package main

import "fmt"
import "time"
import "sort"

func main() {
	start := time.Now()
	fmt.Println(slow_primes(0, 100035))
	duration := time.Since(start)
	fmt.Println(duration)
}

func primes(start, end int) []int {
	primes := make(chan []int)

	chunks := (end - start) / 1000 + 1
	for i := 1; i < chunks; i++ {
		go fast_primes((i-1)*1000 + start, i*1000 + start, primes)
	}
	go fast_primes((chunks-1)*1000 + start, end, primes)

	result := make([]int, 0)
	for i := 0; i < chunks; i++ {
		result = append(result, (<-primes)...)
	}
	sort.Ints(result)
	return result
}

func fast_primes(start, end int, primes chan []int) {
	result := make([]int, 0)
	for i := start; i < end; i++ {
		if is_prime(i) < 0 {
			result = append(result, i)
		}
	}

	primes <- result;
}

func slow_primes(start, end int) []int{
	result := make([]int, 0)
	for i := start; i < end; i++ {
		if is_prime(i) < 0 {
			result = append(result, i)
		}
	}

	return result
}
func is_prime(n int) int {
	for i := 2; i <= n/2; i++ {
		if n % i == 0 {
			return i;
		}
	}
	return -1;
}
