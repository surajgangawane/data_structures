package main

import "fmt"

func main() {
	unsortedItems := []int{4, 6, 2, 5, 7, 9, 5, 3}
	sort(unsortedItems, 0, len(unsortedItems)-1)
	fmt.Println("Here is sortedarray: ", unsortedItems)
}

func sort(arr []int, low, high int) {
	if low < high {
		partitionedIndex := quickSort(arr, low, high)
		sort(arr, low, partitionedIndex-1)
		sort(arr, partitionedIndex+1, high)
	}
}

func quickSort(arr []int, low, high int) int {
	pivot := arr[low]

	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}

		for arr[j] > pivot && j >= low+1 {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]
	return j
}
