package main

import (
	"os"
	"strings"
	"time"
	"aoc.erikson.fi/common"
)

func main() {

	if len(os.Args) < 2 {
		println(strings.Join(os.Args, ", "))
		printUsage()
	}
	runMode, year, day, part, inputMode := parseArgs(os.args[1:])
}



func printUsage (){
	println( "Example usage:" )
	println( "  go run aoc.go <opt|year: default current year til december> <req|day> <opt|part:(def) 1/2> <opt|inputMode: (def) example/full>" )
	println( "  go run aoc.go 2025 1  # run 2025 day 1 part 2 with example input" )
	println( "  go run aoc.go 1       # run 202x day 1 part 1" )
	println()
	println( "  go run aoc.go init <opt|year: default current year until next december> <req|1>" )
	println( "  go run aoc.go init 2025 1" )
	println( "  go run aoc.go init 2" )
	println()
}

func getCurrentAocYear() int{
	currentDate:= time.Now()
	year := currentDate.Year()
	if currentDate.Month() != 12 {
		year -= 1
	}
	return year
}

func parseArgs(args []string) (runMode common.RunMode, year int, day int, part int, inputMode common.InputMode){
	runMode = common.NormalRunMode
	if args[0] == "init" {
		runMode = common.InitRunMode
		args = args[1:]
	}

	value, err := strconv.Atoi(inputString)
	if err != nil {
		println( )
	}
}
