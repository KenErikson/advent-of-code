package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// Read input

	filename := "day§NUMBER§_input.txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		panic(err)
	}

	stringData := string(data)
	stringDataRows := strings.Split(stringData, "\n")

	rows := len(stringDataRows)

	fmt.Println("Read ", filename, " with ", rows, " rows. First row:\n", stringDataRows[0])

	// ### Day 1 ###
	fmt.Println("\n### Day 1 ###\n ")

	// Read data into two lists

	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, value := range stringDataRows {
		splitVal := strings.Split(value, " ")
		leftVal, _ := strconv.Atoi(splitVal[0])
		leftList = append(leftList, leftVal)

		rightVal, _ := strconv.Atoi(splitVal[3])
		rightList = append(rightList, rightVal)
	}

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	println(len(rightList))
	println(len(leftList))

	diff := 0
	for i := 0; i < rows; i++ {
		diffTmp := leftList[i] - rightList[i]
		if diffTmp > 0 {
			diff += diffTmp
		} else {
			diff -= diffTmp
		}
	}

	println("res: ", diff)
}
