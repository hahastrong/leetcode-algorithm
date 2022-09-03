package greedyAlgorithm

import (
	"fmt"
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	fmt.Println(findContentChildren([]int{1,2,3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1,2}, []int{1, 2, 3}))
}


func Test_wiggleMaxLength(t *testing.T) {
	fmt.Println(wiggleMaxLength([]int{1,7,4,9,2,5}))
	fmt.Println(wiggleMaxLength([]int{1,17,5,10,13,15,10,5,16,8}))
	fmt.Println(wiggleMaxLength([]int{1,2,3,4,5,6,7,8,9}))
}

func Test_maxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}

func Test_canJump(t *testing.T) {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
	fmt.Println(canJump([]int{5,9,3,2,1,0,2,3,3,1,0,0}))
}

func Test_jump(t *testing.T) {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{2,3,0,1,4}))
	fmt.Println(jump([]int{2,3,1}))
}

func Test_largestSumAfterKNegations(t *testing.T) {
	fmt.Println(largestSumAfterKNegations([]int{4,2,3}, 1))
	fmt.Println(largestSumAfterKNegations([]int{3,-1,0,2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{2,-3,-1,5,-4}, 2))
}

func Test_canCompleteCircuit(t *testing.T) {
	fmt.Println(canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2}))
	fmt.Println(canCompleteCircuit([]int{2,3,4}, []int{3,4,3}))
	fmt.Println(canCompleteCircuit([]int{4}, []int{5}))
}

func Test_candy(t *testing.T) {
	fmt.Println(candy([]int{1,0,2}))
	fmt.Println(candy([]int{1,2,2}))
}

func Test_lemonadeChange(t *testing.T) {
	fmt.Println(lemonadeChange([]int{5,5,5,10,20}))
	fmt.Println(lemonadeChange([]int{5,5,10,10,20}))
}

func Test_reconstructQueue(t *testing.T) {
	fmt.Println(reconstructQueue([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}))
	fmt.Println(reconstructQueue([][]int{{6,0},{5,0},{4,0},{3,2},{2,2},{1,4}}))
}

func Test_findMinArrowShots(t *testing.T) {
	fmt.Println(findMinArrowShots([][]int{{10,16},{2,8},{1,6},{7,12}}))
	fmt.Println(findMinArrowShots([][]int{{1,2},{3,4},{5,6},{7,8}}))
}

func Test_eraseOverlapIntervals(t *testing.T) {
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3},{3,4},{1,3}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{1,2},{1,2}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3}}))
}

func Test_partitionLabels(t *testing.T) {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
	fmt.Println(partitionLabels("caedbdedda"))
}

func Test_merge(t *testing.T) {
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
	fmt.Println(merge([][]int{{1,4},{4,5}}))
}

func Test_monotoneIncreasingDigits(t *testing.T) {
	fmt.Println(monotoneIncreasingDigits(10))
	fmt.Println(monotoneIncreasingDigits(1234))
	fmt.Println(monotoneIncreasingDigits(332))
	fmt.Println(monotoneIncreasingDigits(100000023))
}


func Test_minCameraCover(t *testing.T) {
	fmt.Println(minCameraCover([][]int{{1,3},{2,6},{8,10},{15,18}}))
	fmt.Println(minCameraCover([][]int{{1,4},{4,5}}))
}
