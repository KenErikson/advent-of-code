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
	if id <= 10 {
		return false
	}

	idLength := len(idString)
	// println(idString, " - ", idLength)
	for i, _ := range idString {

		if i > idLength/2 || idLength == i + 1{
			// println("i gone too far")
			break
		}

		isItEvenSplit := idLength % (i+1)
		if isItEvenSplit != 0 {
			// println("Uneven idLegnth: ", isItEvenSplit )
			continue
		}

		repeats := idLength / (i+1)
		// println("repeats",repeats)
		possibleRepeatString := idString[:(i+1)]
		// println("possibleRepeatString", possibleRepeatString)
		stringToTest := strings.Repeat(possibleRepeatString, repeats)

		// println("Testing string: ", stringToTest)
		if stringToTest == idString {
			println("repeats",repeats)
			println("possibleRepeatString", possibleRepeatString)
			println("InvalidID: ", idString)
			return true
		}
	}

	return false
}
