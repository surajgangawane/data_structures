package main

import "fmt"

func main() {
	arr := []int{1, 8, 7, 56, 90}
	largestNumber := getLargestNumber(arr)
	secondLargest := getSecondLargestElem(arr)

	fmt.Println("Largest number in arr: ", largestNumber)

	if secondLargest == -1 {
		fmt.Println("There is no second largest number")
	} else {
		fmt.Println("Second largest number in arr: ", secondLargest)
	}
}

// Note: One solution for this will be first you can sort array and then pick last element.
// Complexity for sorting algo like merge and quick will ne O(n log n) but for our approach is O(n) which is far better.
func getLargestNumber(arr []int) int {
	largestNumber := arr[0]
	for _, v := range arr {
		if v > largestNumber {
			largestNumber = v
		}
	}

	return largestNumber
}

func getSecondLargestElem(arr []int) int {
	if len(arr) == 1 {
		return -1
	}
	largest := arr[0]
	sLargest := -1

	for _, v := range arr {
		if v > largest {
			sLargest, largest = largest, v
		}
		if v > sLargest && v != largest {
			sLargest = v
		}
	}

	return sLargest
}
