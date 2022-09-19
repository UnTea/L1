package main

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Human is a description of a human
type Human struct {
	Name   string
	Age    int16
	Height float32
	Weight float32
}

// NewHuman is constructor of Human structure
func NewHuman(name string, age int16, height float32, weight float32) Human {
	return Human{
		Name:   name,
		Age:    age,
		Height: height,
		Weight: weight,
	}
}

// String is a string representation of human structure
func (human *Human) String() string {
	return fmt.Sprintf("Name=%v\nAge=%v\nHeight=%v\nWeight=%v\n",
		human.Name, human.Age, human.Height, human.Weight)
}

// SetName is a function that sets a new value for the Name field of the Human object
func (human *Human) SetName(name string) error {
	if len(name) <= 0 {
		return errors.New("name must contain at least one letter")
	}

	human.Name = cases.Title(language.Und).String(cases.Lower(language.Und).String(name))

	return nil
}

// SetAge is a function that sets a new value for the Age field of the Human object
func (human *Human) SetAge(age int16) error {
	if age < 0 {
		return errors.New("age must be greater than or equal to zero")
	}

	human.Age = age

	return nil
}

// SetHeight is a function that sets a new value for the Height field of the Human object
func (human *Human) SetHeight(height float32) error {
	if height < float32(0) {
		return errors.New("height must be greater than zero")
	}

	human.Height = height

	return nil
}

// SetWeight is a function that sets a new value for the Weight field of the Human object
func (human *Human) SetWeight(weight float32) error {
	if weight < float32(0) {
		return errors.New("age must be greater than zero")
	}

	human.Weight = weight

	return nil
}

// Action is structure that contains Human object
type Action struct {
	Human
}

// NewAction is constructor of Action structure
func NewAction(name string, age int16, height float32, weight float32) *Action {
	return &Action{
		NewHuman(name, age, height, weight),
	}
}

func main() {
	action := NewAction("John", 31, 180, 80)
	fmt.Println(action)

	err := action.SetName("Van")
	if err != nil {
		fmt.Println(err)
	}

	err = action.SetAge(25)
	if err != nil {
		fmt.Println(err)
	}

	err = action.SetHeight(150)
	if err != nil {
		fmt.Println(err)
	}

	err = action.SetWeight(65)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(action)
}
