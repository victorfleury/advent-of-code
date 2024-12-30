package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"

	//"slices"
	"strings"
)

func main() {
	fmt.Println("Advent of code day 2")
	input := "input.txt"
	Soluce(input)
}

func Soluce(filepath string) int {
	var validReports int
	// Open file
	data, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	fmt.Println("Opening ", filepath)
	scanner := bufio.NewScanner(data)

	// Rules :
	//The levels are either all increasing or all decreasing.
	//Any two adjacent levels differ by at least one and at most three.

	for scanner.Scan() {
		currReport := strings.Split(scanner.Text(), " ")
		var convertedReport []int
		for _, i := range currReport {
			i, _ := strconv.Atoi(i)
			convertedReport = append(convertedReport, i)
		}
		if IsValidReport(convertedReport) {
			validReports += 1
		}

	}
	fmt.Println("Valid reports : ", validReports)
	return validReports
}

func IsSafe(report []int) bool {
	for i := 0; i < len(report); i++ {
		clone := slices.Clone(report)
		subset := slices.Delete(clone, i, i+1)
		if IsSorted(subset) && IsSafeDistance(subset) {
			return true
		}
	}
	return false
}

func IsValidReport(report []int) bool {
	if IsSorted(report) && IsSafeDistance(report) {
		return true
	} else {
		return IsSafe(report)
	}
}

func IsSorted(report []int) bool {
	if slices.IsSorted(report) {
		return true
	}

	sorted := slices.Clone(report)
	slices.Sort(sorted)
	slices.Reverse(sorted)
	return slices.Equal(sorted, report)

}

func IsSafeDistance(report []int) bool {
	previous_value := -1
	for _, value := range report {
		if previous_value == -1 {
			previous_value = value
			continue
		}
		value := value
		difference := int(math.Abs(float64(previous_value - value)))

		if !(1 <= difference && difference <= 3) {
			// Difference is not in range
			return false
		}
		previous_value = value

	}
	return true
}
