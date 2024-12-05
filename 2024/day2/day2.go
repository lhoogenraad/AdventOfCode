package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	lines := getData()
	total := 0

	for i := range lines {
		line := lines[i]
		if isSafeWithDampener(line) {
			total++
		}
	}
	fmt.Printf("Total:\t%v\n", total)
}

func validLine (lineItems []string) bool {
	var asc bool = determineAsc(lineItems)
	prev, _ := strconv.Atoi(lineItems[0])
	if lineItems[0] == "96" {
		fmt.Println(lineItems, "isAsc:", asc)
	}
	for i := 1; i < len(lineItems); i++ {
		curr, _ := strconv.Atoi(lineItems[i])
		// If diff is positive, we are descending
		diff := curr - prev
		if diff > 3 || diff < -3 || diff == 0{
			return false
		}

		if asc {
			if diff < 0 {
					return false
			}
		} else {
			if lineItems[0] == "96" {
				fmt.Println(curr, prev, diff)
			}
			if diff > 0 {
					return false
			}
		}
		prev = curr
	}

	return true
}


// Helper function to apply the Problem Dampener rule
func isSafeWithDampener(l string) bool {
	line := strings.Split(l, " ")
	// If the report is already safe, return true
	if validLine(line) {
		return true
	}
	fmt.Println("Og line:", line)
	// Try removing each level one by one
	for i := 0; i < len(line); i++ {
		modified := line
		modified = remove(modified, i)
		if validLine(modified) {
			fmt.Println("Modded at index found to be valid:", i, "\tArray:", modified)
			return true
		}
	}

	return false
}

func remove(s []string, i int) []string {
	var list []string
	for x, val := range s {
		if x != i {
			list = append(list, val)
		}
	}
	return list
}

func determineAsc (lineItems []string) bool {
	if len(lineItems) < 2 {
		//idk
		return true
	}
	item1, _ := strconv.Atoi(lineItems[0])
	item2, _ := strconv.Atoi(lineItems[1])

	return item1 < item2
}

func getData() []string {
	var lines []string
	file, err := os.Open("day2/data.txt")
	if err != nil {
		fmt.Println(err)
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return lines
	}
	return lines
}
