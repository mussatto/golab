package channel


//Channel1 creates a string channel and writes into it
func Channel1() {

	stringChannel := make(chan int, 5)

	producer := func(words []int) {
		for _, word := range words {
			stringChannel <- word
		}
	}

	values := []int{}

	for i := 0; i < 1000; i++ {
		values = append(values, i)
	}
	go producer(values)

	//for word := range stringChannel {
	//	fmt.Println(word)
	//	time.Sleep(5000)
	//}

	//time.Sleep(100000)

}
