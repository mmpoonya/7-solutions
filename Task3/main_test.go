package main

import (
	"reflect"
	"testing"
)

func TestGetBeef(t *testing.T) {
	result := getBeef("Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.")
	expect := map[string]int{
		"t-bone":   4,
		"fatback":  1,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
	}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("Expect=%v, result=%v\n", expect, result)
	}
}
