package main

import (
	"time"
)

func main() {

	
	foo := 56

	time.AfterFunc(1*time.Second, func() {
		for i:= 0; i < 10000; i++ {
			go writefun(&foo,i)
			} 
	})

	println("Before ::",foo)

	for i:= 0; i < 10000; i++ {
		go writefun(&foo,i)
	}
	
	time.Sleep(7 * time.Second)
	println("After ::",foo)
}

func writefun(foo *int, boo int){
	*foo = *foo + boo
	
}
