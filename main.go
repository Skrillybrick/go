package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	foo()
	fmt.Println("after foo (or AF, for those in the know)")
	x := 15
	fmt.Println(x)
	fmt.Println(2 + x)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}
