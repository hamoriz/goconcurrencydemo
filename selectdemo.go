package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := producer1()
	ch2 := producer2()
    doneChannel := make(chan int)
	multipleConsumer(ch1,ch2, doneChannel)

	time.Sleep(4 * time.Second)

    doneChannel <- 1
    time.Sleep(20*time.Second)
}

func producer1() <-chan int {

	myChannel := make(chan int, 5)

	go func() {
		defer close(myChannel)

		for i := 0; i < 20000; i++ {
			myChannel <- i
			fmt.Printf("Number %d sent\n", i)
			time.Sleep(1 * time.Second)
		}

	}()

	return myChannel
}

func producer2() <-chan int {

	myChannel := make(chan int, 5)

	go func() {
		defer close(myChannel)

		for i := 0; i < 200000; i++ {
			myChannel <- i * 10
			fmt.Printf("Number %d sent\n", i*10)
			time.Sleep(1 * time.Second)
		}

	}()

	return myChannel
}


func multipleConsumer(producerChannel1 <-chan int, producerChannel2 <-chan int, done <-chan int) {

	go func() {
		for {
			select {
			case <-done: { fmt.Println("Done"); return}

			case smallNumber := <- producerChannel1 :
				fmt.Printf("I got a small number %d\n", smallNumber)
			case  bigNumber := <- producerChannel2:
				fmt.Printf("I got a big number %d\n", bigNumber)

			}
		}
	}()
}
