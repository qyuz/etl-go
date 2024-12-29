package main

import (
	"errors"
	"fmt"
)

func main() {
	// result, err := myFunction()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	result, _ := myFunction()
	fmt.Println(result)
}

func myFunction() (result int, err error) {
	// ... some logic ...
	if true {
		err = errors.New("custom error message")
		return
	}
	// ... more logic ...
	result = 42
	return
}
