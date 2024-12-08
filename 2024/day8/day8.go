package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
var DATA_FILE_PATH string = "day8/data.txt"

type point struct {
	x int
	y int
}

func Solve () {
	fmt.Println(readData())
}


func readData () [][]string {
	var matrix [][]string
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {panic(err)}
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		matrix = append(matrix, line)
	}
	return matrix
}
