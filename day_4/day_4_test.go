package main

import "testing"

func TestSoluce(t *testing.T) {
	input := "input_test.txt"
	got := Soluce(input)
	want := 18

	if got != want {
		t.Errorf("got %v want %v with input %v", got, want, input)
	}
}

func TestSolucePart2(t *testing.T) {
	input := "input_test.txt"
	got := SolucePart2(input)
	want := 9

	if got != want {
		t.Errorf("got %v want %v with input %v", got, want, input)
	}
}
