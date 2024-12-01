package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadFileToString(filename string) string {
	content, err := os.ReadFile("./files/" + filename)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func StringPerLine(content string) []string {
	return strings.Split(content, "\r\n")
}

func StringToIntMatrix(lines []string) [][]int {
	var matrix [][]int

	for _, line := range lines {
		row := strings.Split(line, "   ")
		left, err := strconv.Atoi(row[0])
		if err != nil {
			PanicErr(err)
		}
		right, err := strconv.Atoi(row[1])
		if err != nil {
			PanicErr(err)
		}
		newrow := []int{left, right}
		matrix = append(matrix, newrow)
	}

	return matrix
}

func PanicErr(message error) {
	panic(message)
}
