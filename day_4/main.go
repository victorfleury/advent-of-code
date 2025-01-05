package main

import (
	//"bufio"
	"fmt"
	"os"
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
	Direction{0, 0},
	Direction{0, 1},
	Direction{1, -1},
	Direction{1, 0},
	Direction{1, 1},
}

func main() {
	fmt.Println("Advent of code day 4")
	Soluce("input.txt")
}

func Soluce(filepath string) int {

	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	grid := strings.Split(string(data), "\n")
	fmt.Println(grid)

	var result int
	for i := 0; i < len(grid); i++ {
		columns := len(grid[0])
		for j := 0; j < columns; j++ {
			if string(grid[i][j]) == "X" {
				for _, d := range Search_Directions {
					//fmt.Println(i, j, d)
					if getWord(grid, i, j, d) == "XMAS" {
						result += 1
						fmt.Println("Current result ", result)
					}
				}
			}
		}

	}

	return result
}

func getWord(grid []string, i, j int, d Direction) string {
	word := string(grid[i][j])
	for len(word) < 4 {
		i += d.x
		j += d.y
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return ""
		}
		word += string(grid[i][j])
	}
	//fmt.Println("Word", word)
	return word
}
