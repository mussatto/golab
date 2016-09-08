package channel

import (
	"fmt"
	"time"
)

//Channel1 creates a string channel and writes into it
func Channel1() {

	stringChannel := make(chan int)

	producer := func(words []int) {
		for _, word := range words {
			stringChannel <- word
		}
	}

	values := []int{}

	for i := 0; i < 100000; i++ {
		values = append(values, i)
	}
	go producer(values)

	consumer := func() {
		for true {
			word := <-stringChannel
			fmt.Println(word)
		}
	}
	//spawning 4 consumers?
	for i := 0; i < 5; i++ {
		go consumer()
	}

	time.Sleep(1000000000)

}
