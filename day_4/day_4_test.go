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
