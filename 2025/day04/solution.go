package dayday04

import (
	"strings"
)

func Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	lineLen := len(lines)
	rowLen := len(lines[0])
	println(lineLen)
	println(rowLen)
	
	matrix := [][]rune{}
	for _, line := range lines {
		chars := []rune(line)
		matrix = append(matrix, chars)
	}

	res := 0
	for i, chars := range matrix {
		for j, char := range chars {
			if char == '.' {
				continue
			}

			close := 0
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if k == 0 && l == 0{
						continue
					}
					newI := i + k
					newJ := j + l
					
					if isOutsideCount(newI,lineLen) || isOutsideCount(newJ,rowLen) {
						continue
					}
					
					if matrix[newI][newJ] == '@'{
						close++
					}
				}
			}
			
			if close < 4 {
				res += 1
			}
		}
	}

	return res
}

func isOutsideCount(i, len int) bool {
	return i < 0 || i >= len
}

func Part2(input string) any {
	// replace with Part1 when Part 1 complete
	return 0
}
