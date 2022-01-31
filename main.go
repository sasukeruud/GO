package main

import (
	"fmt"
	//"math/rand"

	algSort "github.com/sasukeruud/GO/Algorithm"
	//test "github.com/sasukeruud/GO/HelloWorld"
)

var (
	intArray [20]int
	strArray [5]string
)

func fillArrayINT(arr *[20]int) {
	var (
		number int = 100
	)
	for i := 0; i < len((*arr)); i++ {
		(*arr)[i] = number

		number--
	}
}

func fillArraySTR(arr *[5]string) {
	(*arr)[0] = "m"
	(*arr)[1] = "v"
	(*arr)[2] = "p"
	(*arr)[3] = "w"
	(*arr)[4] = "a"
	//(*arr)[5] = "c"
}

func writeArray(arr [20]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], ":")
	}
}

func writeArraySTR(arr [5]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], ":")
	}
}

func main() {
	fillArrayINT(&intArray)
	writeArray(intArray)
	fmt.Println("QuickSort")
	algSort.QuickSortINT(intArray[:], 0, len(intArray)-1)
	writeArray(intArray)
	fmt.Println("String array")
	fillArraySTR(&strArray)
	writeArraySTR(strArray)
	algSort.QuickSortSTR(strArray[:], 0, len(strArray)-1)
	fmt.Println("Sorted")
	writeArraySTR(strArray)

	fmt.Scanln()
}
