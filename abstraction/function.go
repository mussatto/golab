package abstraction

import "fmt"

type AfterWorkCallback func(data []int) string

func Work(data []int, callback AfterWorkCallback) string {

	for index, value := range data {
		data[index] = value + index
	}

	return callback(data)
}

func Print(data []int) string {
	fmt.Println("data is: " + fmt.Sprint(data))
	responseString := ""

	for _, val := range data {
		responseString = responseString + " " + fmt.Sprint(val)
	}

	return responseString
}

func PrintReverse(data []int) string {

	fmt.Println("data is: " + fmt.Sprint(data))
	responseString := ""

	for i := len(data) - 1; i >= 0; i-- {
		responseString = responseString + " " + fmt.Sprint(data[i])
	}

	return responseString
}