package main

import (
	"fmt"
)

func test()(ret int){
	 defer func() {
		ret++
	}()
	return 1
}
func test1(arg ...int) (ret []int){
	if 1 == len(arg){
		return arg
	}
	test1(arg[:2]...)
	return test1(arg...)
}
func main()  {
	for i:=0; i<5; i++{
		defer fmt.Print(i)
	}
	fmt.Print(test())
}