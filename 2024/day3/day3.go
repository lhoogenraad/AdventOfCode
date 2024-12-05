package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DATA_FILE_PATH string = "day3/data.txt"

var directions = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1}, 
	{-1, -1},
	{1, -1}, 
	{-1, 1},  
}



func Solve1() {
	total := 0
	matrix := readData()
	for vert := 0; vert < len(matrix); vert++{
		for hor := 0; hor < len(matrix[vert]); hor++ {
			for _, direction := range directions {
				dirVert, dirHor := direction[0], direction[1]
				if checkWord(matrix, hor, vert, dirHor, dirVert, "XMAS") {
					total++
				}
			}
		}
	}
	fmt.Printf("Total: %v", total)
}


func checkWord(grid [][]string, vert, hor, dx, dy int, word string) bool {
	n, m := len(grid), len(grid[0])
	for i, ch := range word {
		// Am lazy
		str := strconv.QuoteRune(ch)
		nx, ny := hor+i*dx, vert+i*dy
		fmt.Println("Found match between:", grid[nx][ny], str)
		if grid[nx][ny] == str {
			fmt.Println("Found match between:", grid[nx][ny], str)
		}
		if nx < 0 || ny < 0 || nx >= n || ny >= m || grid[nx][ny] != str {
			return false
		}
	}
	return true
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
