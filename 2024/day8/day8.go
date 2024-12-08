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
	land := readData()
	var antennaPointsMap map[string][]point = make(map[string][]point)
	// Populate points map
	for y, row := range land {
		for x, val := range row {
			if val == "." {continue}
			p := point{x: x, y: y}
			antennaPointsMap[val] = append(antennaPointsMap[val], p)
		}
	}

	_ = fillAntinodes(antennaPointsMap)
}


func fillAntinodes (points map[string][]point ) [][]string {
	// Lazily grab 2d array of map
	antinodes := readData()

	for key, p := range points {
		fmt.Println(key, "\t", p)
	}

	return antinodes
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
