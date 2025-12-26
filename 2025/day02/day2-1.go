package main

import (
	"fmt"
	"os"
	_ "sort"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"

	. "aoc.erikson.fi/common"
)

func main() {

	// ### Day 2 ###
	fmt.Println("\n### Day 2 - Part 2 ###")

	stringData, stringDataRows, rows := ParseInputFile(os.Args)

	_, _, _ = stringData, stringDataRows, rows

	ranges := strings.Split(stringDataRows[0], ",")

	invalidIds := []int{}
	for _,v  := range ranges {
		rangeSplit := strings.Split(v, "-")
		rangeStart, _ := strconv.Atoi(rangeSplit[0])
		rangeEnd, _ := strconv.Atoi(rangeSplit[1])
		for i := rangeStart; i <= rangeEnd; i++ {
			if isInvalidId(i, strconv.Itoa(i)){
				invalidIds = append(invalidIds, i)
			}
		}
	}

	res := 0
	for _,v  := range invalidIds {
		res += v
	}

	println("Result: ", res )
}

func isInvalidId(id int, idString string) bool{
	if id < 10 {
		return false
	}
	if len(idString) % 2 != 0{
		return false
	}
	if(idString[:len(idString)/2] == idString[len(idString)/2:]){
		return true
	}
	return false
}
