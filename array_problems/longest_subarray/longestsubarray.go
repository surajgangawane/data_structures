package main

import "fmt"

/*
Given an array arr[] containing integers and an integer k, your task is to find the length of the longest subarray
where the sum of its elements is equal to the given value k. It is guaranteed that a valid subarray exists.

Examples:

Input: arr[] = [10, 5, 2, 7, 1, 9], k = 15
Output: 4
Explanation: The subarray [5, 2, 7, 1] has a sum of 15 and length 4.

Input: arr[] = [-1, 2, -3], k = -2
Output: 3
Explanation: The subarray [-1, 2, -3] has a sum of -2 and length 3.

Input: arr[] = [1, -1, 5, -2, 3], k = 3
Output: 4
Explanation: The subarray [1, -1, 5, -2] has a sum of 3 and a length 4.

*/

func main() {
	var arr = []int{1, 2, 3, 1, 1, 1, 0, 0, -1, 1, 1, 3, 3}
	getLongestSubarrayLength(arr, 6)
}

func getLongestSubarrayLength(arr []int, k int) {
	maximumLengthArray := 0
	i := 0
	j := 0
	sum := 0

	for i < len(arr) && j < len(arr) {
		if sum <= k {
			sum = sum + arr[j]
		}
		if sum > k {
			sum = sum - arr[i]
			if sum == k && j-i > maximumLengthArray {
				maximumLengthArray = j - i
			}
			i++
		}
		j++
	}

	fmt.Println("Longest subarray: ", maximumLengthArray)
}
