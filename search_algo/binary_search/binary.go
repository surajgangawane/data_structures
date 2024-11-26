package main

import "fmt"

func main() {
	var arr = []int{3, 4, 6, 7, 9, 12, 16, 17}
	const target = 6
	index := iterativeBinarySearch(arr, target)

	recursiveIndex := recursiveBinarySearch(arr, target, 0, len(arr)-1)

	if index >= 0 {
		fmt.Printf("At index: %d element: %d is present\n", index, target)
	} else {
		fmt.Println("Element not present in array")
	}

	if recursiveIndex >= 0 {
		fmt.Printf("recursiveIndex | At index: %d element: %d is present\n", recursiveIndex, target)
	} else {
		fmt.Println("recursiveIndex | Element not present in array")
	}
}

func iterativeBinarySearch(arr []int, target int) int {
	var low, high = 0, len(arr)

	for low <= high && target <= arr[high-1] && target >= arr[low] {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func recursiveBinarySearch(arr []int, target, low, high int) int {
	if low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			return recursiveBinarySearch(arr, target, low, mid-1)
		} else {
			return recursiveBinarySearch(arr, target, mid+1, high)
		}
	}

	return -1
}
