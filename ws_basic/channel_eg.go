package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func foo_send(c chan int, val int) {
	defer wg2.Done()
	c <- val * val
}
func foo_recv(c chan int) int {
	// return <-c
	val := <-c
	return val
}

func producer(c chan int) {

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go foo_send(c, i+1)
	}

	wg2.Wait()
	close(c)
}

func consumer(c chan int) {
	for item := range c {
		fmt.Println("Value is:", item)
	}
}
func channel_eg() {
	// without size the data comes in different order

	if false {
		fmt.Println("Channel buffer  = unspecified, result in deadlock")
		cVar := make(chan int)
		producer(cVar)
		consumer(cVar)
	}
	{
		fmt.Println("Channel buffer  = 12")
		cVar := make(chan int, 12)
		producer(cVar)
		consumer(cVar)
	}
}
