package main

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	nums := []int{
		1, 3, 5, 7, 9, 8, 6, 4, 2, 0,
	}

	InsertSort(nums)

	fmt.Println(nums)
}

func TestSelectionSort(t *testing.T) {
	nums := []int{
		1, 3, 5, 7, 9, 8, 6, 4, 2, 0, 4, 2, 3,
	}

	SelectionSort(nums)
	fmt.Println(nums)
}

var nums = []int{
	1, 3, 5, 7, 9, 8, 6, 4, 2, 0, 4, 2, 3,
}

var res = []int {
	0, 1, 2, 2, 3, 3, 4, 4, 5, 6, 7, 8, 9,
}

func TestBubbleSort(t *testing.T) {
	testNum := nums[:]
	BubbleSort(testNum)
	fmt.Println(testNum)
}

func TestShellSort(t *testing.T) {
	testNum := nums[:]
	ShellSort(testNum)
	fmt.Println(testNum)
}

func Equals(actual, expect []int) bool {
	if len(actual) != len(expect) {
		return false
	}

	for i:=0; i<len(actual); i++ {
		if actual[i] != expect[i] {
			return false
		}
	}
	return true
}

func TestByName(t *testing.T) {
	for _, sort := range []struct {
		name string
		function func([]int)
	}{
		{"ShellSort", ShellSort},
		{"InsertSort", InsertSort},
		{"Merge", MergeSort},
		{"MergeBU", MergeBUSort},
		{"QuickSort", QuickSort},
		{"Quick3way", QuickSort3way},
	} {
		testNum := make([]int, len(nums))
		copy(testNum, nums)
		sort.function(testNum)
		t.Logf("%s ...\n", sort.name)
		if !Equals(testNum, res) {
			t.Errorf("failed to sort slice: [%s]: %v\n", sort.name, testNum)
		}
	}
	testNum := make([]int, len(nums))
	copy(testNum, nums)
	actual := HeapSort(testNum)
	t.Logf("%s ...\n", "HeapSort")
	if !Equals(actual, res) {
		t.Errorf("failed to sort slice: [%s]: %v\n", "HeapSort", actual)
	}


}
