package main

import "fmt"

func avg_quiz(arr []float64)(avg float64){
	var sum float64
	switch len(arr) {
	case 0:
		return 0
	default:
		for _, value:=range arr{
			sum += value
		}
		avg = sum/float64(len(arr))
	}
	return
}

func swap_quiz(a int , b int)(ret_a int, ret_b int){
	if a>b {
		ret_a, ret_b = b, a
	}else{
		ret_a, ret_b = a, b
	}
	return
}

func main()  {
	var (
			avg float64
		)
	arry := []float64{10.0,11.1,2.3,2.7,10.7,19.00,20.1}
	avg   = avg_quiz(arry)
	a, b := swap_quiz(2,1)
	fmt.Println(a,b)
	fmt.Println(avg)
}