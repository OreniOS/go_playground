package main

import "fmt"

func main() {

	x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println("This is an inner scope level ", increment())

	wr := wrapper()
	fmt.Println(wr())
	fmt.Println(wr())
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
