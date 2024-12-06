package day6

import (
	"bufio"
	"fmt"
	"os"
)

var DATA_FILE_PATH string = "day6/data.txt"

func readData () [][]rune {
	var grid [][]rune = [][]rune{}
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {panic(err)}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		chars := []rune{}
		for _, char := range line {
			chars = append(chars, char)
		}
		grid = append(grid, chars)
	}
	return grid
}



type Position struct {
	x, y int
}

func Solve() {
	// Input map
	grid := readData()

	// Directions: up, right, down, left
	directions := []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	directionIndex := 0 // Start facing up

	// Find the guard's starting position
	var guard Position
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '^' {
				guard = Position{i, j}
				grid[i][j] = '.' // Replace initial guard position with empty space
			}
		}
	}

	visited := make(map[Position]bool)
	visited[guard] = true

	for {
		fmt.Println("(x, y):", guard.x, guard.y)
		// Compute the next position based on current direction
		next := Position{guard.x + directions[directionIndex].x, guard.y + directions[directionIndex].y}

		// Check if the next position is out of bounds
		if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
			fmt.Println("Guard has exited the grid.")
			break
		}

		// Check if the next position is obstructed
		if grid[next.x][next.y] == '#' {
			// Turn right (90 degrees clockwise)
			directionIndex = (directionIndex + 1) % 4
		} else {
			// Move forward
			guard = next
			visited[guard] = true
		}
	}

	// Output the number of distinct positions visited
	fmt.Println("Distinct positions visited:", len(visited))
}

