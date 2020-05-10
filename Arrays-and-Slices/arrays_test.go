package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{2, 3, 4, 5}
	expectedSum := 14
	actualSum := Sum(nums)
	if actualSum != expectedSum {
		t.Errorf("expected %d , but found %d", expectedSum, actualSum)
	}
}

func TestSumAll(t *testing.T) {
	actualSum := SumAll([]int{1, 3, 4}, []int{0, 2, 1})
	expectedSum := []int{8, 3}
	if !reflect.DeepEqual(actualSum, expectedSum) {
		t.Errorf("Expected %v got %v", expectedSum, actualSum)
	}
}
