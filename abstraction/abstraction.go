package abstraction

import "fmt"

type Animal interface {
	Talk() string
}

type Dog struct {
	name string
}

func (dog Dog) Talk() string{
	return "woof woof"
}

type Cat struct{
	name string
}

func (cat Cat) Talk() string{
	return "meow"
}

func createAnimals() []Animal{
	return []Animal{ Dog{"snoopy"}, Cat{"mousty"}, Dog{"rex"}, Cat{"garfield"}, Cat{"fat garfield"}, Dog{"toto"} }
}

func MakeAnimalsTalk(){
	for _, animal := range createAnimals(){
		fmt.Println(animal.Talk())
	}
}