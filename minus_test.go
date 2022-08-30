package main

import (
	"fmt"
	"testing"
)

func TestMinus(t *testing.T) {

	upLimit := (2 << 31 -1) / 10
	lowLimit := - (2 << 31) / 10

	fmt.Printf("upLimit: %d, lowLimit: %d", upLimit, lowLimit)
}