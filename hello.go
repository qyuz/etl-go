package main

import "fmt"

func main() {
	a := 1
	fmt.Println("Hello, World!")
	fmt.Printf("The value of a is: %d\n", a)
	ini(a)
}

func ini(a int) {
	fmt.Printf("Init function %d", a)
}
