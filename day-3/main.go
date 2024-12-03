package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func readFileByLine() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numSlice := make([]int, 0, 1000)

	for scanner.Scan() {
		row := findAllMaches(scanner.Text())
		exctractAndAppendResult(row, &numSlice)
		// log.Println(row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	counter := 0

	for _, val := range numSlice {
		counter += val
	}

	log.Printf("Final result: %v", counter)

}

func findAllMaches(input string) *[]string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(input, -1)
	return &matches

}

func exctractAndAppendResult(stringSlice *[]string, intSlice *[]int) {
	for _, val := range *stringSlice {
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(val, -1)
		num1, err1 := strconv.Atoi(matches[0])
		num2, err2 := strconv.Atoi(matches[1])

		if err1 != nil || err2 != nil {
			log.Panicf("first:%v, second:%v ", err1, err2)
		}
		*intSlice = append(*intSlice, num1*num2)

	}
}

func main() {
	start := time.Now()
	readFileByLine()
	elapsed := time.Since(start)
	log.Printf("The time it took %s", elapsed)
}
