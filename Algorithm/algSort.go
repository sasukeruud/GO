package algSort

//"fmt"

//"fmt"

//Function to swap two int
func swapInt(a, b int) {
	t := a
	a = b
	b = t
}

/* Partition function for quickSort
* From:: https://qvault.io/golang/quick-sort-golang/
 */
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

/* The main function that implements QuickSort
* arr[] --> Array to be sorted,
* low --> Starting index,
* high --> Ending index
* From:: https://qvault.io/golang/quick-sort-golang/ */
func QuickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}
