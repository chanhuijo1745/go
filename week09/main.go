package main

import "fmt"

func main() {
	a := 10
	b := 20
	var pa *int = &a

	fmt.Printf("%T %T\n", &a, pa)
	fmt.Printf("%x %x %x\n", &a, pa, &pa)
	fmt.Println(&a, pa, &pa)
	fmt.Println(*pa)
	pa = &b
	fmt.Println(*pa)
}