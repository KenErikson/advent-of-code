package main

import (
	"fmt"
	"os"
	_ "sort"
	"strconv"
	_ "strconv"
	_ "strings"

	. "aoc.erikson.fi/common"
)

func main() {

	// ### Day 3 ###
	fmt.Println("\n### Day 3 - Part 1 ###")

	stringData, stringDataRows, rows := ParseInputFile(os.Args)

	_, _, _ = stringData, stringDataRows, rows

	joltage := 0
	for _, row  := range stringDataRows {
		println(row)
		maxJoltage := 0
		for i := 0; i < len(row) - 1; i++ {
			valI, _ := strconv.Atoi(row[i:i+1])
			for j := i+1; j < len(row); j++ {
				valJ, _ := strconv.Atoi(row[j:j+1])
				possibleMaxJoltage, _ := strconv.Atoi(fmt.Sprintf("%d%d", valI, valJ))
				if possibleMaxJoltage > maxJoltage{
					println("New max joltage: ", possibleMaxJoltage)
					maxJoltage = possibleMaxJoltage
				}
			}
		}
		println("Max Joltage: ", maxJoltage)
		joltage += maxJoltage
	}

	println("Result: ", joltage)
}
