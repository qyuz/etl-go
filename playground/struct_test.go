package playground

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	i := Struct{str: "world"}
	fmt.Println(i.str)
}
