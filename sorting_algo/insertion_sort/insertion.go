package main

import "fmt"

/*

What is Insertion Sort?
Insertion Sort is a simple comparison-based sorting algorithm. It builds the final sorted array one element at a time. It is much like sorting playing cards in your hands, where you pick up one card at a time and place it in the correct position relative to the cards already in your hand.

How Does Insertion Sort Work?
Here’s how the algorithm works step by step:

Start from the second element (because a single-element array is trivially sorted). Assume that the first element is already sorted.

Pick the current element: For each subsequent element in the array (starting from the second one), you will compare it with the elements before it (those that are already "sorted").

Find the correct position: Move all the elements greater than the current element to one position to the right.

Insert the current element: Insert the current element at the correct position.

Repeat this process for each element in the array.

Detailed Example:
Let’s walk through an example where we have an unsorted array:

[
5
,
2
,
9
,
1
,
5
,
6
]
[5,2,9,1,5,6]
We'll sort this using Insertion Sort.

Step-by-Step Process:
Initial array:

[
5
,
2
,
9
,
1
,
5
,
6
]
[5,2,9,1,5,6]
We start with the second element (index 1), which is 2.

First iteration (current element = 2):

Compare 2 with 5 (the first element).
Since 2 is smaller than 5, move 5 one position to the right.
Insert 2 in the first position.
Array after first iteration:
[
2
,
5
,
9
,
1
,
5
,
6
]
[2,5,9,1,5,6]
Second iteration (current element = 9):

Compare 9 with 5. Since 9 is greater than 5, leave 9 in place.
Array after second iteration:
[
2
,
5
,
9
,
1
,
5
,
6
]
[2,5,9,1,5,6]
Third iteration (current element = 1):

Compare 1 with 9 (move 9 one position to the right).
Compare 1 with 5 (move 5 one position to the right).
Compare 1 with 2 (move 2 one position to the right).
Insert 1 in the first position.
Array after third iteration:
[
1
,
2
,
5
,
9
,
5
,
6
]
[1,2,5,9,5,6]
Fourth iteration (current element = 5):

Compare 5 with 9 (move 9 one position to the right).
Insert 5 in the correct position between 2 and 9.
Array after fourth iteration:
[
1
,
2
,
5
,
5
,
9
,
6
]
[1,2,5,5,9,6]
Fifth iteration (current element = 6):

Compare 6 with 9 (move 9 one position to the right).
Insert 6 in the correct position between 5 and 9.
Array after fifth iteration:
[
1
,
2
,
5
,
5
,
6
,
9
]
[1,2,5,5,6,9]
Now the array is fully sorted:

*/

func insertion_sort(unsortedArr []int, n int) []int {
	for i := 1; i <= n-1; i++ {
		for j := i; j > 0; j-- {
			if unsortedArr[j] < unsortedArr[j-1] {
				unsortedArr[j], unsortedArr[j-1] = unsortedArr[j-1], unsortedArr[j]
			} else {
				break
			}
		}
	}

	return unsortedArr
}

func main() {
	var arr = []int{14, 9, 15, 12, 6, 8, 13}

	sortedArr := insertion_sort(arr, len(arr))
	fmt.Println("Sorted array: ", sortedArr)
}
