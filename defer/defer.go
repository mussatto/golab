package golab

import "fmt"

//Defer1
func defer1() {
	defer fmt.Println("this is going to be printed after return")

	fmt.Println("this is A")
	fmt.Println("this is B")
}

func defer2() {
	//since its LIFO, should print 3210
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}

}
