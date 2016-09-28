package strings

import (
	"fmt"
	"testing"
	"time"
)

func TestConcatenate(t *testing.T) {
	start := time.Now()
	ConcatenateWrong()
	elapsed := time.Since(start)
	fmt.Printf("ConcatenateWrong - Elapsed time is %s\n", elapsed)
	//fmt.Println(wrong)

	start = time.Now()
	ConcatenateRight()
	elapsed = time.Since(start)
	fmt.Printf("ConcatenateRight - Elapsed time is %s\n", elapsed)
	//fmt.Println(right)
}
