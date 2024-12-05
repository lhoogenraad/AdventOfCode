package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var DATA_FILE_PATH string = "day5/data.txt"

func Solve() {
	total := 0
	orders, pages := readData()
	orderMap := buildOrderMap(orders)

	for _, page := range pages {
		if pageValid(page, orderMap) {
			continue
		}
		total += pageCanBeFixed(page, orderMap)
	}

	fmt.Println("Total:", total)
}

func pageCanBeFixed (page []int, orderMap map[int][]int) int {
	for i, num := range page {
		var successors []int = orderMap[num]
		for backtrack := i; backtrack >= 0; backtrack-- {
			backtrackNum := page[backtrack]
			// If a successor number is found to be before our number, swap them around
			if slices.Contains(successors, backtrackNum) {
				temp := page[backtrack]
				page[backtrack] = page[i]
				page[i] = temp
				i = backtrack
			}
		}
	}
	mid := page[len(page)/2]
	return mid
}

func pageValid (page []int, orderMap map[int][]int) bool {
	for i, num := range page {
		var successors []int = orderMap[num]
		for backtrack := i; backtrack >= 0; backtrack-- {
			backtrackNum := page[backtrack]
			// If a successor number is found to be before our number
			if slices.Contains(successors, backtrackNum) {
				return false
			}
		}
	}

	return true
}

func buildOrderMap(orders [][]int) map[int][]int {
	orderMap := make(map[int][]int)

	for _, currMap := range orders {
		before, after := currMap[0], currMap[1]
		orderMap[before] = append(orderMap[before], after)
	}

	return orderMap
}

func readData () ( [][]int, [][]int ){
	var orders [][]int = [][]int{}
	var pages [][]int = [][]int{}
	onOrders := true
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {panic(err)}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		scanText := scanner.Text()
		if scanText == "" {
			onOrders = false
			continue
		}

		if onOrders{
			stringList := strings.Split(scanText, "|")
			intList := []int{}
			for _, page := range stringList {
				pageInt, _ := strconv.Atoi(page)
				intList = append(intList, pageInt)
			}
			orders = append(orders, intList)
		} else {
			stringList := strings.Split(scanText, ",")
			intList := []int{}
			for _, page := range stringList {
				pageInt, _ := strconv.Atoi(page)
				intList = append(intList, pageInt)
			}
			pages = append(pages, intList)
		}
	}
	return orders, pages
}
