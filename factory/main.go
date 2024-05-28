package main

import "fmt"

type Animaler interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Cat"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Dog"
}

func AnimalFactory(animalType string) Animaler {
	switch animalType {
	case "cat":
		return Cat{}
	case "dog":
		return Dog{}
	default:
		return Dog{}
	}
}

func main() {
	animal := AnimalFactory("dog")
	fmt.Println(animal.Speak())

	animal = AnimalFactory("cat")
	fmt.Println(animal.Speak())
}
