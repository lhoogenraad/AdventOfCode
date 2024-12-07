package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var DATA_FILE_PATH string = "day7/input.txt"

// Search tree node
type node struct {
	val int
	add *node
	multiply *node
}

func Solve(){
	dict := readData()
	fmt.Println(dict)
}


func readData() map[int][]int {
	dict := make(map[int][]int)
	readFile, err := os.Open(DATA_FILE_PATH)
	defer readFile.Close()
	if err != nil {panic(err)}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		target, _ := strconv.Atoi(strings.Split(line, ":")[0])
		nums := strings.Split(line, " ")[1:]
		numsArray := make([]int, len(nums))
		for i, num := range nums {
			numsArray[i], _ = strconv.Atoi(num)
		}
		dict[target] = numsArray
	}
	return dict
}
