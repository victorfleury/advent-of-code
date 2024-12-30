package main

import (
	"testing"
)

//func TestValidreports(t *testing.T) {

//input := []string{"8", "10", ",13", "14", "16"}

//got := IsValidReport(input)
//want := true

//if got != want {
//t.Errorf("Got : %v want %v, givent %v", got, want, input)
//}
//}

func TestIsSorted(t *testing.T) {
	t.Run("Test is sorted on sorted input", func(t *testing.T) {
		input := []int{1, 2, 3}
		got := IsSorted(input)
		want := true
		if got != want {
			t.Errorf("Got %v want %v given %v", got, want, input)
		}
	})

	t.Run("Test is sorted on unsorted input", func(t *testing.T) {
		input := []int{3, 2, 1}
		got := IsSorted(input)
		want := true
		if got != want {
			t.Errorf("Got %v, want %v given %v", got, want, input)
		}
	})

	t.Run("Test is sorted on unsorted input", func(t *testing.T) {
		input := []int{3, 1, 2}
		got := IsSorted(input)
		want := false
		if got != want {
			t.Errorf("Got %v, want %v given %v", got, want, input)
		}
	})
}

func TestIsSafeDistance(t *testing.T) {

	t.Run("Test is safe distance", func(t *testing.T) {
		input := []int{7, 6, 4, 2, 1}
		got := IsSafeDistance(input)
		want := true
		if got != want {
			t.Errorf("Got %v, want %v given %v", got, want, input)
		}
	})
	t.Run("Test is unsafe distance", func(t *testing.T) {
		input := []int{1, 2, 7, 8, 9}
		got := IsSafeDistance(input)
		want := false
		if got != want {
			t.Errorf("Got %v, want %v given %v", got, want, input)
		}
	})

}

func TestIsSafe(t *testing.T) {
	input := []int{8, 6, 4, 4, 1}
	got := IsSafe(input)
	want := true

	if got != want {
		t.Errorf("Got %v, want %v given %v", got, want, input)
	}
}

func TestSoluce(t *testing.T) {
	//input := "/home/victor/code/advent-of-code/day_2/input.txt"
	input := "/home/victor/code/advent-of-code/day_2/input_test.txt"
	got := Soluce(input)
	want := 5

	if got != want {
		t.Errorf("Got %v want %v given %v", got, want, input)
	}
}
