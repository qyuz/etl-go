package playground

import "fmt"

type Struct struct {
	str string
}

func main() {
	i := Struct{str: "world"}
	fmt.Println(i.str)
}
