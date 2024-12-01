package main

import (
	"fmt"
	"sort"
	"thefjs14/aoc2024/2024/utils"
)

func Distances() int {
	content := utils.ReadFileToString("1.txt")
	lines := utils.StringPerLine(content)
	matrix := utils.StringToIntMatrix(lines)
	total := 0

	var left_matrix = []int{}
	var right_matrix = []int{}

	for index := range matrix {
		left_matrix = append(left_matrix, matrix[index][0])
		right_matrix = append(right_matrix, matrix[index][1])
	}

	sort.Ints(left_matrix)
	sort.Ints(right_matrix)

	for index := range left_matrix {
		distance := left_matrix[index] - right_matrix[index]
		if distance < 0 {
			distance = -distance
		}
		total = total + distance

	}

	return total
}

func Similarity() int {
	similarity := 0

	content := utils.ReadFileToString("1.txt")
	lines := utils.StringPerLine(content)
	matrix := utils.StringToIntMatrix(lines)

	var left_matrix = []int{}
	var right_matrix = []int{}

	for index := range matrix {
		left_matrix = append(left_matrix, matrix[index][0])
		right_matrix = append(right_matrix, matrix[index][1])
	}

	countMap := make(map[int]int)
	for _, value := range right_matrix {
		countMap[value]++
	}

	for _, value := range left_matrix {
		similarity += value * countMap[value]
	}
	return similarity
}

func main() {
	distance := Distances()
	fmt.Printf("Distance: %d\n", distance)

	similarity := Similarity()
	fmt.Printf("Similarity: %d", similarity)
}
