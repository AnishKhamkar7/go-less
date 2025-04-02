package main

import (
	"fmt"
	"time"
)

func main() {

	go oddNum(10)
 	go evenNum(10)

	time.Sleep(1* time.Second)

}

func evenNum(limit int) {
	for i := range limit {
		if i%2 == 0 {
			fmt.Println("EVEN",i)
			// time.Sleep(time.Millisecond * 500)
		}	
	}
}

func oddNum(limit int) {
	for i := range limit {
		if i%2 != 0 {
			fmt.Println("ODD",i)
		}
	}
}