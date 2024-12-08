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

func Solve() {
	land := readData()
	antennaPointsMap := make(map[string][]point)
	uniqueAntinodes := make(map[point]struct{})

	// Populate points map
	for y, row := range land {
		for x, val := range row {
			if val == "." {
				continue
			}
			p := point{x: x, y: y}
			antennaPointsMap[val] = append(antennaPointsMap[val], p)
		}
	}

	findAntinodes(antennaPointsMap, uniqueAntinodes, len(land), len(land[5]))

	fmt.Println("Total unique antinodes:", len(uniqueAntinodes))
}

func findAntinodes(points map[string][]point, uniqueAntinodes map[point]struct{}, height, width int) {
	for _, coords := range points {
		for i, p1 := range coords {
			for j, p2 := range coords {
				if i == j {
					continue
				}

				antinodes := calculateAntinodes(p1, p2)
				for _, antinode := range antinodes {
					if isValid(antinode, height, width) {
						uniqueAntinodes[antinode] = struct{}{}
					}
				}
			}
		}
	}
}

func calculateAntinodes(p1, p2 point) []point {
	var antinodes []point

	dx := p2.x - p1.x
	dy := p2.y - p1.y

	antinode1 := point{x: p1.x - dx, y: p1.y - dy}
	antinode2 := point{x: p2.x + dx, y: p2.y + dy}

	antinodes = append(antinodes, antinode1, antinode2)
	return antinodes
}


func isValid(p point, height, width int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < width && p.y < height
}



func readData() [][]string {
	var matrix [][]string
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		matrix = append(matrix, line)
	}
	return matrix
}

