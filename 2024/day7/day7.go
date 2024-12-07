package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var DATA_FILE_PATH string = "day7/input.txt"


func Solve(){
	dict := readData()
	total := 0
	i := 0
	for target, nums := range dict {
		fmt.Println("i:", i)
		i++
		if getCalibrationResult(nums[0], nums, 1, target) {
			total += target
		}
	}
	fmt.Println("Total calibration rizzult:", total)
}


func getCalibrationResult(current int, nums []int, i int, target int) bool {
    if i == len(nums) {
        return current == target
    }

    addition := current + nums[i]
    multiplication := current * nums[i]
    concatenation := concatenate(current, nums[i])

    return getCalibrationResult(addition, nums, i+1, target) ||
           getCalibrationResult(multiplication, nums, i+1, target) ||
           getCalibrationResult(concatenation, nums, i+1, target)
}

// Thank you SO
func concatenate(a, b int) int {
	num, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return num
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
