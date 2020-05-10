package main

// Sum calculates Sum of a single array
func Sum(nums []int) (expectedSum int) {
	for _, number := range nums {
		expectedSum += number
	}
	return
}

// SumAll calculates sum of provided array
func SumAll(nums ...[]int) []int {
	var sum []int
	for _, numarr := range nums {
		sum = append(sum, Sum(numarr))
	}
	return sum
}
