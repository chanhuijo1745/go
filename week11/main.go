package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])

	var primes [3]int=[3]int{2,3,5}
	fmt.Println(primes, primes[1])

	test:=[5]bool{true,true,true}
	fmt.Println(test[3])
}