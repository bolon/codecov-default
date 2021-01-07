package main

import "testing"

func TestSumLarge(t *testing.T) {
	var expected = true
	var actual = sum(10, 5)
	if expected != actual {
		t.Errorf("%v == %v\n", expected, actual)
	}
}

// func TestMain(t *testing.T) {
//   main()
// }

// func TestNonCondi(t *testing.T) {
//   var expected = 00
//   var actual = nonCondi(0)
//   if expected != actual {
//     t.Errorf("%v == %v\n", expected, actual)
//   }
// }