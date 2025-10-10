package main

import (
	"sync"
	"time"
)

type Order struct {
	TNumber  int
	PrepTime time.Duration
}

func main() {

	orders := []Order{

		{
			TNumber: 1,
			PrepTime: 2,
		},
		{
			TNumber: 2,
			PrepTime: 8,
		},
		{
			TNumber: 3,
			PrepTime: 4,
		},


	}
	wg := sync.WaitGroup{}
	for _,order := range orders {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			go ProcessData(order)
		}()
	}

	wg.Wait()
}