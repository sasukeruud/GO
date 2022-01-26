package main

import (
	"fmt"
	//"math/rand"

	algSort "github.com/sasukeruud/GO/Algorithm"
	//test "github.com/sasukeruud/GO/HelloWorld"
)

var (
	array [20]int
)

func fillArray(arr *[20]int) {
	var (
		number int = 99
	)
	for i := 0; i < len((*arr))-2; i++ {
		(*arr)[i] = number

		number--
	}
}

func writeArray(arr [20]int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println(arr[i], ":")
	}
}

func main() {
	fillArray(&array)
	writeArray(array)
	fmt.Println("QuickSort")
	algSort.QuickSort(array[:], 0, len(array)-1)
	writeArray(array)

	fmt.Scanln()
}
