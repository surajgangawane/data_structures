package main

import "fmt"

/*
### **Selection Sort Explanation with Example**

**Selection Sort** is a simple comparison-based sorting algorithm that works by repeatedly finding the smallest (or largest, depending on sorting order) element from the unsorted part of the array and swapping it with the first unsorted element.

### **How Selection Sort Works:**
1. **Start with the entire array** and find the smallest element in the array.
2. **Swap the smallest element** with the element at the first position.
3. Now, the first element is considered "sorted." So, move the boundary of the unsorted part to the right by one position.
4. Repeat the process for the remaining unsorted portion of the array until the entire array is sorted.

### **Steps of Selection Sort:**
- **Step 1:** Find the smallest element in the entire unsorted part of the array.
- **Step 2:** Swap this smallest element with the first element of the unsorted portion.
- **Step 3:** Move the boundary of the unsorted part one element to the right.
- **Step 4:** Repeat this until the unsorted portion is reduced to just one element.

### **Example:**

Let's walk through the sorting of an unsorted array using Selection Sort.

Consider the array:

```
[64, 25, 12, 22, 11]
```

### **Pass 1:**
- **Unsorted Array:** `[64, 25, 12, 22, 11]`
- Find the smallest element in the entire array, which is `11`.
- Swap `11` with the first element (`64`).
  - Array after swap: `[11, 25, 12, 22, 64]`
- Now, the first element `11` is in its correct sorted position.

### **Pass 2:**
- **Unsorted Array:** `[25, 12, 22, 64]` (start from index 1)
- Find the smallest element in the unsorted portion (`25, 12, 22, 64`), which is `12`.
- Swap `12` with `25`.
  - Array after swap: `[11, 12, 25, 22, 64]`
- Now, `12` is in its correct position.

### **Pass 3:**
- **Unsorted Array:** `[25, 22, 64]` (start from index 2)
- Find the smallest element in the unsorted portion (`25, 22, 64`), which is `22`.
- Swap `22` with `25`.
  - Array after swap: `[11, 12, 22, 25, 64]`
- Now, `22` is in its correct position.

### **Pass 4:**
- **Unsorted Array:** `[25, 64]` (start from index 3)
- Find the smallest element in the unsorted portion (`25, 64`), which is `25`.
- No swap is needed since `25` is already in its correct position.

### **Pass 5:**
- **Unsorted Array:** `[64]` (only one element left)
- No further sorting is needed because only one element remains.

### **Final Sorted Array:**
```
[11, 12, 22, 25, 64]
```

### **Time Complexity of Selection Sort:**
- **Worst-case time complexity:** O(n²)
- **Best-case time complexity:** O(n²)
- **Average-case time complexity:** O(n²)

  Regardless of the initial order of elements, Selection Sort always makes O(n²) comparisons because it scans the entire unsorted portion of the array in each pass.

### **Space Complexity of Selection Sort:**
- **Space complexity:** O(1)

  Selection Sort is an **in-place** sorting algorithm, meaning it doesn't require any extra memory besides a constant amount for variables like the `minIndex`. Therefore, its space complexity is O(1).

### **Summary of Selection Sort:**
- **Selection Sort** is easy to understand and implement.
- It always performs **n²** comparisons, regardless of whether the array is already sorted.
- It is inefficient for large datasets because of its quadratic time complexity.
*/

// selection_sort function expected unsorted array and length of array as a input
func selection_sort(unsortedArr []int, n int) []int {
	for i := 0; i <= n-2; i++ {
		for j := i + 1; j <= n-1; j++ {
			if unsortedArr[j] < unsortedArr[i] {
				unsortedArr[i], unsortedArr[j] = unsortedArr[j], unsortedArr[i]
			}
		}
	}

	return unsortedArr
}

func main() {
	var unsortedArr = []int{13, 46, 24, 52, 20, 9}

	fmt.Println("Result of selection sort: ", selection_sort(unsortedArr, len(unsortedArr)))
}

// Time complexity for selection sort is as below
// Best: O(n^2)
// Average: O(n^2)
// Worst: O(n^2)
