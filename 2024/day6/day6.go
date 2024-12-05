package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var DATA_FILE_PATH string = "day6/data.txt"

func Solve() {
	grid := readData()
	x, y := getGuardPos(grid)
	uniquePositions := guardRunaround(x, y, grid)
	fmt.Println("Unique pos visited:", uniquePositions)
}


func guardRunaround(x, y int, grid [][]string) int {

	return -1
}



func getGuardPos (grid [][]string) (int, int) {
	for yIndex, row := range grid {
		for xIndex, xVal := range row {
			if xVal == "^" {
				return yIndex, xIndex
			}
		}
	}
	return -1, -1
}


func readData () [][]string {
	var grid [][]string = [][]string{}
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {panic(err)}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), "")
		grid = append(grid, lineSplit)
	}
	return grid
}
