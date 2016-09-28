package abstraction

import (
	"testing"
	"fmt"
)

func TestWork(t *testing.T) {

	input := []int{}
	input2 := []int{}
	for i:=0; i<10; i++{
		input = append(input, i)
		input2 = append(input2, i)
	}

	fmt.Println("Print output:")
	fmt.Println(Work(input, Print))

	fmt.Println("PrintReverse output:")
	fmt.Println(Work(input2, PrintReverse))
}
