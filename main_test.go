package main

import "testing"

func TestCompare(t *testing.T) {
	var expected = false
	var actual = compareAndSum(0, 5)
	if expected != actual {
		t.Errorf("%v == %v\n", expected, actual)
	}
}