package main

import (
	"fmt"
	"reflect"
)

func TypeOf(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("%v type(%T)\n", i, i)
	case int:
		fmt.Printf("%v type(%T)\n", i, i)
	case bool:
		fmt.Printf("%vntype(%T)\n", i, i)
	case chan int:
		fmt.Printf("%v type(%T)\n", i, i)
	}
}

func Reflect(i interface{}) {
	ri := reflect.TypeOf(i)

	fmt.Printf("Type name: %v, Specific kind: %v, String representation: %v\n", ri.Name(), ri.Kind(), ri.String())
}

func main() {
	str := "Pineapple"
	integer := 420
	boolean := true
	channel := make(chan int)

	TypeOf(str)
	TypeOf(integer)
	TypeOf(boolean)
	TypeOf(channel)

	Reflect(str)
	Reflect(integer)
	Reflect(boolean)
	Reflect(channel)
}
