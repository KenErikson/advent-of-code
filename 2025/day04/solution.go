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
					if k == 0 && l == 0 {
						continue
					}
					newI := i + k
					newJ := j + l

					if isOutsideCount(newI, lineLen) || isOutsideCount(newJ, rowLen) {
						continue
					}

					if matrix[newI][newJ] == '@' {
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

	println("Here")
	printMatrix(matrix)
	for {
		currRes := res
		duplicate := make([][]rune, len(matrix))
		for i := range matrix {
			// 2. Create the inner slice for each row
			duplicate[i] = make([]rune, len(matrix[i]))
			// 3. Copy the contents from matrix to duplicate
			copy(duplicate[i], matrix[i])
		}

		for i, chars := range matrix {
			for j, char := range chars {
				if char == '.' {
					continue
				}

				close := 0
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						if k == 0 && l == 0 {
							continue
						}
						newI := i + k
						newJ := j + l

						if isOutsideCount(newI, lineLen) || isOutsideCount(newJ, rowLen) {
							continue
						}

						if matrix[newI][newJ] == '@' {
							close++
						}
					}
				}

				if close < 4 {
					res += 1
					duplicate[i][j] = '.'
				}
			}
		}

		println()
		printMatrix(matrix)

		if currRes == res {
			println("END")
			break
		}

		matrix = make([][]rune, len(duplicate))
		for i := range duplicate {
			// 2. Create the inner slice for each row
			matrix[i] = make([]rune, len(duplicate[i]))
			// 3. Copy the contents from duplicate to duplicate
			copy(matrix[i], duplicate[i])
		}

	}

	return res
}

func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, char := range row {
			if char == '.' {
				print(".")
			}
			if char == '@' {
				print("@")
			}
		}
		println()
	}
}
