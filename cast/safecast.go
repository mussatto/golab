package cast

import "fmt"

func Safecast1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	var firstVar interface{}
	var secondVar interface{}
	secondVar = 10
	firstVar = "this is a string"

	firstString, ok := firstVar.(string)

	if (!ok) {
		fmt.Printf("firstString is not a string, do something about it! Value:'%v'\n", firstString)
	}

	secondString, ok := secondVar.(string)

	if (!ok) {
		fmt.Printf("secondString is not a string, do something about it! Value:'%v'\n", secondString)
	}

	secondString = secondVar.(string)

}
