package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	// Read input
	runRealInput := false
	if len(os.Args) > 1 {
		runRealInput = true
	}

	filename := "example_input.txt"
	if runRealInput {
		filename = "test_input.txt"
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		panic(err)
	}

	stringData := string(data)
	stringDataRows := strings.Split(stringData, "\n")

	// ### Day 2 ###
	fmt.Println("\n### Day 2 ###\n ")

	// Setup

	listOfRows := make([][]int, 0)

	for _, value := range stringDataRows {
		splitVal := strings.Split(value, " ")
		row := make([]int, 0)
		for i := range len(splitVal) {
			intValue, _ := strconv.Atoi(splitVal[i])
			row = append(row, intValue)
		}

		fmt.Printf("row: %v\n", row)
		ascending := true

		diff := 0
		for i := range len(row) - 1 {
			diff = diff + row[i] - row[i+1]
		}
		if diff < 0 {
			ascending = false
		}
		if ascending {
			slices.Reverse(row)
		}
		listOfRows = append(listOfRows, row)
		// fmt.Printf("rowCopy: %v\n", rowCopy)
		fmt.Printf("listOfRows: %v\n", listOfRows)
	}

	// Solve

	result := 0

	// skipNext := false

	for i := range len(listOfRows) {
		row := listOfRows[i]

		fmt.Printf("row: %v\n", row)

		currentVal := row[0]
		fault := false
		firstFault := true
		for j := range len(row) - 1 {
			nextVal := row[j+1]

			diff := nextVal - currentVal

			if diff > 3 || diff < 1 {
				if firstFault {
					firstFault = false
					continue
				} else {
					fault = true
					break
				}
			}

			currentVal = nextVal
		}

		if !fault {
			result += 1
			println("success")
		}
	}

	println(len(listOfRows))
	println("RESULT: ", result)

	println("Expected Example result: 4")
}

// too high 458 and 401 too low
// 459 too high
// 401 too low, again
