package main

import (
	"testing"
)

func TestDecodeText(t *testing.T) {
	result := decodeText("LLRR=")
	if result != "210122" {
		t.Fatalf("Expect=%s, result=%s\n", "210122", result)
	}
}
