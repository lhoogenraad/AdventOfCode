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

	// Map y positions to list of x positions
	visitedPositions := make(map[int][]int)

	// Mark the starting position as visited
	visitedPositions[y] = append(visitedPositions[y], x)
	uniquePositionsVisited++

	for x >= 0 && y >= 0 && x < width && y < height {
		// fmt.Println("uniquePositionsVisited:", uniquePositionsVisited, "y:", y, "x:", x,
		// "direction vertical:", direction[0], "direction horiztonal:", direction[1])
		nextY, nextX := y+direction[0], x+direction[1]

		// Turn if there's an obstacle
		if nextY < 0 || nextY >= height || nextX < 0 || nextX >= width{
			dirIndex = (dirIndex + 1) % len(directions)
			direction = directions[dirIndex]
		} else if grid[nextY][nextX] == "#" {
			fmt.Println("Turning at (x, y):", x, y)
			dirIndex = (dirIndex + 1) % len(directions)
			direction = directions[dirIndex]
			continue
		}

		// Mark new position if not visited
		if !visited(nextY, nextX, visitedPositions) {
			visitedPositions[nextY] = append(visitedPositions[nextY], nextX)
			uniquePositionsVisited++
		}

		// Move to the next position
		y, x = nextY, nextX
	}

	fmt.Println("Exited at y:", y, ", x:", x)
	return uniquePositionsVisited
}


func visited (y, x int, visited map[int][]int) bool {
	if _, exists := visited[y]; !exists {
		return false
	}
	return slices.Contains(visited[y], x)
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
