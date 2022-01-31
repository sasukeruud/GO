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
func partitionINT(arr []int, low, high int) ([]int, int) {
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

/*
* Function to partition a string array
 */
func partitionSTR(arr []string, low, high int) ([]string, int) {
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
func QuickSortINT(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partitionINT(arr, low, high)
		arr = QuickSortINT(arr, low, p-1)
		arr = QuickSortINT(arr, p+1, high)
	}
	return arr
}

/*
*Function to do quickSort on a string array
 */
func QuickSortSTR(arr []string, low, high int) []string {
	if low < high {
		var p int
		arr, p = partitionSTR(arr, low, high)
		arr = QuickSortSTR(arr, low, p-1)
		arr = QuickSortSTR(arr, p+1, high)
	}
	return arr
}
