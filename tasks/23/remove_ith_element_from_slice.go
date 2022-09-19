package main

import "fmt"

func removeWithSaveOrder(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

func removeWithoutSavingOrder(a []int, i int) []int {
	a[i] = a[len(a)-1]

	return a[:len(a)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)
	fmt.Println(removeWithSaveOrder(arr, 3))
	//fmt.Println(removeWithoutSavingOrder(arr, 3))
}
