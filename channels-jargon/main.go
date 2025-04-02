package main

import (
	"fmt"
)

func main() {

	arr := []int{5, 5}
	arr2 := []int{5, 6}
	arr3 := []int{5, 7}

	c := make(chan int)
	go sumOfArray(arr, c)
	go sumOfArray(arr2, c)
	go sumOfArray(arr3, c)
	a,b,d:= <-c, <- c, <- c	
	fmt.Println("C::",a,b,d )
}

func sumOfArray(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	
	c <- sum
	
}
