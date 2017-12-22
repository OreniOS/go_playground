package main

import "fmt"

func main() {

	scope()
}

func scope() {
	x := 0
	fmt.Println("This is an local function scope ", x)
}
