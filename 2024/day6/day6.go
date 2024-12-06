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
    
    originalGrid := readData()

    
    directions := []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

    
    var guardStart Position
    for i := range originalGrid {
        for j := range originalGrid[i] {
            if originalGrid[i][j] == '^' {
                guardStart = Position{i, j}
                originalGrid[i][j] = '.' 
            }
        }
    }

    
    validObstructionCount := 0

    for x := 0; x < len(originalGrid); x++ {
        for y := 0; y < len(originalGrid[x]); y++ {
			fmt.Println(x, y)
            
            if originalGrid[x][y] != '.' {
                continue
            }

            
            grid := copyGrid(originalGrid)
            grid[x][y] = '#'

            
            if causesInfiniteLoop(grid, guardStart, directions) {
                validObstructionCount++
            }
        }
    }

    
    fmt.Println("Valid obstruction positions:", validObstructionCount)
}

func causesInfiniteLoop(grid [][]rune, start Position, directions []Position) bool {
    visited := make(map[Position]int) 
    directionIndex := 0               
    guard := start

    for {
        
        next := Position{guard.x + directions[directionIndex].x, guard.y + directions[directionIndex].y}

        
        if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
            return false 
        }

        
        if grid[next.x][next.y] == '#' {
            
            directionIndex = (directionIndex + 1) % 4
        } else {
            
            guard = next

            
            state := Position{guard.x*10 + directionIndex, guard.y*10 + directionIndex}
            visited[state]++
            if visited[state] > 1 {
                return true 
            }
        }
    }
}

func copyGrid(grid [][]rune) [][]rune {
    copy := make([][]rune, len(grid))
    for i := range grid {
        copy[i] = make([]rune, len(grid[i]))
        copy[i] = append([]rune{}, grid[i]...)
    }
    return copy
}
