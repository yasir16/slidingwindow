
package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

func maxSlidingWindow(nums []int, k int) []int {
	// Start with a window of length zero
	left, right := 0, 0

	// We will use this array as a deque ( a queue that allows both front & back pop)
	// This is a monotonic increasing deque, meaning smaller elements will be popped
	// This is similar to monotonic increasing stack, but can do front pop
	deque := []int{}

	res := []int{}
	for right < len(nums) {
		// If there are elements on top of this deque smaller than current element
		for len(deque) > 0 && nums[right] > nums[deque[len(deque)-1]] {
			// Pop the smaller element
			deque = deque[:len(deque)-1]
		}

		// Append the new element's index
		// We are appending index and not the element itself cause
		// With Index we can make judgement about our window length
		deque = append(deque, right)

		// If left has moved beyond the index of first element, pop
		if left > deque[0] {
			deque = deque[1:]
		}

		// Check if we have a valid window
		if right+1 >= k {
			// Append the first number from deque to the result
			// The first number will always be the largest
			res = append(res, nums[deque[0]])
			left++
		}
		right++
	}

	return res
}
