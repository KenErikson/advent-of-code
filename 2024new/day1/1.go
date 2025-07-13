package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("input_example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	list1 := []int{}
	list2 := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		nmbers := strings.Split(scanner.Text(), "   ")
		num1, err1 := strconv.Atoi(nmbers[0])
		if err1 != nil {
			break
		}
		println(scanner.Text())
		num2, _ := strconv.Atoi(nmbers[1])
		println(num1)
		println(num2)
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	// println(list1)
	for i, _ := range list1 {
		log.Printf("%v. and %v - %v", i, list1[i], list2[i])
		diff := int(math.Abs(float64(list1[i] - list2[i])))
		log.Printf("Diff: %v", diff)
		sum += diff
	}
	println(sum)
}
