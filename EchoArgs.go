package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value  := range os.Args[:]{
		fmt.Printf("index = %v, value = %v\n", index, value)
	}
}