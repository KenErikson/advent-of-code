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
		rowCopy := make([]int, len(row))
		copy(rowCopy, row)
		slices.Reverse(rowCopy)

		listOfRows = append(listOfRows, row)
		listOfRows = append(listOfRows, rowCopy)
	}

	// Solve

	result := 0

	lastSuccess := false
	for i := range len(listOfRows) {
		if i%2 == 0 && lastSuccess {
			continue
		}
		row := listOfRows[i]

		fmt.Printf("row: %v\n", row)

		breakin := false
		firstFault := true
		lastVal := -1
		for j := range len(row) {
			currentVal := row[j]
			if lastVal < 0 {
				lastVal = currentVal
				continue
			}

			diff := currentVal - lastVal

			if diff > 3 || diff < 1 {
				if firstFault {
					firstFault = false
					continue
				} else {
					breakin = true
					break
				}
			}

			lastVal = currentVal
		}

		if breakin {
			println("UNSAFE ")
			lastSuccess = false
		} else {
			println("SAFE ")
			lastSuccess = true
			result += 1
		}
	}

	println("RESULT: ", result)

	println("Expected Example result: 4")
}

// too high 458 and 401 too low