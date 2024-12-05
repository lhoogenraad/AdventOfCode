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
		if validLine(line) {
			total++
		}
	}
	fmt.Printf("Total:\t%v\n", total)
}

func validLine (line string) bool {
	lineItems := strings.Split(line, " ")
	var asc bool = determineAsc(lineItems)
	prev, _ := strconv.Atoi(lineItems[0])
	fmt.Println(lineItems, asc)

	for i := 1; i < len(lineItems); i++ {
		numErrors := 0
		curr, _ := strconv.Atoi(lineItems[i])
		// If diff is positive, we are descending
		diff := curr - prev
		fmt.Println(prev, curr, diff)
		if diff > 3 || diff < -3 || diff == 0{
			return false
		}

		if asc {
			if diff < 0 {
				if numErrors >= 1 {
					return false
				} else {
					numErrors++
				}
			}
		} else {
			if diff > 0 {
				if numErrors >= 1 {
					return false
				} else {
					numErrors++
				}
			}
		}
		prev = curr
	}

	return true
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
