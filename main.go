package main

import (
	"fmt"
	//"math/rand"

	algSort "github.com/sasukeruud/GO/Algorithm"
	//test "github.com/sasukeruud/GO/HelloWorld"
)

var (
	Array [20]int
)

func fillArray(arr []int) {
	var (
		number int = 100
	)
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = number

		number--
	}
}

func writeArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println(arr[i], ":")
	}
}

func main() {
	fillArray(Array[:])
	writeArray(Array[:])
	fmt.Println("QuickSort")
	algSort.QuickSort(Array[:], 0, len(Array)-1)
	writeArray(Array[:])

	fmt.Scanln()
}
