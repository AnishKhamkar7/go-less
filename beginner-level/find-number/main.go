package main

import (
	"fmt"
)

type Result struct {
	index int
	found bool
}

func main() {

	sorted_arr := []int{5, 10, 20, 25, 31, 45}

	target := 50
	fmt.Println("Here is the Index of the target element",find_number_in_sorted_array(sorted_arr,target))

}

func find_number_in_sorted_array(arr []int, tar int) Result {
	l := 0
	r := len(arr) -1 

	for l <= r {
		mid := (l + r) / 2

		if arr[mid] == tar {
			return Result{index: mid, found: true}
		}

		if arr[mid] > tar {
			r = mid - 1

		}else{
			l = mid + 1
		}
	}

	return Result{index: -1, found: false}
}