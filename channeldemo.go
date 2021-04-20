package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	ch := producer()

	consumer(ch)
	wg.Wait()
}


func producer() <-chan int {

	myChannel := make(chan int,5)

	go func() {
		defer close(myChannel)
		defer wg.Done()
		for i := 0; i < 20; i++ {
			myChannel <- i
			fmt.Printf("Number %d sent\n",i)
			time.Sleep(1 * time.Second)
		}

	}()

	return myChannel
}

func consumer(producerChannel <-chan int) {

	go func() {
		defer wg.Done()
		for i:= range producerChannel {
			time.Sleep(1 * time.Second)
			fmt.Printf("I got %d\n",i)
		}
	}()
}
