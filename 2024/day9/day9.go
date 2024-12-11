package day9

import (
	"bufio"
	"fmt"
	"os"
)

var DATA_FILE_PATH string = "day9/input.txt"


func Solve() {
	readData()
}


func readData() {
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {panic(err)}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("New line", line)
	}
	return 
}
