package main

import (
	"fmt"
	"time"
)

func ProcessData(data Order) string {

	fmt.Printf("Preparing order for table %d...",data.TNumber)

	time.Sleep(data.PrepTime)

	fmt.Printf("Order Ready got table %d",data.TNumber)
	return "ended"
}
