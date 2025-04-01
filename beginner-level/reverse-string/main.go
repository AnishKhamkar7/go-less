package main

import "fmt"

func main() {

	var newStr string = "toReverseString"

	fmt.Println("Reversed STRING:::::::::",reverseString(newStr))

}

func reverseString(abc string) string {
	var mutableStr = ""
	
	for i:= len(abc)-1 ; i >= 0; i-- {
		mutableStr = mutableStr + string(abc[i])
	}

	return mutableStr
}