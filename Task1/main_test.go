package main

import (
	"testing"
)

func TestGetMaxPath(t *testing.T) {
	input := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	result := getMaxPath(input)
	if result != 237 {
		t.Fatalf("Expect=%d, result=%d\n", 237, result)
	}
}
