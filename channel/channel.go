package channel

import (
	"fmt"
	"time"
)

//Channel1 creates a string channel and writes into it
func Channel1() {

	stringChannel := make(chan int, 5)

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

	for word := range stringChannel {
		fmt.Println(word)
		time.Sleep(5000000)
	}

	time.Sleep(1000000000)

}
