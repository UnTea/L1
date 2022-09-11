package main

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strings"
)

func examiner(T reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", T.String(), "and kind is", T.Kind())

	switch T.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")

		examiner(T.Elem(), depth+1)

	case reflect.Struct:
		for i := 0; i < T.NumField(); i++ {
			f := T.Field(i)

			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
		}
	}
}

type Human struct {
	Name   string
	Age    int16
	Height float32
	Weight float32
}

func NewHuman(name string, age int16, height float32, weight float32) Human {
	return Human{
		Name:   name,
		Age:    age,
		Height: height,
		Weight: weight,
	}
}

func (human *Human) String() string {
	return fmt.Sprintf("Name=%v\nAge=%v\nHeight=%v\nWeight=%v\n",
		human.Name, human.Age, human.Height, human.Weight)
}

func (human *Human) SetName(name string) error {
	if len(name) <= 0 {
		return errors.New("name must contain at least one letter")
	}

	human.Name = cases.Title(language.Und).String(cases.Lower(language.Und).String(name))

	return nil
}

func (human *Human) SetAge(age int16) error {
	if age < 0 {
		return errors.New("age must be greater than or equal to zero")
	}

	human.Age = age

	return nil
}

func (human *Human) SetHeight(height float32) error {
	if height < float32(0) {
		return errors.New("height must be greater than zero")
	}

	human.Height = height

	return nil
}

func (human *Human) SetWeight(weight float32) error {
	if weight < float32(0) {
		return errors.New("age must be greater than zero")
	}

	human.Weight = weight

	return nil
}

type Action struct {
	Human
}

func NewAction(name string, age int16, height float32, weight float32) *Action {
	return &Action{
		NewHuman(name, age, height, weight),
	}
}

func main() {
	action := NewAction("John", 31, 180, 80)
	actionType := reflect.TypeOf(action)

	examiner(actionType, 0)
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

	actionType = reflect.TypeOf(action)
	examiner(actionType, 0)
	fmt.Println(action)
}
