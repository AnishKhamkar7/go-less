package main

import "fmt"

func main() {

	num := 11

	var checkNumber = isEven(num)

	fmt.Println("iSeven:",checkNumber)

}

func isEven(a int) bool {
	if a % 2 == 0 {
		return true
	}

	return false
}