package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hellow world\n")

	const(
		  a = iota
	      b
	)
	fmt.Printf("result a+b = %d", a+b)

}


