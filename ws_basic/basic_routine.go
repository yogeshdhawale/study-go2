package main

import (
	"fmt"
	"sync"
	"time"
)

// usage of go, sync.WaitGroup, defer, panic, recover

var wg sync.WaitGroup

func f1(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, i+1)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}

func routine_eg_1() {

	fmt.Println("1 Go routines eg ...")
	{
		wg.Add(1)
		go f1("Hey")

		wg.Add(1)
		go f1("There")

		wg.Wait()
		wg.Add(1)
		go f1("Last")
		wg.Wait()
	}
	fmt.Println("1 Go routines eg WAIT imapct ...")
	{
		wg.Add(1)
		go f1("Hey")

		wg.Add(1)
		go f1("There")
		wg.Add(1)
		go f1("Last")

		wg.Wait()
	}
}
func f2(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s, i+1)
		time.Sleep(time.Millisecond * 10)
	}
}

func routine_eg_2() {
	fmt.Println("2 Go routines eg ...")
	{
		wg.Add(1)
		go f1("Hey")

		wg.Add(1)
		go f1("There")

		wg.Wait()
		wg.Add(1)
		go f2("Last")
		wg.Wait()
	}
	fmt.Println("2 Go routines eg WAIT imapct ...")
	{
		wg.Add(1)
		go f1("Hey")

		wg.Add(1)
		go f1("There")
		wg.Add(1)
		go f2("Last")

		wg.Wait()
	}
}

func f3_panic(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s, i+1)
		time.Sleep(time.Millisecond * 10)
		if i == 1 {
			panic("***ERROR HERE***")
		}
	}
}

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup function with error:", r)
	}
}

func routine_eg_3() {

	fmt.Println("3 Go routines eg WAIT imapct ...")
	{
		wg.Add(1)
		go f1("Hey")

		wg.Add(1)
		go f2("There")
		wg.Add(1)
		go f3_panic("Last")

		wg.Wait()
	}
	fmt.Println("3 Reached end of function")
}
func routine_eg() {

	// simple go routine with waitgroup
	//routine_eg_1()

	// with defer
	//routine_eg_2()

	// with panic
	// stacking of defers - not a good idea
	routine_eg_3()
}
