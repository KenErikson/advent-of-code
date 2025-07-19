package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// inputExample := true

	filename := "day3/input_example.txt"
	
	if len(os.Args) > 1 {
		// inputExample = false
		filename = "day3/input.txt"
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		panic(err)
	}

	stringData := string(data)
	stringDataRows := strings.Split(strings.ReplaceAll(stringData, "\r", ""), "\n")

	res := 0
	for _, v := range stringDataRows {
		str := strings.Split(v, "mul(")
		for _, v2 := range str {

			b4, _, _ := strings.Cut(v2, ")")
			intStrs := strings.Split(b4, ",")
			if len(intStrs) == 2 {
				mul1, err := strconv.Atoi(intStrs[0])
				if err != nil{
					continue
				}
				mul2, err := strconv.Atoi(intStrs[1])
				
				if err != nil{
					continue
				}

				
				res += mul1 * mul2	
			}
		}
	}

	fmt.Println(res)

}
