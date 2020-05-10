package main

import "testing"

func TestSum(t *testing.T) {
	nums := []int{2, 3, 4, 5}
	expectedSum := 14
	actualSum := Sum(nums)
	if actualSum != expectedSum {
		t.Errorf("expected %d , but found %d", expectedSum, actualSum)
	}
}
