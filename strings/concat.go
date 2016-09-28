package strings

import "bytes"

const maxIterations = 100000

func ConcatenateWrong() string {
	myVar := ""
	for i := 0; i < maxIterations; i++ {
		myVar = myVar + GenerateRandom()
	}
	return myVar
}

func ConcatenateRight() string {
	var buffer bytes.Buffer
	for i := 0; i < maxIterations; i++ {
		buffer.WriteString(GenerateRandom())
	}
	return buffer.String()
}