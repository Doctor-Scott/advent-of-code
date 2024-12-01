package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseRowString(rowStr string, leftSlice []int, rightSlice []int) ([]int, []int) {
	splitRow := strings.Split(rowStr, "   ")

	leftString := splitRow[0]
	rightString := splitRow[1]

	leftNumber, _ := strconv.Atoi(leftString)
	rightNumber, _ := strconv.Atoi(rightString)
	return append(leftSlice, leftNumber), append(rightSlice, rightNumber)

}

func findDiff(left int, right int) int {
	if left > right {
		return left - right
	}
	if left < right {
		return right - left
	}
	return 0
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	inputString := string(input)
	splitStr := strings.Split(inputString, "\n")
	var leftArr []int
	var rightArr []int
	for _, row := range splitStr {
		if row != "" {
			leftArr, rightArr = parseRowString(row, leftArr, rightArr)
		}
	}
	slices.Sort(leftArr)
	slices.Sort(rightArr)
	count := 0
	for i := 0; i < len(leftArr); i++ {
		count += findDiff(leftArr[i], rightArr[i])
	}
	fmt.Println(count)
	// 2430334
}
