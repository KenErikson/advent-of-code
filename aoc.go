package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type RunMode int
const (
	NormalRunMode RunMode = iota
	InitRunMode
)

type InputMode int
const (
	ExampleInputMode InputMode = iota
	FullInputMode
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}

	runMode, year, day, part, inputMode := parseArgs(os.Args[1:])

	switch runMode {
	case NormalRunMode:
		run(year, day, part, inputMode)
	case InitRunMode:
		initDay(year, day)
	}
}

func run(year int, day int, part int, inputMode InputMode) {
	dayStr := fmt.Sprintf("day%02d", day)
	dirPath := filepath.Join(strconv.Itoa(year), dayStr)

	// 1. Build the Regex for -run
	// Format: TopLevelTest/SubTest  (e.g., TestPart1/Example)
	
	// Start with the Function Name
	testRegex := "TestPart"
	if part == 1 {
		testRegex += "1"
	} else if part == 2 {
		testRegex += "2"
	}
	// If part == 0, "TestPart" matches both TestPart1 and TestPart2

	// Append the Subtest Name
	if inputMode == ExampleInputMode {
		testRegex += "/Example"
	} else if inputMode == FullInputMode {
		testRegex += "/Full"
	}

	// 2. Build Arguments
	// -count=1 bypasses the test cache
	args := []string{"test", "-v", "-count=1", "./" + dirPath, "-run", testRegex}

	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Pretty print what we are doing
	modeStr := "Example"
	if inputMode == FullInputMode { modeStr = "Full" }
	partStr := "Both Parts"
	if part != 0 { partStr = fmt.Sprintf("Part %d", part) }

	fmt.Printf("📂 Running %d/%s (%s, %s)...\n", year, dayStr, partStr, modeStr)
	
	// Run
	_ = cmd.Run()
}

func initDay(year int, day int) {
	dayStr := fmt.Sprintf("day%02d", day)
	targetDirectory := filepath.Join(strconv.Itoa(year), dayStr)

	if err := os.MkdirAll(targetDirectory, 0755); err != nil {
		panic(err)
	}

	// Template map
	files := map[string]string{
		"solution.go":      "common/setup/solution.go.example",
		"solution_test.go": "common/setup/solution_test.go.example",
		"input_example.txt": "", 
		"input_full.txt":    "",
	}

	for fileName, templatePath := range files {
		targetPath := filepath.Join(targetDirectory, fileName)
		content := ""

		if templatePath != "" {
			rawContent, err := os.ReadFile(templatePath)
			if err != nil {
				panic(fmt.Sprintf("❌ Missing template: %s", templatePath))
			}
			content = string(rawContent)
			content = strings.ReplaceAll(content, "$DAY$", dayStr)
			content = strings.ReplaceAll(content, "$YEAR$", strconv.Itoa(year))
		}

		createFile(targetPath, content)
	}

	fmt.Printf("✅ Initialized %s/%s\n", strconv.Itoa(year), dayStr)
}

func createFile(path string, content string) {
	// Exclusive create (fail if exists)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsExist(err) { return } // Skip if exists
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}

func parseArgs(args []string) (runMode RunMode, year, day, part int, inputMode InputMode) {
	// --- DEFAULTS ---
	runMode = NormalRunMode
	inputMode = ExampleInputMode // Default to Example
	part = 1                     // Default to Part 1

	// Determine Year
	now := time.Now()
	year = now.Year()
	if now.Month() < 12 { year-- }

	// Check for INIT
	if len(args) > 0 && args[0] == "init" {
		runMode = InitRunMode
		args = args[1:]
	}

	// Scan Args
	ints := []int{}
	for _, arg := range args {
		if arg == "example" { inputMode = ExampleInputMode; continue }
		if arg == "full"    { inputMode = FullInputMode;    continue }
		
		val, err := strconv.Atoi(arg)
		if err == nil { ints = append(ints, val) }
	}

	// Assign Ints (Year / Day / Part)
	if len(ints) > 0 && ints[0] > 2000 {
		year = ints[0]
		ints = ints[1:]
	}

	if len(ints) > 0 {
		day = ints[0]
		ints = ints[1:]
	} else {
		panic("❌ You must specify a Day.")
	}

	if len(ints) > 0 {
		part = ints[0] // Override default part 1 if specified
	}

	return
}

func printUsage() {
	fmt.Println(`Usage:
  go run aoc.go <day>                 # Run Part 1 Example (Default)
  go run aoc.go <day> full            # Run Part 1 Full
  go run aoc.go <day> 2               # Run Part 2 Example
  go run aoc.go <day> 2 full          # Run Part 2 Full
  go run aoc.go <day> 0 full          # Run All Parts Full
  
  go run aoc.go init <day>            # Init current year`)
	os.Exit(1)
}
