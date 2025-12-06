package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	. "aoc.erikson.fi/common"
)

func main() {
	if len(os.Args) < 2 {
		println(strings.Join(os.Args, ", "))
		printUsage()
	}
	runMode, year, day, part, inputMode := parseArgs(os.Args[1:])

	println("RunMode", runMode)
	println("Year", year)
	println("Day", day)
	println("Part", part)
	println("InputMode", inputMode)

	switch runMode {
	case NormalRunMode:
		run(strconv.Itoa(year), strconv.Itoa(day), strconv.Itoa(part), inputMode)
	case InitRunMode:
		initDay(year, day)
	default:
		End("Dunno")
	}
}

func printUsage() {
	println("Example usage:")
	println("  go run aoc.go <opt|year: default current year til december> <req|day> <opt|part:(def) 1/2> <opt|inputMode: (def) example/full>")
	println("  go run aoc.go 2025 1  # run 2025 day 1 part 2 with example input")
	println("  go run aoc.go 1       # run 202x day 1 part 1")
	println()
	println("  go run aoc.go init <opt|year: default current year until next december> <req|1>")
	println("  go run aoc.go init 2025 1")
	println("  go run aoc.go init 2")
	println()
	os.Exit(1)
}

func run(year string, day string, part string, inputMode InputMode) {
	dirPath := filepath.Join(year, "day"+day)
	goPath := filepath.Join(dirPath, "day"+day+"-" + part + ".go")
	args := []string{"run", goPath, dirPath}
	if inputMode == FullInputMode {
		args = append(args, "full")
	}
	cmd := exec.Command("go", args... )

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	println(" - " + strings.Join(cmd.Args, " "))

	err := cmd.Run()

    if err != nil {
        fmt.Println("Error:", err)
    }
}

func initDay(year int, day int) {
	targetDirectory := filepath.Join(fmt.Sprintf("%d", year), fmt.Sprintf("day%d", day))
	if createErr := os.MkdirAll(targetDirectory, 0755); createErr != nil {
		panic(createErr)
	}

	inputFileExample := filepath.Join(targetDirectory, "input_example.txt")
	inputFileFull := filepath.Join(targetDirectory, "input_full.txt")
	part1File := filepath.Join(targetDirectory, fmt.Sprintf("day%d-1.go", day))
	part2File := filepath.Join(targetDirectory, fmt.Sprintf("day%d-2.go", day))

	contentBytes, readErr := os.ReadFile(filepath.Join("common", "setup", "dayExample.go.example"))
	if readErr != nil {
		panic(readErr)
	}

	dayExampleString := string(contentBytes)

	part1Content := strings.ReplaceAll(dayExampleString, "$DAY$", strconv.Itoa(day))
	part1Content = strings.ReplaceAll(part1Content, "$PART$", "1")
	part1Content = strings.ReplaceAll(part1Content, "$YEAR$", strconv.Itoa(year))
	part2Content := strings.ReplaceAll(dayExampleString, "$DAY$", strconv.Itoa(day))
	part2Content = strings.ReplaceAll(part2Content, "$PART$", "2")
	part2Content = strings.ReplaceAll(part2Content, "$YEAR$", strconv.Itoa(year))

	createFile(inputFileExample, "")
	createFile(inputFileFull, "")
	createFile(part1File, part1Content)
	createFile(part2File, part2Content)

}

func createFile(fileToCreate string, content string) {
	destinationFile, createDestErr := os.OpenFile( // Open -> create, fail if exists
		fileToCreate,
		os.O_CREATE|os.O_EXCL|os.O_WRONLY,
		0644,
	)
	if createDestErr != nil {
		panic(createDestErr)
	}
	defer destinationFile.Close()

	_, writeErr := destinationFile.WriteString(content)
	if writeErr != nil {
		panic(writeErr)
	}
}

func getYearArgOrCurrentAocYear(intArgs []int) (intArgsOut []int, year int) {
	if intArgs[0] > 2000 && intArgs[0] < 2100 {
		year = intArgs[0]
		intArgsOut = intArgs[1:]
		return
	}

	currentDate := time.Now()
	year = currentDate.Year()
	if currentDate.Month() != 12 {
		year -= 1
	}
	intArgsOut = intArgs
	return
}

func parseArgs(args []string) (runMode RunMode, year int, day int, part int, inputMode InputMode) {
	intArgs, runMode, inputMode := parseRunMode(args)
	intArgs, year = getYearArgOrCurrentAocYear(intArgs)
	day = 0
	part = 1

	day = intArgs[0]
	if day > 31 || day < 1 {
		End(fmt.Sprintf("Day is not as expected %d", day))
	}
	switch {
	case len(intArgs) == 1:
	case len(intArgs) == 2:
		part = intArgs[1]
		if part != 1 && part != 2 {
			End(fmt.Sprintf("Part is not 1 or 2: %d", part))
		}
	default:
		End("Unexpected intag length: " + strconv.Itoa(len(intArgs)))
	}

	return
}

func parseRunMode(args []string) (intArgs []int, runMode RunMode, inputMode InputMode) {
	runMode = NormalRunMode
	inputMode = ExampleInputMode

	if args[0] == "init" {
		runMode = InitRunMode
		args = args[1:]
	}

	if len(args) < 1 {
		End("not enough args left: " + strings.Join(args, ", "))
	}

	if args[len(args)-1] == "full" || args[len(args)-1] == "example" {
		if args[len(args)-1] == "full" {
			inputMode = FullInputMode
		}
		args = args[:len(args)-1]
	}

	intArgs = []int{}
	for _, v := range args {
		intArg, err := strconv.Atoi(v)
		if err != nil {
			End("Remaining args should be ints, but was: "+strings.Join(args, ","), err)
		}
		intArgs = append(intArgs, intArg)
	}

	if len(intArgs) < 1 {
		End("Not enough int args")
	}

	return
}
