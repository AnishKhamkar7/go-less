package main

import "fmt"

func main() {
	var username string = "Foo"
	// var age = 20
	// isDeveloper := true
	fmt.Println("hello ",username)

	sum := add(10,20)

	fmt.Println("Here is the addtion",sum)

	var arr = []int{1,2,4,5}

	for i:= 0; i<= sum; i++{
		arr = append(arr,i)
		fmt.Println(arr)
	}

	user := User{Name: "FOO",Age: 2}
	fmt.Println("Struct Implementation:",user)
}

func add (a int, b int  )int{
	return a + b
}

type User struct {
	Name string
	Age int
}