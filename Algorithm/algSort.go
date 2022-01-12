package algSort

//"fmt"

//Function to swap two int
func swapInt(a, b int) {
	var t int = a
	a = b
	b = t
}

/*
* Partition function for quickSort
 */
func partition(arr []int, low, high int) int {
	var (
		piviot int = arr[high]
		i      int = low - 1
	)

	for j := low; j <= high-1; j++ {
		if arr[j] < piviot {
			i++
			swapInt(arr[i], arr[j])
		}
	}

	swapInt(arr[i+1], arr[high])
	return i + 1
}

/* The main function that implements QuickSort
arr[] --> Array to be sorted,
low --> Starting index,
high --> Ending index */
func QuickSort(arr []int, low, high int) {
	var (
		pi int = partition(arr, low, high)
	)

	if low < high {
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}
