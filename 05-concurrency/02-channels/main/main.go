package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	//<-ch
	//fmt.Println("Here")

	go func() {
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("got %d\n", val)

	fmt.Println("-----")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("-----")

	// example with closing of channel (to avoid deadlock) and getting values in range
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
}
