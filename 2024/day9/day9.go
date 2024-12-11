package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var DATA_FILE_PATH string = "day9/input.txt"


func Solve() {
	rawArray := readData()
	fmt.Println(rawArray)
}


func readData() []int {
	var items []int
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {panic(err)}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := scanner.Text()
		num,_ := strconv.Atoi(char)
		items = append(items, num)
	}
	return items
}
