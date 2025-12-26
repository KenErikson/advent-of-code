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

	// ### Day 1 ###
	fmt.Println("\n### Day 1 - Part 1 ###")

	stringData, stringDataRows, rows := ParseInputFile(os.Args)
	
	_, _, _ = stringData, stringDataRows, rows

	state := 50

	zeroCount := 0

	for _, v  := range stringDataRows {
		direction:=1
		if v[0] == 'L' {
			direction = -1
		}

		steps, err := strconv.Atoi(v[1:]) 

		if err != nil{
			End("!!!")
		}
		zeroCount += steps / 100
		steps = steps % 100
		println("Steps: ", steps , " ", direction)
		justState := state
		state += ( direction * steps )
		if(state<0){
			state += 100
			if justState != 0 {
				zeroCount += 1
			} 
		}
		if state / 100 > 0 && state != 100{
			zeroCount += 1
		}
		state = state % 100
		if state == 0 {
			zeroCount++
		}
		println("state", state)
		println("z ", zeroCount)
	}

	println("Result: ", zeroCount )
}
