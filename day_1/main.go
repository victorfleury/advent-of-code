package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of code day 1")

	// Open file
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	list1, list2 := []int{}, []int{}

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")
		//fmt.Println(values[0], "testst ", values[1])
		val1, _ := strconv.Atoi(values[0])
		val2, _ := strconv.Atoi(values[1])
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	// Sort slices
	sort.Ints(list1)
	sort.Ints(list2)

	sums := []int{}

	for i, _ := range list1 {
		fmt.Println(i, list1[i], list2[i])

		sums = append(sums, int(math.Abs(float64(list1[i]-list2[i]))))

	}

	var result int

	for _, v := range sums {
		result += v
	}

	fmt.Println(result)

}
