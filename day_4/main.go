package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Direction struct {
	x, y int
}

var Search_Directions = []Direction{
	Direction{-1, -1},
	Direction{-1, 0},
	Direction{-1, 1},
	Direction{0, -1},
	Direction{0, 1},
	Direction{1, -1},
	Direction{1, 0},
	Direction{1, 1},
}

var New_Directions = []Direction{
	Direction{-1, -1},
	Direction{-1, 1},
	Direction{1, -1},
	Direction{1, 1},
}

func main() {
	fmt.Println("Advent of code day 4")
	result := Soluce("input.txt")
	fmt.Println("Result Part 1", result)
	result = SolucePart2("input.txt")
	fmt.Println("Result Part 2", result)
}

func Soluce(filepath string) int {

	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// Split and strip any empty lines
	grid := slices.DeleteFunc(
		strings.Split(string(data), "\n"),
		func(e string) bool {
			return e == ""
		},
	)

	var result int
	for row := 0; row < len(grid)-1; row++ {
		for col := 0; col < len(grid[0])-1; col++ {
			if string(grid[row][col]) == "X" {
				for _, d := range Search_Directions {
					word := getWord(grid, row, col, d)
					if word == "XMAS" {
						result += 1
					}
				}
			}
		}

	}
	return result
}

func getWord(grid []string, row, col int, d Direction) string {
	word := string(grid[row][col])
	for len(word) < 4 {
		row += d.x
		col += d.y
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return ""
		}
		word += string(grid[row][col])
	}
	return word
}

func SolucePart2(filepath string) int {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// Split and strip any empty lines
	grid := slices.DeleteFunc(
		strings.Split(string(data), "\n"),
		func(e string) bool {
			return e == ""
		},
	)
	var result int
	for row := range grid {
		for col := range grid[0] {
			if string(grid[row][col]) != "A" {
				continue
			}
			if getX(grid, row, col) {
				result += 1
			}
		}
	}
	return result
}

func getX(grid []string, row, col int) bool {

	if row <= 0 || row >= len(grid)-1 || col <= 0 || col >= len(grid[0])-1 {
		return false
	}
	top_left := grid[row-1][col-1]
	top_right := grid[row-1][col+1]
	bottom_left := grid[row+1][col-1]
	bottom_right := grid[row+1][col+1]

	if isPattern(top_left, bottom_right) && isPattern(top_right, bottom_left) {
		return true
	}
	return false
}

// Helper function to check if two characters form the pattern "MS" or "SM"
func isPattern(a, b byte) bool {
	return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
}
