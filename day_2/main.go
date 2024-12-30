package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	//"slices"
	"strings"
)

func main() {
	fmt.Println("Advent of code day 2")

	// Open file
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	fmt.Println("Opening input.txt")
	scanner := bufio.NewScanner(data)

	var validReports int

	// Rules :
	//The levels are either all increasing or all decreasing.
	//Any two adjacent levels differ by at least one and at most three.

	for scanner.Scan() {
		currReport := strings.Split(scanner.Text(), " ")
		if IsValidReport(currReport) {
			validReports += 1
		}
		//if !slices.IsSorted(currReportValues) {
		//continue
		//}
		//validReports += 1
		////for value := range currReportValues {

		////}

	}
	//fmt.Println("Valid reports : ", validReports)
}

func IsValidReport(report []string) bool {
	if !IsSorted(report) || !IsSafeDistance(report) {
		return false
	}
	return false
}

func IsSorted(report []string) bool {
	if slices.IsSorted(report) {
		return true
	}

	sorted := slices.Clone(report)
	slices.Sort(sorted)
	slices.Reverse(sorted)

	return slices.Equal(sorted, report)

}

func IsSafeDistance(report []string) bool {
	previous_value := -1
	for _, value := range report {
		if previous_value == -1 {
			previous_value, _ = strconv.Atoi(value)
			continue
		}
		value, _ := strconv.Atoi(value)
		difference := previous_value - value

		if !(1 <= difference && difference <= 3) {
			return false
		}
		previous_value = value

	}
	return true
}
