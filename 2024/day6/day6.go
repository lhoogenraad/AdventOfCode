package day6

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var DATA_FILE_PATH string = "day6/data.txt"

func Solve() {
	grid := readData()
	x, y := getGuardPos(grid)
	uniquePositions := guardRunaround(x, y, grid)
	fmt.Println("Unique pos visited:", uniquePositions)
}

// Format: [y, x]
var directions [][]int = [][]int{
	{1, 0}, // Up
	{0, 1}, //Right
	{-1, 0}, //Down
	{0, -1}, //Left
}
func guardRunaround(x, y int, grid [][]string) int {
	uniquePositionsVisited := 0
	// Start out going upwards
	dirIndex := 0
	direction := directions[dirIndex]

	// Declare grid height and width
	height, width := len(grid), len(grid[0])

	// Map y positions to list of x positions (I am lazy)
	visitedPositions := make( map[int][]int)

	for x > 0 && y > 0 && x < width && y < height {
		nextY, nextX := (y + direction[0]), (x + direction[1])
		if grid[nextY][nextX] == "#" {
			fmt.Println("Found turn point", nextY, nextX)
			newDirIndex := (dirIndex + 1) % len(directions)
			dirIndex = newDirIndex
			direction = directions[newDirIndex]

		}
		if !visited(y, x, visitedPositions) {
			uniquePositionsVisited += 1
		}
		visitedPositions[y] = append(visitedPositions[y], x)
		y += direction[0]
		x += direction[1]
	}
	fmt.Println("Exited at y:", y, ", x:", x)
	return uniquePositionsVisited
}


func visited (y, x int, visited map[int][]int) bool {
	fmt.Println("Checking y:", y, "x:", x, visited[y], "contains x:", slices.Contains(visited[y], x))
	if !slices.Contains(visited[y], x) {
		return false
	}
	return true
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
