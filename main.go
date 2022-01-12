package main

import (
	"Algorithm/algSort"
	"fmt"
)

var (
	Array [20]int
)

func fillArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.int
	}
}

func writeArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}

func main() {
	fillArray()
	writeArray(Array)
	algSort.QuickSort(Array, 0, len(Array))
	writeArray(Array)
}
