package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pair struct {
	first  int
	second int
}

func readFileByLine(processLine func(string) []int) {

	counter := 0

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num1 := processLine(scanner.Text())
		res := PairSliceFromSlice(num1)
		if isReportSafe(res) {
			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("%v reports are safe", counter)

}

func isReportSafe(pairs []Pair) bool {
	if len(pairs) < 1 {
		log.Panic("result shorter than two")
	}

	if pairs[0].first > pairs[0].second {
		for _, val := range pairs {
			result := val.first - val.second
			if result < 1 || result > 3 {
				return false
			}
		}
		return true
	}

	if pairs[0].first < pairs[0].second {
		for _, val := range pairs {
			result := val.second - val.first
			if result < 1 || result > 3 {
				return false
			}
		}
		return true
	}
	return false
}

func PairSliceFromSlice(slice []int) []Pair {

	pair := make([]Pair, 0)

	if len(slice) < 2 {
		return pair
	}

	for idx, val := range slice {
		if idx == 0 {
		} else {
			pair = append(pair, Pair{first: slice[idx-1], second: val})
		}
	}

	return pair
}

func processLine(input string) []int {
	result := strings.Fields(input)

	nums := make([]int, 0, 10)

	for _, val := range result {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Panic("error:", err)
		}
		nums = append(nums, num)
		// nums[idx] = num
	}

	return nums

}

func main() {
	start := time.Now()
	readFileByLine(processLine)
	elapsed := time.Since(start)
	log.Printf("The time it took %s", elapsed)
}
