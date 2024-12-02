package main

import (
	"fmt"
	"os"
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

func countOccurances(number int, rightList []int) int {
	if len(rightList) == 0 {
		// base case
		return 0
	}

	for i, rightNumber := range rightList {
		if rightNumber == number {
			// dont need to pass in the full list, just the bit we still need to search
			return 1 + countOccurances(number, rightList[i+1:])
		}
	}
	// no occurances
	return 0
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	inputString := string(input)
	splitStr := strings.Split(inputString, "\n")
	var leftList []int
	var rightList []int
	for _, row := range splitStr {
		if row != "" {
			leftList, rightList = parseRowString(row, leftList, rightList)
		}
	}

	var count int
	for i := 0; i < len(leftList); i++ {
		currentListItem := leftList[i]
		occurances := countOccurances(currentListItem, rightList)
		if occurances != 0 {
			count += currentListItem * occurances
		}
	}

	// 28786472
	fmt.Println(count)
}
