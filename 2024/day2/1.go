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

	rows := len(stringDataRows)

	fmt.Println("Read ", filename, " with ", rows, " rows. First row:\n", stringDataRows[0])

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
		listOfRows = append(listOfRows, row)
	}

	// Solve

	result := 0

	for i := range len(listOfRows) {
		ascending := true
		if listOfRows[i][0] < listOfRows[i][1] {
			ascending = false
		}
		row := listOfRows[i]
		if ascending {
			slices.Reverse(row)
		}
		fmt.Printf("row: %v\n", row)

		breakin := false
		lastVal := -1
		for j := range len(listOfRows[i]) {
			currentVal := listOfRows[i][j]
			if lastVal < 0 {
				lastVal = currentVal
				continue
			}

			diff := currentVal - lastVal

			if diff > 3 || diff < 1 {
				breakin = true
				break
			}

			lastVal = currentVal
		}

		if breakin {
			println("UNSAFE")
		} else {
			println("SAFE")
			result += 1
		}
	}

	println("RESULT: ", result)

	println("Expected Example result: 2")
}
