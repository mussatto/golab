package slices

import "fmt"

func DoAppendSlices(){
	myStringSlice := []string{"first", "second", "third"}
	//ERROR!
	//myStringSlice = append(myStringSlice, []string{"fourth", "fift"})
	myStringSlice = append(myStringSlice, []string{"fourth", "fift"}...)
	myStringSlice = append(myStringSlice, "sixth", "seventh")
	fmt.Println(myStringSlice)
}