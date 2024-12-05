package day1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"sort"
	"os"
)

var DATA_FILE_PATH = "./day1/data.txt"
var list_one []int
var list_two []int

func Solve() {
	readData()
	partTwo()
}

// On^2 go brr
func partTwo() {
	var total int = 0

	for i := range list_one {
		leftNum := list_one[i]
		for j := range list_two {
			rightNum := list_two[j]
			if leftNum == rightNum {
				total += leftNum
			}
		}
	}
	fmt.Println("Total similarity score:", total)
}

func partOne() {
	var diffTotal int = 0
	sort.Ints(list_one[:])
	sort.Ints(list_two[:])

	for i := range list_one {
		var diff int = findPositiveDiff(list_one[i], list_two[i]) 
		diffTotal += diff
	}

	fmt.Println("Diffs:", diffTotal)
}

func findPositiveDiff (i int, a int) int {
	if i > a {
		return i - a
	} else {
		return a - i
	}
}

func readData () {
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {panic(err)}
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) == 4 {
			num1, _ := strconv.Atoi(line[0])
			num2, _ := strconv.Atoi(line[3])
			list_one = append(list_one, num1)
			list_two = append(list_two, num2)
		}
	}
}
