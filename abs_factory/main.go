package main

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Cat"
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Dog"
}

type AnimalFactory interface {
	CreateAnimal() Animal
}

type CatFactory struct{}

func (c CatFactory) CreateAnimal() Animal {
	return &Cat{}
}

type DogFactory struct{}

func (d DogFactory) CreateAnimal() Animal {
	return &Dog{}
}

type Shower interface {
	ShowAnimal() string
}

type PetShopA struct {
	AnimalFactory
}

func (p PetShopA) ShowAnimal() string {
	return p.AnimalFactory.CreateAnimal().Speak()
}

type PetShopB struct {
	AnimalFactory
}

func (p PetShopB) ShowAnimal() string {
	return p.AnimalFactory.CreateAnimal().Speak()
}

func PetShopFactory(petShopType string, factory AnimalFactory) Shower {
	switch petShopType {
	case "A":
		return &PetShopA{AnimalFactory: factory}
	case "B":
		return &PetShopB{AnimalFactory: factory}
	default:
		return &PetShopA{AnimalFactory: factory}
	}
}

func main() {
	dog := DogFactory{}
	petShopA := PetShopFactory("A", dog)
	fmt.Println(petShopA.ShowAnimal())

	cat := CatFactory{}
	petShopB := PetShopB{AnimalFactory: cat}
	fmt.Println(petShopB.ShowAnimal())
}
