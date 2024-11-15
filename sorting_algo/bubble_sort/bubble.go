package main

import "fmt"

/*
### **Bubble Sort Explanation with Example**

Bubble Sort is a simple comparison-based sorting algorithm. The main idea is to repeatedly step through the list, compare adjacent elements, and swap them if they are in the wrong order. This process is repeated for each element in the array, gradually "bubbling" the largest unsorted element to the end of the array on each pass.

#### **Steps of Bubble Sort:**
1. Compare the first two adjacent elements of the array.
2. If the first element is greater than the second, swap them.
3. Move to the next pair of adjacent elements and repeat the comparison and swap if necessary.
4. After one pass through the entire array, the largest element will be at the end of the array.
5. Repeat the process for the remaining unsorted part of the array (i.e., exclude the last element that is already in its correct place).
6. Continue until no more swaps are needed.

### **Bubble Sort Algorithm in Detail:**
For an array of size `n`, the algorithm performs `n-1` passes, where each pass moves the largest unsorted element to its correct position at the end.

### **Example:**

Consider the array:
```
[5, 1, 4, 2, 8]
```

### **Pass 1:**
- Compare `5` and `1`: Since `5 > 1`, swap them.
  - Array becomes: `[1, 5, 4, 2, 8]`
- Compare `5` and `4`: Since `5 > 4`, swap them.
  - Array becomes: `[1, 4, 5, 2, 8]`
- Compare `5` and `2`: Since `5 > 2`, swap them.
  - Array becomes: `[1, 4, 2, 5, 8]`
- Compare `5` and `8`: No swap needed since `5 < 8`.
  - Array stays: `[1, 4, 2, 5, 8]`

After Pass 1, the largest element `8` is in its correct position at the end.

### **Pass 2:**
- Compare `1` and `4`: No swap needed since `1 < 4`.
  - Array stays: `[1, 4, 2, 5, 8]`
- Compare `4` and `2`: Since `4 > 2`, swap them.
  - Array becomes: `[1, 2, 4, 5, 8]`
- Compare `4` and `5`: No swap needed since `4 < 5`.
  - Array stays: `[1, 2, 4, 5, 8]`

After Pass 2, the second largest element `5` is in its correct position.

### **Pass 3:**
- Compare `1` and `2`: No swap needed since `1 < 2`.
  - Array stays: `[1, 2, 4, 5, 8]`
- Compare `2` and `4`: No swap needed since `2 < 4`.
  - Array stays: `[1, 2, 4, 5, 8]`

After Pass 3, the array is sorted, so the algorithm can stop here.

### **Final Sorted Array:**
[1, 2, 4, 5, 8]

### **Time Complexity Recap:**
- **Worst-case:** O(n²), when the array is sorted in reverse order.
- **Best-case:** O(n), when the array is already sorted (with optimization).
- **Average-case:** O(n²).
*/

// bubble_sort function expected unsorted array and length of array as a input
func bubble_sort(unsortedArr []int, n int) []int {
	var swapped = false
	for lastIndex := n - 1; lastIndex >= 1; lastIndex-- {
		for i := 0; i < lastIndex; i++ {
			if unsortedArr[i] > unsortedArr[i+1] {
				swapped = true
				unsortedArr[i], unsortedArr[i+1] = unsortedArr[i+1], unsortedArr[i]
			}
		}

		if !swapped {
			break
		}
	}

	return unsortedArr
}

func main() {
	//var unsortedArr = []int{13, 46, 24, 52, 20, 9}
	var unsortedArr = []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Result of selection sort: ", bubble_sort(unsortedArr, len(unsortedArr)))
}
