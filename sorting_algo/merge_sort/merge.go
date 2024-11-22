package main

import "fmt"

/*
Merge sort is sorting technique with follow divide and merge rule.
Steps to perform merge sort
1. Divide entire array to a single element/node.
2. Then perfom merging of two array at same level.

For ex.

Array [3, 1, 2, 4, 1, 5, 6, 2, 4]

*Divide
**First iteration
[3,1,2,4,1] [5,6,2,4]

**Second iteration
[3,1,2] [4,1] [5,6] [2,4]

**Third iteration
[3,1] [2] [4] [1] [5] [6] [2] [4]

** Fourth iteration
[3] [1]

*Mering: While merging we have to merge elements in correct order asc/desc as expected.
Lets consider we have to place elements in ascending order.

We will create/have temp array which will have sorted elements

Now we will start comparing elements from both array and will place smaller element first in temp array.

so here 1 is smaller than 3, so 1 will be on 0 position and 3 will be on 1st position in temp array

then copy these elements from temp array to original array at its respective position

array after merging will be

[1,3] [2]

Now we will compare above two array and result after merging will be
[1,2,3]

Now we will merge array [4] and [1]. Result after merging will be
[1,4]

now from left side we have two array to be merged.
first: [1, 2, 3]
second: [1, 4]

after merging result will be
[1, 1, 2, 3, 4]

Similar process will be follow on right side. Array will be
[2, 4, 5, 6]

Now to get final sorted array we will need to merge two arrays i.e [1, 1, 2, 3, 4] & [2, 4, 5, 6]

[1, 1, 2, 2, 3, 4, 4, 5, 6]

Time complexity will be = O(n log base 2 n)

n = max n number of times will need to iterate to merge element in one thats too for last merge. For other its can be
n / 2, n / 4, n / 8. We have to consider worst one so n.

log base 2 n = In divide process, we divide array into two which comes to n / 2, n /4, n / 8. Whenever there is n / 2 happening
we have to consider complexity log base 2 n. Its also step will required to divide array until its one.

If length of array is 16 and log base 2 16 comes out to be 4. So it will required four iteration/steps to break down
array into single element

Space complexity = O(n) = temp array
*/

func sort(arr []int, low, high int) {
	if low == high {
		return
	}
	mid := (low + high) / 2

	sort(arr, low, mid)
	sort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func merge(arr []int, low, mid, high int) {
	var temp []int
	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if arr[left] < arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

func main() {
	unsortArray := []int{3, 1, 2, 4, 1, 5, 6, 2, 4}

	sort(unsortArray, 0, len(unsortArray)-1)

	fmt.Println("Array sorted using merge sort: ", unsortArray)
}
