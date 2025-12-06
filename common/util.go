package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func End(msg string, errs ...error){
	println("")
	println("Exiting with message: " + msg)
	println("")
	if len(errs) > 0 && errs[0] != nil {
		println("Error(s):")
		for i := 0; i < len(errs); i++ {
			println(errs[i])
		}
	}
	println("")
	println("Exiting")
	os.Exit(1)
}

func ParseInputFile( args []string )(stringData string, stringDataRows []string, rows int ){
	targetDir := args[1]
	
	inputFile := "input_example.txt"
	if len(args) > 2 {
		if args[2] == "full" {
			inputFile = "input_full.txt"
		}
	}

	fmt.Printf(" - %s \n", inputFile)

	inputFile = filepath.Join(targetDir, inputFile)

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Can't read file:", inputFile)
		panic(err)
	}

	stringData = string(data)
	stringDataRows = strings.Split(stringData, "\n")
	rows = len(stringDataRows)

	fmt.Println("Read ", inputFile, " with ", rows, " rows. First row:")
	println(stringDataRows[0])
	println()

	return
}