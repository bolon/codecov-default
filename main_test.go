package main

import "testing"

func TestSumLarge(t *testing.T) {
	var expected = true
	var actual = sum(10, 5)
	if expected != actual {
		t.Errorf("%v == %v\n", expected, actual)
	}
}