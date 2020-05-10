package main

// Sum calculates Sum of a single array
func Sum(nums []int) (expectedSum int) {
	for _, number := range nums {
		expectedSum += number
	}
	return
}
